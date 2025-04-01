package config

// Config holds configuration for the application.
type Config struct {
	Port string
}

// Load returns default configuration or reads from a file/env.
func Load() Config {
	return Config{
		Port: ":8080",
	}
}
