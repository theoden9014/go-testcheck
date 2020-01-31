package notestpkg

import (
	"fmt"
	"go/build"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = `
Check exists the test in the package.
defaults exclude packages only "main".
`

var Analyzer = &analysis.Analyzer{
	Name: "notestpkg",
	Doc:  Doc,
	Run:  run,
}

type StringsValue struct {
	strs map[string]struct{}
}

func (sv *StringsValue) String() string {
	return fmt.Sprint(sv.strs)
}

func (sv *StringsValue) Set(s string) error {
	if sv.strs == nil {
		sv.strs = make(map[string]struct{}, 5)
	}
	sv.strs[s] = struct{}{}

	return nil
}

func (sv *StringsValue) IsInclude(s string) bool {
	_, ok := sv.strs[s]
	return ok
}

var ignorePkg StringsValue

func init() {
	// set default ignore packages
	ignorePkg.Set("main")
	Analyzer.Flags.Var(&ignorePkg, "ignore", "ignore packages")
}

func run(pass *analysis.Pass) (interface{}, error) {
	var pkgname, pkgpath string

	// skip main and extest packages
	switch pkgname = pass.Pkg.Name(); {
	case ignorePkg.IsInclude(pkgname):
		return nil, nil
	case strings.HasSuffix(pkgname, "_test"):
		return nil, nil
	}
	pkgpath = pass.Pkg.Path()

	if len(pass.Files) == 0 {
		return nil, nil
	}

	pos := pass.Files[0].Pos()
	filename := pass.Fset.File(pos).Name()
	dir := filepath.Dir(filename)
	pkg, err := build.Default.ImportDir(dir, build.IgnoreVendor)
	if err != nil {
		return nil, err
	}

	// check exist the test files in the package
	if len(pkg.TestGoFiles) > 0 {
		return nil, nil
	}
	if len(pkg.XTestGoFiles) > 0 {
		return nil, nil
	}

	return nil, fmt.Errorf("%q has not test files", pkgpath)
}
