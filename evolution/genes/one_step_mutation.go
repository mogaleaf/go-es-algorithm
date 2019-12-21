package genes

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func NewGeneUnCorrelatedOneStepMutation(valueSize int, rangeInit []float64) GeneI {
	one := GeneUnCorrelatedOneStepMutation{
		Sigma: 1.0,
	}
	one.Theta = 1.0 / math.Sqrt(float64(valueSize))
	one.Values = generateValues(valueSize, rangeInit)
	return &one
}

type GeneUnCorrelatedOneStepMutation struct {
	Values           []float64
	Sigma            float64
	Theta            float64
	ValueFitness     float64
	ValueFitnessCalc bool
}

func (gene *GeneUnCorrelatedOneStepMutation) Mutate() {
	dist := distuv.Normal{
		Mu:    0,          // Mean of the normal distribution
		Sigma: gene.Sigma, // Standard deviation of the normal distribution
	}

	for i := 0; i < len(gene.Values); i++ {
		gene.Values[i] += dist.Rand()
	}
}

func (gene *GeneUnCorrelatedOneStepMutation) MutateParams() {
	dist := distuv.Normal{
		Mu:    0, // Mean of the normal distribution
		Sigma: 1, // Standard deviation of the normal distribution
	}
	distT := distuv.Normal{
		Mu:    0,          // Mean of the normal distribution
		Sigma: gene.Theta, // Standard deviation of the normal distribution
	}
	t := distT.Rand()
	newSigma := gene.Sigma * math.Exp(t*dist.Rand())
	if newSigma < 0.0002 {
		newSigma = 0.0002
	}
	gene.Sigma = newSigma
}

func (gene *GeneUnCorrelatedOneStepMutation) Recombine(other ...GeneI) GeneI {
	newValues := make([]float64, len(gene.Values))
	newSigma := gene.Sigma
	for i, v := range gene.Values {
		newValues[i] = v
	}
	for _, p := range other {
		for i := range newValues {
			newValues[i] += p.(*GeneUnCorrelatedOneStepMutation).Values[i]

		}
		newSigma += p.(*GeneUnCorrelatedOneStepMutation).Sigma
	}
	for i := range newValues {
		newValues[i] /= float64(len(other) + 1)
	}
	newSigma /= float64(len(other) + 1)
	return &GeneUnCorrelatedOneStepMutation{
		Values: newValues,
		Sigma:  gene.Sigma,
		Theta:  gene.Theta,
	}
}

func (gene *GeneUnCorrelatedOneStepMutation) GetValues() []float64 {
	return gene.Values
}

func (gene *GeneUnCorrelatedOneStepMutation) GetSigma() []float64 {
	return []float64{gene.Sigma}
}

func (gene *GeneUnCorrelatedOneStepMutation) IsFitnessAvailable() bool {
	return gene.ValueFitnessCalc
}
func (gene *GeneUnCorrelatedOneStepMutation) GetCalculatedFitness() float64 {
	return gene.ValueFitness
}
func (gene *GeneUnCorrelatedOneStepMutation) SaveFitness(v float64) {
	gene.ValueFitness = v
	gene.ValueFitnessCalc = true
}
