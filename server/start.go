package server

import (
	"log"

	"github.com/valyala/fasthttp"
)

func (s *Server) Start() {
	log.Printf("starting server at " + s.Address)
	if err := fasthttp.ListenAndServe(s.Address, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
