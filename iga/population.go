package iga

import (
	"fmt"
)

type Population [POPULATION_NUM]*Gene

func NewPopulation() *Population {
	p := &Population{}
	for key, _ := range p {
		p[key] = newGene()
	}
	return p
}

func (p *Population) Printf() {
	for _, gene := range p {
		fmt.Println(gene.toString())
	}
}
func (p *Population) Len() int {
	return len(p)
}

func (p *Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Population) Less(i, j int) bool {
	return p[i].evaluate() < p[j].evaluate()
}

func (p *Population) Crossover() {
	var i, j int
	boundry := int(SELECT_RATE * float64(len(p)))
	for key, _ := range p[:boundry] {
		for {
			i = randInt(boundry, len(p))
			j = randInt(boundry, len(p))
			if i != j {
				break
			}
		}
		p[key] = crossover(p[i], p[j])
	}
}

func (p *Population) Mutation() {
	for key, _ := range p {
		p[key].wheelRadius = mutation(p[key].wheelRadius)
		p[key].begin = mutation(p[key].begin)
		p[key].pointNum = mutation(p[key].pointNum)
		p[key].breastWide = mutation(p[key].breastWide)
	}

}
