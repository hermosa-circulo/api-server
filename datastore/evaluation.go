package datastore

import (
	"github.com/hermosa-circulo/iga-controller/api"
)

var Evaluation *EvaluationStore

func init() {
	Evaluation := NewEvaluationStore()
}

type EvaluationStore struct {
}

func (s *EvaluationStore) Get(id string) (*api.Evaluation, error) {
	return &api.Evaluation{}, nil
}

func (s *EvaluationStore) Set(g api.Evaluation) error {

	return nil
}

func NewEvaluationStore() *EvaluationStore {
	return &EvaluationStore{}
}
