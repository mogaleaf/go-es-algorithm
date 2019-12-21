package recorder

import "go-es-algorithm/evolution/genes"

type Recorder interface {
	CreatedOffspring(gs []genes.GeneI, iter int)
	MutatedOffspring(gs []genes.GeneI, iter int)
	NextGeneration(gs []genes.GeneI, iter int)
}
