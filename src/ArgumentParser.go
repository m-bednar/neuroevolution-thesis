package main

import (
	"flag"
	"log"
)

/* Arguments:
[name]				[type]		[flag]
env. file			string		-env
pop size			int			-pop
max generations		int			-maxg
min success rate 	float		-mins
steps				int			-steps
mutation strength	float		-mutstr
tournament size		int			-tsize
output directory    string      -out
capture modifier    int			-cmod
*/

type ProgramArguments struct {
	enviromentFile   string
	popSize          int
	maxGenerations   int
	minSuccessRate   float64
	steps            int
	mutationStrength float64
	tournamentSize   int
	outputPath       string
	captureModifier  int
}

func ParseProgramArguments() *ProgramArguments {
	var arguments = ProgramArguments{}

	flag.StringVar(&arguments.enviromentFile, "env", "", "Path to file with enviroment data.")
	flag.IntVar(&arguments.popSize, "pop", 0, "Number of generated microbes in each generation.")
	flag.IntVar(&arguments.maxGenerations, "maxg", 0, "Maximum of generated generations.")
	flag.Float64Var(&arguments.minSuccessRate, "mins", 0.0, "Minimum success rate of population to terminate program. (0.1 = 10%, 1 = 100%, etc.)")
	flag.IntVar(&arguments.steps, "steps", 0, "Number of steps of each microbe in every generation simulation.")
	flag.Float64Var(&arguments.mutationStrength, "mutstr", 0.0, "Gauss mutation strenght. (0.1 = 10%, 1 = 100%, etc.)")
	flag.IntVar(&arguments.tournamentSize, "tsize", 0, "Number of microbes from whom parent is selected in tournament.")
	flag.StringVar(&arguments.outputPath, "out", "", "Path of output directory.")
	flag.IntVar(&arguments.captureModifier, "cmod", 1, "Generation capture modifier. Only every nth generation will be captured on video.")
	flag.Parse()

	if arguments.enviromentFile == "" {
		log.Fatal("No enviroment file was specified")
	}
	if arguments.popSize < 2 {
		log.Fatal("Pop size cannot be less than 2")
	}
	if arguments.maxGenerations < 1 {
		log.Fatal("Number of max generations cannot be less than 1")
	}
	if arguments.minSuccessRate < 0 || arguments.minSuccessRate > 1.0 {
		log.Fatal("Min success rate must be in range 0.0-1.0")
	}
	if arguments.steps < 1 {
		log.Fatal("Number of steps cannot be less than 1")
	}
	if arguments.mutationStrength < 0 || arguments.minSuccessRate > 1.0 {
		log.Fatal("Mutation strength  must be in range 0.0-1.0")
	}
	if arguments.tournamentSize < 2 || arguments.tournamentSize > arguments.popSize {
		log.Fatal("Tournament size must be in range 2-pop")
	}
	if arguments.tournamentSize > arguments.popSize {
		log.Fatal("Tournament size cannot be more than population size")
	}
	if arguments.outputPath == "" {
		log.Fatal("Output path not set.")
	}

	return &arguments
}
