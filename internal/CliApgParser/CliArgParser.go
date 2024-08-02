package CliApgParser

import (
	"flag"
	"strings"
)

type CliParser struct {
	Action    string
	Sources   []string
	TargetDir string
}

//type Source struct {
//	Path string
//	Hash string
//}

func GetArgs() CliParser {
	var init bool
	var sources string
	var targetDir string
	var action string
	flag.StringVar(&sources, "s", "", "Sources parameter")
	flag.StringVar(&targetDir, "t", "", "target dir")
	flag.StringVar(&action, "a", "status", "action")
	flag.BoolVar(&init, "i", false, "set this param from initialize project")
	flag.Parse()
	if init {
		action = "init"
	}
	sourcesArr := strings.Split(sources, " ")
	CliParserCar := CliParser{
		Action:    action,
		Sources:   sourcesArr,
		TargetDir: targetDir,
	}
	return CliParserCar
}
