package hctl

import (
	"log"

	"google.golang.org/grpc"
)

func NewConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:13009", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	return conn
}
