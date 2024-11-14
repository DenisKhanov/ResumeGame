package user

import (
	"bufio"
	"context"
	"fmt"
	"github.com/DenisKhanov/ResumeGame/internal/client/state"
	"github.com/DenisKhanov/ResumeGame/internal/client/validate"
	"github.com/fatih/color"
	"os"
	"strings"
)

// GRPCUser is an interface that defines methods for user registration and login.
// Implementations of this interface should provide the actual functionality.
type GRPCUser interface {
	RegisterUser(ctx context.Context, login string) error
	LoginUser(ctx context.Context, login string) error
}

// ServiceUser is a struct that provides user-related functionalities.
// It contains a reference to a DataService and a ClientState.
type ServiceUser struct {
	userService GRPCUser           // The user service implementation
	State       *state.ClientState // Client state management
}

// NewServiceUserClient initializes a new DataProvider with the given user service and client state.
// It returns a pointer to the newly created DataProvider.
func NewServiceUserClient(u GRPCUser, state *state.ClientState) *ServiceUser {
	return &ServiceUser{
		userService: u,
		State:       state,
	}
}

// RegisterUser prompts the user for their login credentials to register.
// It handles input validation and calls the user service to perform the registration.
func (u *ServiceUser) RegisterUser(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	red := color.New(color.FgRed).SprintFunc()

	var login string

	yellowBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(yellowBold("Input 'login' to create new user:"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input login as %s: ", yellow("'all symbols length (6-12)'"))
	scanner.Scan()
	data := scanner.Text()
	login = strings.TrimSpace(data)
	if err := validate.CheckLogin(login); err != nil {
		fmt.Printf("%v, %s\n", err, red(" please try again"))
		return
	}

	if err := u.userService.RegisterUser(ctx, login); err != nil {
		fmt.Println(red(" please try again"))
		return
	} else {
		u.State.SetLogin(login)
	}
}

// LoginUser prompts the user for their login credentials to log in.
// It handles input validation and calls the user service to perform the login.
func (u *ServiceUser) LoginUser(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	red := color.New(color.FgRed).SprintFunc()

	var login string

	yellowBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(yellowBold("Input 'login password' to login:"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow("Input your login"))
	scanner.Scan()
	data := scanner.Text()
	login = strings.TrimSpace(data)
	if len(login) == 0 {
		fmt.Println(red("Login must not be empty please try again"))
		return
	}

	if err := u.userService.LoginUser(ctx, login); err != nil {
		fmt.Println(red(" User don't found, please try again"))
		return
	} else {
		u.State.SetLogin(login)
	}
}
