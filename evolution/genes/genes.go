package genes

import (
	"math/rand"
	"time"
)

type GeneI interface {
	Mutate()
	MutateParams()
	Recombine(other ...GeneI) GeneI
	GetValues() []float64
	GetSigma() []float64
	IsFitnessAvailable() bool
	GetCalculatedFitness() float64
	SaveFitness(float64)
}

func generateValues(size int, rangeInit []float64) []float64 {
	var values []float64
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		i2 := rand.Float64() * (rangeInit[1] - rangeInit[0])
		values = append(values, i2)
	}
	return values
}
