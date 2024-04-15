package http

import (
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	SearchHandler *SearchHandler
	logger        *zap.Logger
}

func NewServer(searchHandler *SearchHandler, logger *zap.Logger) *Server {
	return &Server{SearchHandler: searchHandler, logger: logger}
}

func (s *Server) Start(addr string) error {
	s.logger.Info("HTTP server starting", zap.String("address", addr))

	http.Handle("/files/search", s.SearchHandler)

	return http.ListenAndServe(addr, nil)
}
