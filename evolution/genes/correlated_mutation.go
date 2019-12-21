package genes

func NewGeneCorrelatedMutation(valueSize int, rangeInit []float64) GeneI {
	return &GeneCorrelatedMutation{}
}

type GeneCorrelatedMutation struct {
	Values           []float64
	Sigma            []float64
	Alpha            []float64
	ValueFitness     float64
	ValueFitnessCalc bool
}

func (gene *GeneCorrelatedMutation) Mutate() {

}

func (gene *GeneCorrelatedMutation) MutateParams() {

}

func (gene *GeneCorrelatedMutation) Recombine(other ...GeneI) GeneI {
	return nil
}

func (gene *GeneCorrelatedMutation) GetValues() []float64 {
	return gene.Values
}

func (gene *GeneCorrelatedMutation) GetSigma() []float64 {
	return gene.Sigma
}

func (gene *GeneCorrelatedMutation) IsFitnessAvailable() bool {
	return gene.ValueFitnessCalc
}
func (gene *GeneCorrelatedMutation) GetCalculatedFitness() float64 {
	return gene.ValueFitness
}
func (gene *GeneCorrelatedMutation) SaveFitness(v float64) {
	gene.ValueFitness = v
	gene.ValueFitnessCalc = true
}
