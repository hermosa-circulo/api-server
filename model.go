package main

import (
	"encoding/json"
	"log"
)

type Evaluation struct {
	Value    int
	UserName string
}

type Gene struct {
	id          int
	Evaluations []Evaluation
	WheelRadius int
	Begin       int
	PointNum    int
	BreastWide  int
}

func (g Gene) CalcEvaluation() int {
	return 1
}

func (g Gene) ToString() string {
	str_byte, err := json.Marshal(g)
	if err != nil {
		log.Fatal("JSON marshal error:", err)
	}

	return string(str_byte)
}

func NewGene(str []byte) *Gene {

	data := new(Gene)
	if err := json.Unmarshal(str, data); err != nil {
		log.Fatal("JSON Unmarshal error:", err)
	}

	return data
}

func NewEvaluation(value int, user string) *Evaluation {
	return &Evaluation{Value: value, UserName: user}
}
