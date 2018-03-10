package hctl

import (
	"fmt"
	"log"

	"github.com/hermosa-circulo/iga-controller/api"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func GetEvaluation(c *cli.Context) error {

	client := api.NewEvaluationsClient(NewConnection())

	req := &api.GetEvaluationRequest{}
	res, err := client.Get(context.Background(), req)
	if err != nil {
		log.Fatalln("Get:", err)
	}
	fmt.Println(res)
	fmt.Println("GetEvaluationRequest success!!!")
	return nil
}

func SetEvaluation(c *cli.Context) error {

	client := api.NewEvaluationsClient(NewConnection())

	req := &api.SetEvaluationRequest{}
	res, err := client.Set(context.Background(), req)
	if err != nil {
		log.Fatalln("Send:", err)
	}
	fmt.Println(res)
	fmt.Println("SetEvaluationRequest success!!!")
	return nil
}
