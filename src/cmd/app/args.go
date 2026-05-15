package main

import (
	"errors"
	"flag"
)

func Args() (bool, error) {
	outputPath := outPath()
	inputPath := inPath()
	flag.Parse()
	if outputPath == "" {
		return false, errors.New("output path required")
	}
	if inputPath == "" {
		return false, errors.New("input path required")
	}
	return true, nil
}

func outPath() string {
	var outputPath string
	flag.StringVar(&outputPath, "o", "", "output file path")
	flag.StringVar(&outputPath, "out", "", "output file path")
	return outputPath
}

func inPath() string {
	var inputPath string
	flag.StringVar(&inputPath, "i", "", "input file path")
	flag.StringVar(&inputPath, "in", "", "input file path")
	return inputPath
}
