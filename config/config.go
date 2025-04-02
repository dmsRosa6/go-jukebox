package config

// Config holds configuration for the application.
type Config struct {
	Port string
}

// Load returns default configuration or reads from a file/env.
func LoadDefaultConfig() Config {
	return Config{
		Port: ":8080",
	}
}

// Options holds the application runtime options
type Options struct {
	Capacity int
}

//Load the default options
func LoadDefaultOptions() Options{
	return Options{
		Capacity: 50,
	}
}

