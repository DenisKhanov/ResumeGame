package client

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/DenisKhanov/ResumeGame/internal/client/config"
	servicedata "github.com/DenisKhanov/ResumeGame/internal/client/services/data"
	serviceuser "github.com/DenisKhanov/ResumeGame/internal/client/services/user"
	"github.com/DenisKhanov/ResumeGame/pkg/logcfg"
	"github.com/DenisKhanov/ResumeGame/pkg/tlsconfig"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	tlsCreds "google.golang.org/grpc/credentials"
	"os"
)

// App represents the application structure responsible for initializing dependencies
// and running the clientGRPC.
type App struct {
	serviceProvider *serviceProvider // The service provider for dependency injection
	config          *config.Config   // The configuration object for the application
	clientUser      *serviceuser.ServiceUser
	clientData      *servicedata.ServiceData
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

// Run starts the application and runs the grpc_game clientGRPC.
func (a *App) Run() {
	a.runGameClient()
}

// initDeps initializes all dependencies required by the application.
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initTLS,
		a.initServiceProvider,
		a.initGameGRPCClient,
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
	logcfg.RunLoggerConfig(a.config.EnvLogLevel, logFileName) //initialise logger
	return nil
}

func (a *App) initTLS(_ context.Context) error {
	newTls, err := tlsconfig.NewClientTLS(a.config.ClientCert, a.config.ClientKey, a.config.ClientCa)
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize tls")
		return err
	}
	a.tls = newTls
	return nil
}

// initServiceProvider initializes the service provider for dependency injection.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGameGRPCServer initializes the  serverGRPC with interceptors.
func (a *App) initGameGRPCClient(_ context.Context) error {

	conn, err := grpc.NewClient(a.config.GRPCServer, grpc.WithTransportCredentials(tlsCreds.NewTLS(a.tls)))
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize GRPC conn")
		return err
	}
	a.clientData = a.serviceProvider.DataClientService(conn)
	a.clientUser = a.serviceProvider.UserClientService(conn)
	return nil
}

// runGameClient starts the gRPC server with graceful shutdown.
func (a *App) runGameClient() {
	ctx := context.Background()
	scanner := bufio.NewScanner(os.Stdin)

	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for {
		if a.clientUser.State.GetLogin() != "" {
			fmt.Printf(green("\nYou are authorized as %s\n"), blue(a.clientUser.State.GetLogin()))
		} else {
			fmt.Printf(red("\nYou are not authorized, please login or register\n\n"))
		}
		//TODO реализовать подпункты
		fmt.Println(blue("Input command number to proceed\n"))
		fmt.Println("[1] - login")
		fmt.Println("[2] - register")
		fmt.Println(blue("---------------------------------------------"))
		fmt.Println("[3] - load about game info")
		fmt.Println("[4] - load about creator info")
		fmt.Println("[5] - load creator's experience")
		fmt.Println("[6] - load creator's skills")
		fmt.Println("[7] - load all creator's projects")
		fmt.Println("[8] - load creator's contacts")
		fmt.Println(blue("------------"))
		fmt.Println(red("[0] - quit"), blue("|"))
		fmt.Println(blue("------------"))
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			a.clientUser.LoginUser(ctx)
		case "2":
			a.clientUser.RegisterUser(ctx)
		case "3":
			a.clientData.GetGameInfo(ctx)
		case "4":
			a.clientData.GetAboutOwner(ctx)
		case "5":
			a.clientData.GetPreviousJobs(ctx)
		case "6":
			a.clientData.GetSkills(ctx)
		case "7":
			a.clientData.GetProjectList(ctx)
		case "8":
			a.clientData.GetContacts(ctx)
		case "0":
			fmt.Println("Application shutdown.")
			return
		default:
			fmt.Println("Unknown command, please try again")
		}
	}
}
