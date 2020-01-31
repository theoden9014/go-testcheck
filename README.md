# testcheck

## Install
```
go get -u github.com/theoden9014/go-testcheck/cmd/testcheck
```

## Analyzer

### notestpkg
```
import "github.com/theoden9014/go-testcheck/passes/notestpkg"
```

Check exists the test in the package. defaults exclude packages only "main".
Additional ignore packages by `-ignore` flag.

e.g.
```
testcheck -ignore github.com/theoden9014/go-testcheck/foo \
          -ignore github.com/theoden9014/go-testcheck/bar \
          ./...
```
