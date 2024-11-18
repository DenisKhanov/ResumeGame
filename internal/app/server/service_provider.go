package server

import (
	grpcdata "github.com/DenisKhanov/ResumeGame/internal/server/api/grpc/data"
	repodata "github.com/DenisKhanov/ResumeGame/internal/server/repositories/data"
	srevicedata "github.com/DenisKhanov/ResumeGame/internal/server/services/data"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// serviceProvider manages the dependency injection for http_shortener-related components.
type serviceProvider struct {
	repositoryData srevicedata.RepoData // Repository interface for grpc_game-related data_info
	serviceData    grpcdata.ServData    // Service interface for grpc_game-related operations
	grpcData       *grpcdata.GRPCData   //GRPC for grpc_game-related operations
}

// newServiceProvider creates a new instance of the service provider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GameDataRepository returns the repository for data_info-related user_data_info.
// If dbPool is nil, it initializes an in-memory repository, otherwise initializes a database repository.
func (s *serviceProvider) GameDataRepository(dbPool *pgxpool.Pool) srevicedata.RepoData {
	logrus.Info("Creating new ResumeGame Data Repository.")
	if s.repositoryData == nil {
		logrus.Info("Initializing repository data_info.")
		if dbPool == nil {
			logrus.Infof("DB pool is nil.")
			return nil
			//TODO продумать хранение на жестком
		} else {
			s.repositoryData = repodata.NewPostgresData(dbPool)
		}
	}
	return s.repositoryData
}

// GameDataService returns the service for data_info-related operations.
func (s *serviceProvider) GameDataService(dbPool *pgxpool.Pool) grpcdata.ServData {
	if s.serviceData == nil {
		s.serviceData = srevicedata.NewServiceData(
			s.GameDataRepository(dbPool),
			dbPool,
		)
	}
	return s.serviceData
}

// GameDataGRPC returns GRPC data server operations
func (s *serviceProvider) GameDataGRPC(dbPool *pgxpool.Pool) *grpcdata.GRPCData {
	logrus.Info("Creating ResumeGame GRPC.")
	if s.grpcData == nil {
		logrus.Info("Initializing grpc data_info.")
		gameGRPC := grpcdata.NewGRPCData(s.GameDataService(dbPool))
		s.grpcData = gameGRPC
	}
	return s.grpcData
}
