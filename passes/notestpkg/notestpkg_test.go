package notestpkg_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/theoden9014/go-testcheck/passes/notestpkg"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	type flag struct {
		name  string
		value string
	}

	tests := []struct {
		dir   string // testdata directory (testdata/src/{{.dir}})
		flags []flag
		want  []string // Error messages
	}{
		{"a", nil, nil},
		{"b", nil, []string{`error analyzing notestpkg@b: "b" has not test files`}},
		{"b", []flag{flag{"ignore", "b"}}, nil},
		{"c", nil, nil},
		{"d", nil, nil},
	}

	for _, tt := range tests {
		var got []string
		t2 := errorFunc(func(s string) { got = append(got, s) })

		for _, f := range tt.flags {
			notestpkg.Analyzer.Flags.Set(f.name, f.value)
		}
		analysistest.Run(t2, testdata, notestpkg.Analyzer, tt.dir)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got:\n%s\nwant:\n%s",
				strings.Join(got, "\n"),
				strings.Join(tt.want, "\n"))
		}
	}
}

type errorFunc func(string)

func (f errorFunc) Errorf(format string, args ...interface{}) {
	f(fmt.Sprintf(format, args...))
}
