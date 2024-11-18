package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/DenisKhanov/ResumeGame/internal/server/config"
	"github.com/DenisKhanov/ResumeGame/pkg/db/postgres"
	"github.com/DenisKhanov/ResumeGame/pkg/logcfg"
	tlsCreds "google.golang.org/grpc/credentials"

	protodata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
	"github.com/DenisKhanov/ResumeGame/pkg/tlsconfig"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// App represents the application structure responsible for initializing dependencies
// and running the serverGRPC.
type App struct {
	serviceProvider *serviceProvider // The service provider for dependency injection
	config          *config.Config   // The configuration object for the application
	dbPool          *pgxpool.Pool    // The connection pool to the database
	serverGRPC      *grpc.Server     // The serverGRPC instance
	tls             *tls.Config
}

// NewApp creates a new instance of the application.
func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Run starts the application and runs the grpc_game serverGRPC.
func (a *App) Run() {
	a.runGameServer()
}

// initDeps initializes all dependencies required by the application.
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initTLS,
		a.initDBConnection,
		a.runMigrations,
		a.initServiceProvider,
		a.initGameGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// initConfig initializes the application configuration.
func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	a.config = cfg
	logFileName := "GameServer.log"
	logcfg.RunLoggerConfig(a.config.EnvLogLevel, logFileName)
	return nil
}

func (a *App) initTLS(_ context.Context) error {
	newTls, err := tlsconfig.NewServerTLS(a.config.ServerCert, a.config.ServerKey, a.config.ServerCa)
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize tls")
		return err
	}
	a.tls = newTls
	return nil
}

// initDBConnection initializes the connection to the database.
func (a *App) initDBConnection(ctx context.Context) error {
	logrus.Infof("DB config: %+v", a.config.DatabaseURI)
	if a.config.DatabaseURI == "" {
		logrus.Infof("config DatabaseURI is empty")
		return fmt.Errorf("config DatabaseURI is empty")
	}
	confPool, err := pgxpool.ParseConfig(a.config.DatabaseURI)
	if err != nil {
		logrus.WithError(err).Error("Error parsing config")
		return err
	}
	confPool.MaxConns = 50
	confPool.MinConns = 10
	a.dbPool, err = pgxpool.NewWithConfig(ctx, confPool)
	if err != nil {
		logrus.WithError(err).Error("Don't connect to DB")
		return err
	}
	if err = a.dbPool.Ping(ctx); err != nil {
		return err
	}
	return nil
}

func (a *App) runMigrations(_ context.Context) error {
	if err := postgres.MigrationsUp(a.dbPool); err != nil {
		logrus.WithError(err).Error("migrations up failed")
		return err
	}
	logrus.Info("Migrations done")
	return nil
}

// initServiceProvider initializes the service provider for dependency injection.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGameGRPCServer initializes the  serverGRPC with interceptors.
func (a *App) initGameGRPCServer(_ context.Context) error {
	gameDataGRPC := a.serviceProvider.GameDataGRPC(a.dbPool)

	server := grpc.NewServer(grpc.Creds(tlsCreds.NewTLS(a.tls)))
	reflection.Register(server)
	a.serverGRPC = server

	// registration service
	protodata.RegisterResumeDataV1Server(server, gameDataGRPC)
	return nil

}

// runGameServer starts the gRPC server with graceful shutdown.
func (a *App) runGameServer() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer close(signalChan)

	var wg sync.WaitGroup

	wg.Add(1)
	//run gRPC server
	go func() {
		defer wg.Done()
		listen, err := net.Listen("tcp", a.config.GRPCServer)
		if err != nil {
			logrus.Error(err)
		}

		logrus.Infof("Starting server gRPC on: %s", a.config.GRPCServer)
		if err = a.serverGRPC.Serve(listen); err != nil {
			logrus.WithError(err).Error("The server gRPC  failed to start")
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			a.dbPool.Close()
			wg.Done()
		}()
		for {
			select {
			case sig := <-signalChan:
				a.serverGRPC.GracefulStop()
				logrus.Infof("gRPC servers stopped with signal: %v", sig)
				return
			}
		}
	}()

	wg.Wait()
	logrus.Info("Server exited")
}
