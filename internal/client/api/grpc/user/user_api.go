package user

import (
	"context"

	pb "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/user"
)

// GRPCUserClient is a client for interacting with the user service via gRPC.
// It holds a reference to the user service client interface.
type GRPCUserClient struct {
	userService pb.ResumeUserV1Client
}

// NewUserPBClient initializes a new DataPBClient with the provided user service client.
// It returns a pointer to the newly created DataPBClient.
func NewUserPBClient(u pb.ResumeUserV1Client) *GRPCUserClient {
	return &GRPCUserClient{
		userService: u,
	}
}

// LoginUser attempts to log in a user with the provided login credentials.
// It sends a login request to the user service and returns the user's token if successful.
func (u *GRPCUserClient) LoginUser(ctx context.Context, login string) error {
	req := &pb.SignInRequest{
		Login: login,
	}

	_, err := u.userService.SignIn(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// RegisterUser registers a new user with the provided login credentials.
// It sends a registration request to the user service and returns the user's token if successful.
func (u *GRPCUserClient) RegisterUser(ctx context.Context, login string) error {
	req := &pb.SignUpRequest{
		Login: login,
	}

	_, err := u.userService.SignUp(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
