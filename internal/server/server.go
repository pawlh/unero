package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	return &Server{Port: port}
}

func (s *Server) Start() {
	router := gin.Default()

	registerRoutes(router)

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	if err := router.Run(fmt.Sprintf("0.0.0.0:%d", s.Port)); err != nil {
		panic(err)
	}
}
