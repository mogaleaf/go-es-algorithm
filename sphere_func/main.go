package main

import (
	"fmt"
	"go-es-algorithm/evolution"
	"go-es-algorithm/evolution/genes"
)

func main() {
	//withOneToOne()
	withOneToN()

}

func withOneToN() {
	es := evolution.NewES(sphere, winSphere,
		evolution.WithNumberIterationMax(200000),
		evolution.WithOffSpringSize(700),
		evolution.WithParentsNumber(2),
		evolution.WithRangeInit([]float64{-30.0, 30.0}),
		evolution.WithSelectionType(evolution.MuCommaLambda),
		evolution.WithType(evolution.N_Step_mutation),
		evolution.WithValuesSize(7),
		evolution.WithPopulationSize(100),
	)

	oneOne, i := es.Evolve()

	for _, v := range oneOne.(*genes.GeneUnCorrelatedNStepMutation).Values {
		println(fmt.Sprintf("%0.10f", v))
	}
	println()
	println()
	println(fmt.Sprintf("%0.10f", sphere(oneOne.GetValues())))
	println(fmt.Sprintf("%d", i))
}

func withOneToOne() {
	oneOne, i := evolution.EvolveOneToOne(sphereOne, winsphereOne, 3, []float64{-30.0, 30.0}, 10, 0.9, 200000)

	for _, v := range oneOne.Values {
		println(fmt.Sprintf("%0.10f", v))
	}
	println()
	println()
	println(fmt.Sprintf("%0.10f", sphereOne(oneOne)))
	println(fmt.Sprintf("%d", i))
}

func sphere(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v * v
	}
	return sum
}

func winSphere(values []float64) bool {
	return sphere(values) < 0.00000000002
}

func sphereOne(g *genes.GeneOneOne) float64 {
	sum := 0.0
	for _, v := range g.Values {
		sum += (v * v)
	}
	return sum
}

func winsphereOne(g *genes.GeneOneOne) bool {
	return sphereOne(g) == 0.0
}
