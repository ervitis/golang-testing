package server

import (
	"fmt"
)

type Server struct {
	Port string
	Addr string
}

type UrlIface interface {
	Url() string
	FullUrl(...string) string
	PortAddress() string
}

func (s *Server) Url() string {
	return fmt.Sprintf("%s:%s", s.Addr, s.Port)
}

func (s *Server) PortAddress() string {
	return fmt.Sprintf(":%s", s.Port)
}

func (s *Server) FullUrl(endpointResource ...string) string {
	if endpointResource != nil {
		return fmt.Sprintf("%s/%s", s.Url(), endpointResource)
	}
	return s.Url()
}
