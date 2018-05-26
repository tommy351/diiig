package server

import "github.com/tommy351/diiig/dao"

type Server struct {
	Host string
	Port int

	TopicDAO dao.TopicDAO
}
