package testcheck

import (
	"github.com/theoden9014/go-testcheck/passes/notestpkg"
	"golang.org/x/tools/go/analysis"
)

// Analyzers is all analysis.Analyzer in testcheck
var Analyzers = []*analysis.Analyzer{
	notestpkg.Analyzer,
}
