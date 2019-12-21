package evolution

import "go-es-algorithm/evolution/genes"

type ByFitness []struct {
	pop         genes.GeneI
	CalcFitness func(genes.GeneI) float64
}

func (a ByFitness) Len() int      { return len(a) }
func (a ByFitness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool {
	return a[i].CalcFitness(a[i].pop) < a[j].CalcFitness(a[j].pop)
}
