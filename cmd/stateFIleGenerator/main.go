package main

import (
	"fmt"
	"github.com/SidorkinAlex/stateFileGenerator/internal/CliApgParser"
	"github.com/SidorkinAlex/stateFileGenerator/internal/CourceAnalyser"
	"log"
)

func main() {

	Args := CliApgParser.GetArgs()
	log.Println(Args)
	if Args.Action == "init" {
		CourceAnalyser.Anaslyse(Args.Sources[0])
	}
	fmt.Println("\n")
	fmt.Println("\n")
}
