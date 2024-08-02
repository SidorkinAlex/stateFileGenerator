package main

import (
	"fmt"
	"log"
	"stateFileGenerator/internal/CliApgParser"
	"stateFileGenerator/internal/CourceAnalyser"
	"stateFileGenerator/internal/copyFiler"
)

func main() {

	Args := CliApgParser.GetArgs()
	log.Println(Args)
	if Args.Action == "init" {
		CourceAnalyser.Anaslyse(Args.Sources[0])
		copyFiler.Copy(Args.Sources[0]+"/.result.csv", Args.Sources[0], Args.TargetDir)
	}
	if Args.Action == "status" {
		CourceAnalyser.CheckHashes(Args)
	}
	fmt.Println("\n")
	fmt.Println("\n")
}
