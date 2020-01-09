package analysis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/full"
	"path/filepath"
)

type JavaFullApp struct {
}

func NewJavaFullApp() JavaFullApp {
	return *&JavaFullApp{}
}

func (j *JavaFullApp) AnalysisPath(codeDir string, classes []string, identNodes []domain.JIdentifier) []domain.JClassNode {
	files := cocafile.GetJavaFiles(codeDir)
	return j.AnalysisFiles(identNodes, files, classes)
}

func (j *JavaFullApp) AnalysisFiles(identNodes []domain.JIdentifier, files []string, classes []string) []domain.JClassNode {
	var nodeInfos []domain.JClassNode

	var identMap = make(map[string]domain.JIdentifier)
	for _, ident := range identNodes {
		identMap[ident.GetClassFullName()] = ident
	}

	for _, file := range files {
		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := full.NewJavaFullListener(identMap, file)
		listener.AppendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes := listener.GetNodeInfo()
		nodeInfos = append(nodeInfos, nodes...)
	}

	return nodeInfos
}
