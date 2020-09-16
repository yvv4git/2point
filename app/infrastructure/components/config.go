package components

import (
	"fmt"
)

const defaultPort = 8080
const defaultPrefix = "[srv] "

// Config main config of app.
type Config struct {
	port   int
	prefix string
}

// SetPort is a setter for property port.
func (c *Config) SetPort(port int) {
	c.port = port
}

// SetPrefix is a setter for property prefix.
func (c *Config) SetPrefix(prefix string) {
	c.prefix = prefix
}

// GetPort is a getter.
func (c *Config) GetPort() int {
	return c.port
}

// GetPrefix is a getter.
func (c *Config) GetPrefix() string {
	return c.prefix
}

// GenerateServerAddress for generate server address.
func (c *Config) GenerateServerAddress() string {
	return fmt.Sprintf(":%d", c.port)
}

// NewConfig is a constructor.
func NewConfig() *Config {
	var config Config
	config.SetPort(defaultPort)
	config.SetPrefix(defaultPrefix)
	return &config
}
