package evolution

import (
	"go-es-algorithm/evolution/genes"
)

func EvolveOneToOne(CalcFitness func(*genes.GeneOneOne) float64, Win func(*genes.GeneOneOne) bool, size int, rangeInit []float64, k int, c float64, NumberIterationMax int) (*genes.GeneOneOne, int) {
	sigma := 1.0
	gs := 0
	individual := genes.NewGeneOneOne(size, rangeInit)
	if Win(individual) {
		return individual, 0
	}
	for i := 0; i < NumberIterationMax; i++ {
		mutate := individual.Mutate(sigma)

		if CalcFitness(mutate) < CalcFitness(individual) {
			gs++
			individual = mutate
		}
		//Rechenberg 1/5 rule
		if (i % k) == 0 {
			ps := float64(gs) / float64(k)
			r := 1.0 / 5.0
			if ps > r {
				sigma /= c
			} else if ps < r {
				sigma *= c
			}
			gs = 0
		}

		if Win(individual) {
			return individual, i
		}
	}
	return individual, NumberIterationMax
}
