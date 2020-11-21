package http

// Config is http service config.
type Config struct {
	Hostname string `json:"hostname" env:"HTTP_HOSTNAME" env-default:"0.0.0.0"`
	Port     string `json:"port" env:"PORT" env-default:"8080"`
}

// Config is http client config.
type ConfigClient struct {
	Hostname string `json:"hostname" env:"HTTP_CLIENT_HOSTNAME" env-default:"0.0.0.0"`
	Port     string `json:"port" env:"HTTP_CLIENT_PORT" env-default:"8080"`
}
