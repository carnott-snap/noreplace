package noreplace_test

import (
	"fmt"
	"testing"

	"github.com/gostaticanalysis/noreplace"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	cases := []struct {
		pkg        string
		hasReplace bool
	}{
		{"a", false},
		{"b", true},
	}

	testdata := analysistest.TestData()

	for _, tt := range cases {
		tt := tt
		t.Run(tt.pkg, func(t *testing.T) {
			var hasErr bool
			analysistest.Run(testutil.Filter(t, func(format string, args ...interface{}) bool {
				if !tt.hasReplace {
					t.Fatalf("unexpected error:%s", fmt.Sprintf(format, args...))
				}
				hasErr = true
				return false
			}), testdata, noreplace.Analyzer, tt.pkg)

			if tt.hasReplace && !hasErr {
				t.Fatal("go.mod has replace directives but noreplace cannot find them")
			}
		})
	}

}
