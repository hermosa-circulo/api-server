package iga

type GeneRepository interface {
	GetAllGene() []*Gene
	GetGeneById(id int) (*Gene, error)
	SaveGene(gene *Gene) error
}

type DBGeneRepository struct {
}

func NewGeneRepository() *DBGeneRepository {
	return &DBGeneRepository{}
}

func (r *DBGeneRepository) GetAllGene() []*Gene {
	return nil
}

func (r *DBGeneRepository) GetGeneById(id int) (*Gene, error) {
	return &Gene{}, nil
}

func (r *DBGeneRepository) SaveGene(gene *Gene) error {
	return nil
}

var r *DBGeneRepository

func GetGeneRepository() (r GeneRepository) {
	if r == nil {
		r = NewGeneRepository()
	}
	return r
}
