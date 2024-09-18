package handler

import (
	"simpl-commerce/internal"
	repository "simpl-commerce/repository/user"
)

type Server struct {
	Cfg        internal.Config
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(cfg internal.Config, opts NewServerOptions) *Server {
	return &Server{
		Cfg:        cfg,
		Repository: opts.Repository,
	}
}
