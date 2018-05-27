package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const topTopicSize = 20

// Home renders the home page which contains the top topics.
func (s *Server) Home(c *gin.Context) {
	topics, err := s.TopicDAO.Range(0, topTopicSize-1)

	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "home", gin.H{
		"Topics": topics,
	})
}
