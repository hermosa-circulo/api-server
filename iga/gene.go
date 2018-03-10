package iga

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Gene struct {
	id          int
	wheelRadius string
	begin       string
	pointNum    string
	breastWide  string
}

func newGene() *Gene {
	return &Gene{
		wheelRadius: NewGAPara(rand.Intn(WHEELRADIUS_MAX)),
		begin:       NewGAPara(rand.Intn(BEGIN_MAX)),
		pointNum:    NewGAPara(rand.Intn(POINTNUM_MAX)),
		breastWide:  NewGAPara(rand.Intn(BREASTWIDE_MAX)),
	}
}

func (gene *Gene) toString() string {
	return fmt.Sprint(gene.evaluate())
}

func (g *Gene) evaluate() int {
	result1 := toInt(g.wheelRadius) * WHEELRADIUS_MAX / PARA_MAX
	result2 := toInt(g.begin) * BEGIN_MAX / PARA_MAX
	result3 := toInt(g.pointNum) * POINTNUM_MAX / PARA_MAX
	result4 := toInt(g.breastWide) * BREASTWIDE_MAX / PARA_MAX
	return result1 + result2 + result3 + result4
}

//NewGAPara intを20桁の2進数文字列に変換
func NewGAPara(base int) string {
	var str string
	for i := 0; i < 10; i++ {
		if rand.Float64() < 0.5 {
			str += "1"
		} else {
			str += "0"
		}
	}

	return str
}

func toInt(para string) int {
	result, _ := strconv.ParseInt(para, 2, 0)
	return int(result)
}
