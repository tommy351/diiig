package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tommy351/diiig/dao/daomock"
	"github.com/tommy351/diiig/model"
)

func TestServer_Home(t *testing.T) {
	var topicDAO *daomock.TopicDAO

	testReq := func() *httptest.ResponseRecorder {
		s := &Server{
			TopicDAO: topicDAO,
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		s.router(gin.New()).ServeHTTP(w, req)

		return w
	}

	t.Run("Success", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		topics := []model.Topic{
			{Name: "a", Score: 1},
			{Name: "b", Score: 0},
		}
		topicDAO.On("Range", 0, topTopicSize-1).
			Return(topics, nil).
			Once()

		w := testReq()
		assert.Equal(t, http.StatusOK, w.Code)
		topicDAO.AssertExpectations(t)
	})

	t.Run("DAO error", func(t *testing.T) {
		topicDAO := new(daomock.TopicDAO)
		topicDAO.On("Range", 0, topTopicSize-1).
			Return(nil, errors.New("random error")).
			Once()

		assert.Panics(t, func() {
			testReq()
		})
	})
}
