package state

const (
	perm = 0o755 // Default permission for created directories (rwxr-xr-x)
)

// ClientState holds the state information of the client,
// including authorization status, token, login, and working directory.
type ClientState struct {
	login string // User login
}

// NewClientState creates and returns a new instance of ClientState.
func NewClientState() *ClientState {
	return &ClientState{}
}

// GetLogin retrieves the current login of the client.
func (c *ClientState) GetLogin() string {
	return c.login
}

// SetLogin sets the login for the client.
func (c *ClientState) SetLogin(login string) {
	c.login = login
}
