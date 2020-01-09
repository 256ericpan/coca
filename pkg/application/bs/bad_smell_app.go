package bs

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/bs_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/bs"
	"path/filepath"
)

var nodeInfos []bs_domain.BsJClass

type BadSmellApp struct {
}

func NewBadSmellApp() *BadSmellApp {
	return &BadSmellApp{}
}

func (j *BadSmellApp) AnalysisPath(codeDir string) *[]bs_domain.BsJClass {
	nodeInfos = nil
	files := cocafile.GetJavaFiles(codeDir)
	for index := range files {
		nodeInfo := bs_domain.NewJFullClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := bs.NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.GetNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, nodeInfo)
	}

	return &nodeInfos
}

func (j *BadSmellApp) IdentifyBadSmell(nodeInfos *[]bs_domain.BsJClass, ignoreRules []string) []bs_domain.BadSmellModel {
	bsList := AnalysisBadSmell(*nodeInfos)

	mapIgnoreRules := make(map[string]bool)
	for _, ignore := range ignoreRules {
		mapIgnoreRules[ignore] = true
	}

	filteredBsList := bs_domain.FilterBadSmellList(bsList, mapIgnoreRules)
	return filteredBsList
}
