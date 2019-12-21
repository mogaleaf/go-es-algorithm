package evolution

import (
	"go-es-algorithm/evolution/genes"
	"go-es-algorithm/recorder"
	"math/rand"
	"sort"
	"time"
)

type ES struct {
	CalcFitness        func([]float64) float64
	Win                func([]float64) bool
	ValuesSize         int
	RangeInit          []float64
	NumberIterationMax int
	Type               GeneType
	PopulationSize     int
	OffSpringSize      int
	ParentsNumber      int
	SelectionType      SelectionType
	Recorder           recorder.Recorder
}

func (e *ES) Evolve() (genes.GeneI, int) {
	population := e.generatePopulation()

	for _, individual := range population {
		if e.Win(individual.GetValues()) {
			return individual, 0
		}
	}

	for i := 0; i < e.NumberIterationMax; i++ {

		children := e.createChildren(population)
		if e.Recorder != nil {
			e.Recorder.CreatedOffspring(children, i)
		}
		mutateParams(children)
		mutate(children)

		if e.Recorder != nil {
			e.Recorder.MutatedOffspring(children, i)
		}

		population = e.selectNextGeneration(children, population)

		if e.Recorder != nil {
			e.Recorder.NextGeneration(population, i)
		}

		for _, individual := range population {
			if e.Win(individual.GetValues()) {
				return individual, i
			}
		}
	}

	//todo select the best ?
	return population[0], e.NumberIterationMax
}

func (e *ES) selectNextGeneration(children []genes.GeneI, population []genes.GeneI) []genes.GeneI {
	var parentsAndOffSpring []struct {
		pop         genes.GeneI
		CalcFitness func(genes.GeneI) float64
	}
	for _, in := range children {
		parentsAndOffSpring = append(parentsAndOffSpring,
			struct {
				pop         genes.GeneI
				CalcFitness func(genes.GeneI) float64
			}{pop: in, CalcFitness: e.CalculateAndSaveFitness},
		)
	}
	switch e.SelectionType {
	case MuPlusLambda:
		for _, in := range population {
			parentsAndOffSpring = append(parentsAndOffSpring,
				struct {
					pop         genes.GeneI
					CalcFitness func(genes.GeneI) float64
				}{pop: in, CalcFitness: e.CalculateAndSaveFitness},
			)
		}
	case MuCommaLambda:
	}
	var newGen []genes.GeneI
	sort.Sort(ByFitness(parentsAndOffSpring))

	for i := 0; i < len(population); i++ {
		newGen = append(newGen, parentsAndOffSpring[i].pop)
	}
	rand.Shuffle(len(newGen), func(i, j int) { newGen[i], newGen[j] = newGen[j], newGen[i] })
	return newGen

}

func (e *ES) generatePopulation() []genes.GeneI {
	var population []genes.GeneI
	switch e.Type {
	case Correlated_mutation:
		for i := 0; i < e.PopulationSize; i++ {
			population = append(population, genes.NewGeneCorrelatedMutation(e.ValuesSize, e.RangeInit))
		}
	case N_Step_mutation:
		for i := 0; i < e.PopulationSize; i++ {
			population = append(population, genes.NewGeneUnCorrelatedNStepMutation(e.ValuesSize, e.RangeInit))
		}
	case One_Step_mutation:
		for i := 0; i < e.PopulationSize; i++ {
			population = append(population, genes.NewGeneUnCorrelatedOneStepMutation(e.ValuesSize, e.RangeInit))
		}
	}

	return population
}

func (e *ES) createChildren(population []genes.GeneI) []genes.GeneI {
	rand.Seed(time.Now().UnixNano())
	var children []genes.GeneI
	for i := 0; i < e.OffSpringSize; i++ {
		position1 := rand.Intn(len(population))
		parents1 := population[position1]
		var others []genes.GeneI
		for j := 0; j < e.ParentsNumber; j++ {
			position2 := rand.Intn(len(population))
			parents2 := population[position2]
			others = append(others, parents2)
		}
		children = append(children, parents1.Recombine(others...))
	}
	return children

}

func (e *ES) CalculateAndSaveFitness(g genes.GeneI) float64 {
	if g.IsFitnessAvailable() {
		return g.GetCalculatedFitness()
	}
	fitness := e.CalcFitness(g.GetValues())
	g.SaveFitness(fitness)
	return fitness
}

func mutateParams(children []genes.GeneI) {
	for _, individual := range children {
		individual.MutateParams()
	}
}

func mutate(children []genes.GeneI) {
	for _, individual := range children {
		individual.Mutate()
	}
}
