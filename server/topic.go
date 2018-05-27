package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createTopicForm struct {
	Topic string `form:"topic" binding:"required,lte=255"`
}

type voteTopicForm struct {
	Topic string `form:"topic" binding:"required"`
	Score int    `form:"score" binding:"required"`
}

// CreateTopic creates a new topic and redirects to the home page.
func (s *Server) CreateTopic(c *gin.Context) {
	var form createTopicForm

	if err := c.Bind(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.TopicDAO.Create(form.Topic); err != nil {
		panic(err)
	}

	redirectToHome(c)
}

// VoteTopic adds the score of the topic and redirects to the home page.
func (s *Server) VoteTopic(c *gin.Context) {
	var form voteTopicForm

	if err := c.Bind(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.TopicDAO.Vote(form.Topic, form.Score); err != nil {
		panic(err)
	}

	redirectToHome(c)
}

func redirectToHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
