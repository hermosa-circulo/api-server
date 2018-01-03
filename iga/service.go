package iga

import (
	"context"
	"fmt"
	"sort"

	"github.com/hermosa-circulo/iga-controller/api"
)

type GenesService struct{}

func (*GenesService) Get(ctx context.Context, r *api.GetGeneRequest) (*api.GetGeneResponse, error) {
	return &api.GetGeneResponse{
		Genes: []*api.Gene{
			&api.Gene{
				Id:          1,
				WheelRadius: 10,
				Begin:       12,
				PointNum:    1,
				BreastWide:  10,
			},
		},
	}, nil
}

type EvaluationsService struct{}

func (*EvaluationsService) Get(ctx context.Context, r *api.GetEvaluationRequest) (*api.GetEvaluationResponse, error) {
	return &api.GetEvaluationResponse{
		Evaluations: []*api.Evaluation{
			&api.Evaluation{
				Geneid: 1,
				Value:  1,
			},
		},
	}, nil
}

func (*EvaluationsService) Set(ctx context.Context, r *api.SetEvaluationRequest) (*api.SetEvaluationResponse, error) {
	fmt.Println("start")
	population := NewPopulation()
	for i := 0; i < GENERATION; i++ {
		population.Crossover()
		population.Mutation()
		sort.Sort(population)
	}
	population.Printf()
	return &api.SetEvaluationResponse{}, nil
}
