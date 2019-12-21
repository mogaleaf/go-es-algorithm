package genes

import (
	"gonum.org/v1/gonum/stat/distuv"
)

func NewGeneOneOne(size int, rangeInit []float64) *GeneOneOne {
	one := GeneOneOne{}
	one.Values = generateValues(size, rangeInit)
	return &one
}

type GeneOneOne struct {
	Values []float64
}

func (gene *GeneOneOne) Mutate(sigma float64) *GeneOneOne {
	mutated := GeneOneOne{
		Values: append([]float64{}, gene.Values...),
	}
	dist := distuv.Normal{
		Mu:    0,     // Mean of the normal distribution
		Sigma: sigma, // Standard deviation of the normal distribution
	}
	for i := 0; i < len(gene.Values); i++ {
		mutated.Values[i] += dist.Rand()
	}
	return &mutated
}
