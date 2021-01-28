package http

// Config is http service config.
type Config struct {
	Hostname string `json:"hostname" env:"HTTP_HOSTNAME" env-default:"0.0.0.0"`
	Port     string `json:"port" env:"HTTP_PORT" env-default:"8080"`
}

// ClientConfig is http service config.
type ClientConfig struct {
	Domain string `json:"domain" env:"HTTP_CLIENT_DOMAIN"`
}
