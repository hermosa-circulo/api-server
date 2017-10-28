package iga

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

//mock
func (g Gene) CalcEvaluation() int {
	return 1
}

func NewGene() *Gene {
	return &Gene{}
}

func NewEvaluation(value int, user string) *Evaluation {
	return &Evaluation{value, user}
}

//mock
func Merge(gene1 Gene, gene2 Gene) Gene {
	return gene1
}
