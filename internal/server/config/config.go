package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config holds the application configuration parameters.
// Each field corresponds to an expected environment variable.
type Config struct {
	EnvLogLevel string // Log level for the application (e.g., DEBUG, INFO)
	DatabaseURI string // Connection string for the database
	GRPCServer  string // Address of the gRPC server
	ServerCert  string // Path to the server's SSL certificate
	ServerKey   string // Path to the server's SSL key
	ServerCa    string // Path to the server's CA file
}

// New initializes a new Config instance by loading environment variables from a .env file.
// It returns a pointer to the Config struct and an error if any of the environment variables are missing or invalid.
func New() (*Config, error) {
	err := godotenv.Load("server.env")
	if err != nil {
		return nil, fmt.Errorf("new load .env: %w", err)
	}

	config := &Config{}
	config.EnvLogLevel = os.Getenv("LOG_LEVEL")
	config.DatabaseURI = os.Getenv("DATABASE_URI")
	config.GRPCServer = os.Getenv("GRPC_SERVER")
	config.ServerCert = os.Getenv("SERVER_CERT_FILE")
	config.ServerKey = os.Getenv("SERVER_KEY_FILE")
	config.ServerCa = os.Getenv("SERVER_CA_FILE")

	return config, nil
}
