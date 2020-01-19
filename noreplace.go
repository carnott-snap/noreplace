package noreplace

import (
	"errors"

	"github.com/gostaticanalysis/modfile"
	xmodfile "golang.org/x/mod/modfile"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "noreplace",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		modfile.Analyzer,
	},
}

const Doc = "noreplace finds replace directive in go.mod"

func run(pass *analysis.Pass) (interface{}, error) {
	mf := pass.ResultOf[modfile.Analyzer].(*xmodfile.File)
	if mf != nil && len(mf.Replace) != 0 {
		return nil, errors.New("go.mod file has replace directives")
	}
	return nil, nil
}
