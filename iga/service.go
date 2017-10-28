package iga

type GeneService struct {
	r GeneRepository
}

func NewGeneService() *GeneService {
	return &GeneService{
		GetGeneRepository(),
	}
}

func (s *GeneService) GetAllGene() []*Gene {
	return s.r.GetAllGene()
}

func (s *GeneService) GetGeneById(id int) (*Gene, error) {
	return s.r.GetGeneById(id)
}

func (s *GeneService) SaveGene(gene *Gene) error {
	return s.r.SaveGene(gene)
}

var s *GeneService

//getter
func GetGeneService() *GeneService {
	if s == nil {
		s = NewGeneService()
	}
	return s
}
