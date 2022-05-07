package api

type Server struct {
	config *Config
}

func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (a *Server) Start() error {
	return nil
}
