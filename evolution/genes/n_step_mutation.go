package genes

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func NewGeneUnCorrelatedNStepMutation(valueSize int, rangeInit []float64) GeneI {
	mutation := GeneUnCorrelatedNStepMutation{}
	mutation.Theta1 = 1.0 / math.Sqrt(2.1*float64(valueSize))
	mutation.Theta2 = 1.0 / math.Sqrt(1.9*float64(valueSize))
	mutation.Values = generateValues(valueSize, rangeInit)
	mutation.Sigma = generateValues(valueSize, []float64{0, 5})
	return &mutation
}

type GeneUnCorrelatedNStepMutation struct {
	Values           []float64
	Sigma            []float64
	Theta1           float64
	Theta2           float64
	ValueFitness     float64
	ValueFitnessCalc bool
}

func (gene *GeneUnCorrelatedNStepMutation) Mutate() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(gene.Values); i++ {
		dist := distuv.Normal{
			Mu:    0,             // Mean of the normal distribution
			Sigma: gene.Sigma[i], // Standard deviation of the normal distribution
		}
		f := dist.Rand()

		gene.Values[i] += f
	}
}

func (gene *GeneUnCorrelatedNStepMutation) MutateParams() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(gene.Sigma); i++ {
		dist := distuv.Normal{
			Mu:    0, // Mean of the normal distribution
			Sigma: 1, // Standard deviation of the normal distribution
		}
		distT1 := distuv.Normal{
			Mu:    0,           // Mean of the normal distribution
			Sigma: gene.Theta1, // Standard deviation of the normal distribution
		}
		distT2 := distuv.Normal{
			Mu:    0,           // Mean of the normal distribution
			Sigma: gene.Theta2, // Standard deviation of the normal distribution
		}
		t1 := distT1.Rand()
		t2 := distT2.Rand()
		exp := math.Exp(t1*dist.Rand() + t2*dist.Rand())
		newSigma := gene.Sigma[i] * exp
		if newSigma < 0.00000000002 {
			newSigma = 0.00000000002
		}

		gene.Sigma[i] = newSigma
	}

}

func (gene *GeneUnCorrelatedNStepMutation) Recombine(other ...GeneI) GeneI {
	rand.Seed(time.Now().UnixNano())
	newValues := make([]float64, len(gene.Values))
	for i, v := range gene.Values {
		newValues[i] = v
	}
	for _, p := range other {
		for i := range newValues {
			newValues[i] += p.(*GeneUnCorrelatedNStepMutation).Values[i]
		}
	}
	for i := range newValues {
		newValues[i] /= float64(len(other) + 1)
	}

	newSigmas := make([]float64, len(gene.Sigma))
	for i := range newSigmas {
		intn := rand.Intn(len(other) + 1)
		if intn < len(other) {
			newSigmas[i] = other[intn].(*GeneUnCorrelatedNStepMutation).Sigma[i]
		} else {
			newSigmas[i] = gene.Sigma[i]
		}
	}
	return &GeneUnCorrelatedNStepMutation{
		Values: newValues,
		Theta1: gene.Theta1,
		Theta2: gene.Theta2,
		Sigma:  newSigmas,
	}
}

func (gene *GeneUnCorrelatedNStepMutation) GetValues() []float64 {
	return gene.Values
}

func (gene *GeneUnCorrelatedNStepMutation) GetSigma() []float64 {
	return gene.Sigma
}

func (gene *GeneUnCorrelatedNStepMutation) IsFitnessAvailable() bool {
	return gene.ValueFitnessCalc
}
func (gene *GeneUnCorrelatedNStepMutation) GetCalculatedFitness() float64 {
	return gene.ValueFitness
}
func (gene *GeneUnCorrelatedNStepMutation) SaveFitness(v float64) {
	gene.ValueFitness = v
	gene.ValueFitnessCalc = true
}
