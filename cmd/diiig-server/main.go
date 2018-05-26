package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/tommy351/diiig/dao"
	"github.com/tommy351/diiig/server"
)

func main() {
	s := new(server.Server)
	parser := flags.NewParser(s, flags.Default)

	if _, err := parser.Parse(); err != nil {
		code := 1

		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}

		os.Exit(code)
	}

	s.TopicDAO = dao.NewMemoryTopicDAO()

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
