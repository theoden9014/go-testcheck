package main

import (
	"github.com/theoden9014/go-testcheck"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(testcheck.Analyzers...)
}
