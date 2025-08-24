package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application configuration
// The struct tags `envconfig` are used to map environment variables to the struct fields
// Example: `envconfig:"DB_USER"` maps the DB_USER environment variable to the DBUser field
// The `required:"true"` tag ensures that the application will not start if the variable is not set
type Config struct {
	DBUser        string `envconfig:"DB_USER" required:"true"`
	DBPassword    string `envconfig:"DB_PASSWORD" required:"true"`
	DBHost        string `envconfig:"DB_HOST" required:"true"`
	DBPort        string `envconfig:"DB_PORT" required:"true"`
	DBName        string `envconfig:"DB_NAME" required:"true"`
	DatabaseURL   string `envconfig:"DATABASE_URL"` // Optional, can be built from other fields if not provided
	WebServerPort string `envconfig:"WEB_SERVER_PORT" default:"8080"`
	GraphQLPort   string `envconfig:"GRAPHQL_SERVER_PORT" default:"8081"`
	GRPCPort      string `envconfig:"GRPC_SERVER_PORT" default:"50051"`
}

// GetDatabaseURL returns the database connection string
func (c *Config) GetDatabaseURL() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

// LoadConfig loads the configuration from environment variables
func LoadConfig(filenames ...string) (*Config, error) {
	// Load the .env file. We ignore the error because the configuration can also be provided by system environment variables.
	_ = godotenv.Load(filenames...)

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing environment variables: %w", err)
	}

	return &cfg, nil
}
