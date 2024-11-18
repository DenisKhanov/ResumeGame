package client

import (
	dataapi "github.com/DenisKhanov/ResumeGame/internal/client/api/grpc/data"
	userapi "github.com/DenisKhanov/ResumeGame/internal/client/api/grpc/user"
	servicedata "github.com/DenisKhanov/ResumeGame/internal/client/services/data"
	serviceuser "github.com/DenisKhanov/ResumeGame/internal/client/services/user"
	"github.com/DenisKhanov/ResumeGame/internal/client/state"
	pbdata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
	pbuser "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// serviceProvider manages the dependency injection for http_shortener-related components.
type serviceProvider struct {
	//TODO перенести состояние в уровень инициализации приложения
	state         *state.ClientState
	userClientAPI pbuser.ResumeUserV1Client
	dataClientAPI pbdata.ResumeDataV1Client
	userClient    serviceuser.GRPCUser
	dataClient    servicedata.GRPCData
	userService   *serviceuser.ServiceUser
	dataService   *servicedata.ServiceData
}

// newServiceProvider creates a new instance of the service provider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{
		state: state.NewClientState(),
	}
}

func (s *serviceProvider) UserClientAPI(client grpc.ClientConnInterface) pbuser.ResumeUserV1Client {
	if s.userClientAPI == nil {
		logrus.Info("Initializing new user client API.")
		s.userClientAPI = pbuser.NewResumeUserV1Client(client)
	}
	return s.userClientAPI
}

func (s *serviceProvider) DataClientAPI(conn grpc.ClientConnInterface) pbdata.ResumeDataV1Client {
	if s.dataClientAPI == nil {
		logrus.Info("Initializing new data conn API.")
		s.dataClientAPI = pbdata.NewResumeDataV1Client(conn)
	}
	return s.dataClientAPI
}

// UserClient returns GRPC user client operations
func (s *serviceProvider) UserClient(conn grpc.ClientConnInterface) serviceuser.GRPCUser {
	if s.userClient == nil {
		logrus.Info("Initializing user client.")
		s.userClient = userapi.NewUserPBClient(s.UserClientAPI(conn))
	}
	return s.userClient
}

// DataClient returns GRPC data client operations
func (s *serviceProvider) DataClient(conn grpc.ClientConnInterface) servicedata.GRPCData {
	if s.dataClient == nil {
		logrus.Info("Initializing data client.")
		s.dataClient = dataapi.NewDataPBClient(s.DataClientAPI(conn))
	}
	return s.dataClient
}

func (s *serviceProvider) UserClientService(conn grpc.ClientConnInterface) *serviceuser.ServiceUser {
	if s.userService == nil {
		logrus.Info("Initialising user client service")
		s.userService = serviceuser.NewServiceUserClient(s.UserClient(conn), s.state)
	}
	return s.userService
}

func (s *serviceProvider) DataClientService(conn grpc.ClientConnInterface) *servicedata.ServiceData {
	if s.dataService == nil {
		logrus.Info("Initialising data client service")
		s.dataService = servicedata.NewServiceDataService(s.DataClient(conn), s.state)
	}
	return s.dataService
}
