package hctl

import (
	"fmt"
	"log"

	"github.com/hermosa-circulo/iga-controller/api"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func GetGene(c *cli.Context) error {
	client := api.NewGenesClient(NewConnection())

	req := &api.GetGeneRequest{}
	res, err := client.Get(context.Background(), req)
	if err != nil {
		log.Fatalln("Get:", err)
	}
	fmt.Println(res)
	fmt.Println("GetGeneRequest succeess!!!")
	return nil
}
