package unused

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"path/filepath"
	"testing"
)

func TestRemoveUnusedImportApp_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/refactor/unused"
	codePath = filepath.FromSlash(codePath)
	cocatest.ResetGitDir(codePath)

	deps1, _, _ := cocatest.BuildAnalysisDeps(codePath)
	g.Expect(len(deps1[0].Imports)).To(Equal(3))

	app := NewRemoveUnusedImportApp(codePath)
	results := app.Analysis()

	g.Expect(len(results)).To(Equal(1))

	errorLines := BuildErrorLines(results[0])
	g.Expect(errorLines).To(Equal([]int{3,4,5}))

	app.Refactoring(results)

	deps, _, _ := cocatest.BuildAnalysisDeps(codePath)
	g.Expect(len(deps[0].Imports)).To(Equal(0))
	cocatest.ResetGitDir(codePath)
}
