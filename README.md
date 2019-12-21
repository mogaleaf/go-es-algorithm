# go-es-algorithm

Small project to use Evolutionary strategy algorithm

```
es := evolution.NewES(
        // Fitness function = func([]float64)float64
        calcAckley, 
        // Is it a winner function = func([]float64)bool
        winAckley,
        //Iteration max number
		evolution.WithNumberIterationMax(200000),
		// Children (offspring) size
		evolution.WithOffSpringSize(700),
		// Parents recombination number
		evolution.WithParentsNumber(2),
		// Range of the xi values
		evolution.WithRangeInit([]float64{-30.0, 30.0}),
		// Type of selection ( u + l ) or (u,l)
		evolution.WithSelectionType(evolution.MuCommaLambda),
		// Type of mutation one/n/covariance
		evolution.WithType(evolution.N_Step_mutation),
		// xi number
		evolution.WithValuesSize(30),
		// u size
		evolution.WithPopulationSize(100),
	)
```
