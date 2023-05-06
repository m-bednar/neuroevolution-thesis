/**
 * @project neuroevolution
 * @file    ArgumentParser.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

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
steps				int			-steps
mutation strength	float		-mutstr
tournament size		int			-tsize
output directory    string      -out
capture modifier    int			-cmod
neural net. scheme 	string 		-nn
*/

/*
Structure for storing parsed program arguments.
*/
type ProgramArguments struct {
	enviromentFile      string
	populationSize      int
	maxGenerations      int
	stepsCount          int
	mutationStrength    float64
	tournamentSize      int
	outputPath          string
	captureModifier     int
	neuralNetworkScheme string
}

/*
Parses and validates program arguments and then returns them as suitable structure.
*/
func ParseProgramArguments() *ProgramArguments {
	arguments := ProgramArguments{}

	flag.StringVar(&arguments.enviromentFile, "env", "", "Path to file with enviroment data.")
	flag.IntVar(&arguments.populationSize, "pop", 0, "Number of generated microbes in each generation.")
	flag.IntVar(&arguments.maxGenerations, "maxg", 0, "Maximum of generated generations.")
	flag.IntVar(&arguments.stepsCount, "steps", 0, "Number of steps of each microbe in every generation simulation.")
	flag.Float64Var(&arguments.mutationStrength, "mutstr", 0.0, "Gauss mutation strenght. (0.1 = 10%, 1 = 100%, etc.)")
	flag.IntVar(&arguments.tournamentSize, "tsize", 0, "Number of microbes from whom parent is selected in tournament.")
	flag.StringVar(&arguments.outputPath, "out", "", "Path of output directory.")
	flag.IntVar(&arguments.captureModifier, "cmod", 0, "Generation capture modifier. Only every nth generation will be captured on video.")
	flag.StringVar(&arguments.neuralNetworkScheme, "nn", "", "Scheme of neural network hidden layers, pattern: [count]x[width].")
	flag.Parse()

	if arguments.enviromentFile == "" {
		log.Fatal("No enviroment file was specified")
	}
	if arguments.populationSize < 2 {
		log.Fatal("Pop size cannot be less than 2")
	}
	if arguments.maxGenerations < 1 {
		log.Fatal("Number of max generations cannot be less than 1")
	}
	if arguments.stepsCount < 1 {
		log.Fatal("Number of steps cannot be less than 1")
	}
	if arguments.mutationStrength < 0 || arguments.mutationStrength > 1.0 {
		log.Fatal("Mutation strength  must be in range 0.0-1.0")
	}
	if arguments.tournamentSize < 2 || arguments.tournamentSize > arguments.populationSize {
		log.Fatal("Tournament size must be in range 2-[pop]")
	}
	if arguments.tournamentSize > arguments.populationSize {
		log.Fatal("Tournament size cannot be more than population size")
	}
	if arguments.outputPath == "" {
		log.Fatal("Output path not set.")
	}
	if arguments.neuralNetworkScheme == "" {
		log.Fatal("Neural network scheme not set.")
	}
	if arguments.captureModifier < 1 {
		log.Fatal("Capture modifier cannot be less than 1")
	}

	return &arguments
}
