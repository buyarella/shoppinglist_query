package config

// Config is used to configure the service
type Config struct {
	Port int
}

// Get retrieves the config from environment
func Get() Config {
	return Config{
		Port: 1337,
	}
}
