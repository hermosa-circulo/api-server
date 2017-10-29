package iga

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"log"
	"strconv"
)

const (
	directory = "gene"
	etcdPort  = ":2379"
	etcdHost  = "127.0.0.1"
)

type GeneRepository interface {
	GetAllGene() []Gene
	GetGeneById(id int) (*Gene, error)
	SaveGene(gene *Gene) error
}

type DBGeneRepository struct {
}

var etcd client.KeysAPI

func NewGeneRepository() *DBGeneRepository {
	cfg := client.Config{
		Endpoints: []string{"http://" + etcdHost + etcdPort},
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)
	kAPI := client.NewKeysAPI(c)

	if err != nil {
		log.Fatal(err)
	}

	o := client.SetOptions{Dir: true}
	_, err = kAPI.Set(context.Background(), "/"+directory, "", &o)

	if err != nil {
		log.Fatal(err)
	}

	etcd = kAPI

	return &DBGeneRepository{}
}

func (r *DBGeneRepository) GetAllGene() []Gene {

	resp, err := etcd.Get(context.Background(), "/"+directory, nil)

	if err != nil {
		log.Fatal(err)
	}

	for _, n := range resp.Node.Nodes {
		fmt.Printf("Key: %q, Value: %q\n", n.Key, n.Value)
	}

	return []Gene{}
}

func (r *DBGeneRepository) GetGeneById(id int) (*Gene, error) {

	resp, err := etcd.Get(context.Background(), joinPath(id), nil)

	if err != nil {
		log.Fatal(err)
	}
	return NewGene([]byte(resp.Node.Value)), nil
}

func (r *DBGeneRepository) SaveGene(gene *Gene) error {

	_, err := etcd.Create(context.Background(), joinPath(gene.id), gene.ToString())

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *DBGeneRepository) DeleteGene(id int) error {

	_, err := etcd.Delete(context.Background(), joinPath(id), &client.DeleteOptions{})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

var r *DBGeneRepository

func GetGeneRepository() (r GeneRepository) {
	if r == nil {
		r = NewGeneRepository()
	}
	return r
}

func joinPath(id int) string {
	return "/" + directory + "/" + strconv.Itoa(id)
}
