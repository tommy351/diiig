package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tommy351/diiig/dao"
)

// Server represents an instance of a server.
type Server struct {
	Host string `long:"host" env:"HOST" description:"the host of the server"`
	Port int    `long:"port" env:"PORT" default:"4000" description:"the port of the server"`

	TopicDAO dao.TopicDAO
}

// Serve starts the server.
func (s *Server) Serve() error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return s.router(gin.Default()).Run(addr)
}

func (s *Server) router(r *gin.Engine) *gin.Engine {
	r.SetHTMLTemplate(homeTemplate)

	r.GET("/", s.Home)
	r.POST("/topics", s.CreateTopic)
	r.POST("/vote", s.VoteTopic)

	return r
}
