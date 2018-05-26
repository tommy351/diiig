package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tommy351/diiig/dao"
)

type Server struct {
	Host string `long:"host" env:"HOST" description:"the host of the server"`
	Port int    `long:"port" env:"PORT" default:"4000" description:"the port of the server"`

	TopicDAO dao.TopicDAO
}

func (s *Server) Serve() error {
	r := gin.Default()
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	r.LoadHTMLGlob("templates/*")

	r.GET("/", s.Home)
	r.POST("/topics", s.CreateTopic)
	r.POST("/vote", s.VoteTopic)

	return r.Run(addr)
}
