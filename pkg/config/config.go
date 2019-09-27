package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is used to configure the service
type Config struct {
	Port                     int    `envconfig:"PORT" default:"1337"`
	DatabaseConnectionString string `envconfig:"DB_CONNECTIONSTRING" required:"true"`
}

// Get retrieves the config from environment
func Get() Config {
	config := Config{}

	err := envconfig.Process("shoppinglist_query", &config)
	if err != nil {
		panic("config error: " + err.Error())
	}

	return config
}
