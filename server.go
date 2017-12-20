package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Server struct {
	Config    Config
	Processor *Processor
}

func NewServer(config Config) (*Server, error) {
	processor, err := NewProcessor(config)
	if err != nil {
		return nil, errors.Wrap(err, "[Server]: failed to create new server")
	}
	return &Server{Config: config, Processor: processor}, nil
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/api/instances", makeHandler(apiReadInstances, s.Processor)).Methods("GET")

	// SITE
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./site/")))

	srv := http.Server{
		Addr:         "0.0.0.0:8181",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// update instances periodically
	go func(p *Processor) {
		ticker := time.NewTicker(5 * time.Minute)
		for t := range ticker.C {
			log.Println("upating instances", t)
			err := p.UpdateInstances()
			if err != nil {
				log.Println(err)
			}
			log.Println("done upating instances")
		}
	}(s.Processor)

	err := srv.ListenAndServe()
	return errors.Wrap(err, "[Server]: server error")
}
