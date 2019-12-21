package evolution

import "go-es-algorithm/recorder"

type GeneType uint8

type EvolveOpts func(es *ES)

func WithRecorder(r recorder.Recorder) EvolveOpts {
	return func(e *ES) {
		e.Recorder = r
	}
}

func WithPopulationSize(populationSize int) EvolveOpts {
	return func(e *ES) {
		e.PopulationSize = populationSize
	}
}

func WithValuesSize(valueSize int) EvolveOpts {
	return func(e *ES) {
		e.ValuesSize = valueSize
	}
}

func WithRangeInit(r []float64) EvolveOpts {
	return func(e *ES) {
		e.RangeInit = r
	}
}

func WithNumberIterationMax(it int) EvolveOpts {
	return func(e *ES) {
		e.NumberIterationMax = it
	}
}

func WithType(t GeneType) EvolveOpts {
	return func(e *ES) {
		e.Type = t
	}
}

func WithOffSpringSize(s int) EvolveOpts {
	return func(e *ES) {
		e.OffSpringSize = s
	}
}

func WithParentsNumber(p int) EvolveOpts {
	return func(e *ES) {
		e.ParentsNumber = p
	}
}

func WithSelectionType(selection SelectionType) EvolveOpts {
	return func(e *ES) {
		e.SelectionType = selection
	}
}

const (
	One_Step_mutation = iota
	N_Step_mutation
	Correlated_mutation
)

type SelectionType uint8

const (
	MuCommaLambda SelectionType = iota
	MuPlusLambda
)

func NewES(CalcFitness func([]float64) float64, Win func([]float64) bool, options ...EvolveOpts) *ES {
	e := &ES{
		CalcFitness:        CalcFitness,
		Win:                Win,
		NumberIterationMax: 100,
		PopulationSize:     100,
		ParentsNumber:      2,
		SelectionType:      MuCommaLambda,
		OffSpringSize:      700,
		Type:               N_Step_mutation,
		RangeInit:          []float64{-10, 10},
		ValuesSize:         10,
	}
	for _, option := range options {
		option(e)
	}
	return e
}
