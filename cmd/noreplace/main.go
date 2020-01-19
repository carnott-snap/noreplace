package main

import (
	"github.com/gostaticanalysis/noreplace"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noreplace.Analyzer) }
