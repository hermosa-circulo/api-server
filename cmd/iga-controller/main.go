package main

import (
	"log"
	"net"

	"github.com/hermosa-circulo/iga-controller/api"
	"github.com/hermosa-circulo/iga-controller/iga"
	"google.golang.org/grpc"
)

// https://www.aaai.org/Papers/Symposia/Spring/2008/SS-08-03/SS08-03-006.pdf

func main() {
	l, err := net.Listen("tcp", ":13009")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	api.RegisterGenesServer(s, &iga.GenesService{})
	api.RegisterEvaluationsServer(s, &iga.EvaluationsService{})
	s.Serve(l)
}
