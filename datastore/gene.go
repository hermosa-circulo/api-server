package datastore

import (
	"github.com/hermosa-circulo/iga-controller/api"
)

var Gene *GeneStore

func init() {
	Gene := NewGeneStore()
}

type GeneStore struct {
}

func (s *GeneStore) Get(id string) (*api.Gene, error) {

	return &api.Gene{}, nil
}

func (s *GeneStore) Set(g api.Gene) error {
	return nil
}

func NewGeneStore() *GeneStore {
	return &GeneStore{}
}
