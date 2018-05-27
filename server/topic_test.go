package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tommy351/diiig/dao/daomock"
)

func postForm(url string, values url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func assertRedirect(t *testing.T, w *httptest.ResponseRecorder, code int, location string) {
	assert.Equal(t, code, w.Code)
	assert.Equal(t, location, w.HeaderMap.Get("Location"))
}

func repeat(s string, times int) string {
	result := ""

	for i := 0; i < times; i++ {
		result += s
	}

	return result
}

func TestServer_CreateTopic(t *testing.T) {
	var topicDAO *daomock.TopicDAO

	testReq := func(values url.Values) *httptest.ResponseRecorder {
		s := &Server{
			TopicDAO: topicDAO,
		}
		w := httptest.NewRecorder()
		req, _ := postForm("/topics", values)

		s.router(gin.New()).ServeHTTP(w, req)
		return w
	}

	t.Run("Success", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		topicDAO.On("Create", "foo").Return(nil).Once()
		w := testReq(url.Values{
			"topic": {"foo"},
		})
		assertRedirect(t, w, http.StatusFound, "/")
		topicDAO.AssertExpectations(t)
	})

	t.Run("Topic is required", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		w := testReq(url.Values{})
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Topic exceeded max length", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		w := testReq(url.Values{
			"topics": {repeat("a", 256)},
		})
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("DAO error", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		topicDAO.On("Create", "foo").
			Return(errors.New("random error")).
			Once()

		assert.Panics(t, func() {
			testReq(url.Values{
				"topic": {"foo"},
			})
		})
	})
}

func TestServer_VoteTopic(t *testing.T) {
	var topicDAO *daomock.TopicDAO

	testReq := func(values url.Values) *httptest.ResponseRecorder {
		s := &Server{
			TopicDAO: topicDAO,
		}
		w := httptest.NewRecorder()
		req, _ := postForm("/vote", values)

		s.router(gin.New()).ServeHTTP(w, req)
		return w
	}

	t.Run("Success", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		topicDAO.On("Vote", "foo", 1).Return(nil).Once()
		w := testReq(url.Values{
			"topic": {"foo"},
			"score": {"1"},
		})
		assertRedirect(t, w, http.StatusFound, "/")
		topicDAO.AssertExpectations(t)
	})

	t.Run("Topic is required", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		w := testReq(url.Values{})
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Score is required", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		w := testReq(url.Values{
			"topic": {"foo"},
		})
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("DAO error", func(t *testing.T) {
		topicDAO = new(daomock.TopicDAO)
		topicDAO.On("Vote", "foo", 1).
			Return(errors.New("random error")).
			Once()

		assert.Panics(t, func() {
			testReq(url.Values{
				"topic": {"foo"},
				"score": {"1"},
			})
		})
	})
}
