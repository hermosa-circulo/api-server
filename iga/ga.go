package iga

import (
	"fmt"
	"math/rand"
)

func crossover(gene1 *Gene, gene2 *Gene) *Gene {

	crossPoint1 := randInt(0, len(gene1.wheelRadius))
	crossPoint2 := randInt(0, len(gene1.begin))
	crossPoint3 := randInt(0, len(gene1.pointNum))
	crossPoint4 := randInt(0, len(gene1.breastWide))
	newPara1 := gene1.wheelRadius[:crossPoint1] + gene2.wheelRadius[crossPoint1:]
	newPara2 := gene1.begin[:crossPoint2] + gene2.begin[crossPoint2:]
	newPara3 := gene1.pointNum[:crossPoint3] + gene2.pointNum[crossPoint3:]
	newPara4 := gene1.breastWide[:crossPoint4] + gene2.breastWide[crossPoint4:]

	return &Gene{
		wheelRadius: newPara1,
		begin:       newPara2,
		pointNum:    newPara3,
		breastWide:  newPara4,
	}
}

func mutation(para string) string {
	var result string
	for index, _ := range para {
		if rand.Float64() < MUTATION_RATE {
			result += string(invertBit(para[index]))
		} else {
			result += string(para[index])
		}
	}
	return result
}

func invertBit(bit byte) byte {
	if bit == ASCII_0 {
		return ASCII_1
	} else if bit == ASCII_1 {
		return ASCII_0
	} else {
		fmt.Println("Unknown bit pattern.", bit)
		return ASCII_0
	}

}
