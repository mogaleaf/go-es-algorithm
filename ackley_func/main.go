package main

import (
	"fmt"
	"go-es-algorithm/evolution"
	"go-es-algorithm/evolution/genes"
	"math"
)

func main() {
	withOneToN()
}

func withOneToN() {
	es := evolution.NewES(calcAckley, winAckley,
		evolution.WithNumberIterationMax(200000),
		evolution.WithOffSpringSize(700),
		evolution.WithParentsNumber(2),
		evolution.WithRangeInit([]float64{-30.0, 30.0}),
		evolution.WithSelectionType(evolution.MuCommaLambda),
		evolution.WithType(evolution.N_Step_mutation),
		evolution.WithValuesSize(30),
		evolution.WithPopulationSize(100),
		evolution.WithRecorder(&Recorder{}),
	)

	oneOne, i := es.Evolve()

	for _, v := range oneOne.GetValues() {
		println(fmt.Sprintf("%0.10f", v))
	}
	println()
	println()
	println(fmt.Sprintf("%0.10f", calcAckley(oneOne.GetValues())))
	println(fmt.Sprintf("%d", i))
}

func calcAckley(value []float64) float64 {
	sum := 0.0
	sum2 := 0.0
	for _, v := range value {
		sum += v * v
		sum2 += math.Cos(2.0 * math.Pi * v)
	}
	sum *= 1.0 / float64(len(value))
	sum2 *= 1.0 / float64(len(value))
	sum = math.Sqrt(sum)
	sum *= -0.2
	sum = -20 * math.Exp(sum)
	sum2 = -1.0 * math.Exp(sum2)
	return sum + sum2 + 20 + math.E
}

func winAckley(g []float64) bool {
	return calcAckley(g) < 7.48*math.Pow(10.0, -8.0)
}

type Recorder struct {
}

func (*Recorder) CreatedOffspring(gs []genes.GeneI, iter int) {

}
func (*Recorder) MutatedOffspring(gs []genes.GeneI, iter int) {

}
func (*Recorder) NextGeneration(population []genes.GeneI, iter int) {
	if iter%10 == 0 {
		println()
		println("----------")
		println(fmt.Sprintf("%d", iter))
		mean := 0.0
		for _, individual := range population {
			mean += individual.GetCalculatedFitness()
		}
		println(fmt.Sprintf("%0.10f", mean/float64(len(population))))

	}
}
