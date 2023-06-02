package staticlint

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var ErrCheckAnalyzer = &analysis.Analyzer{
	Name: "mainexit",
	Doc:  "check for in main no os.Exit",
	Run:  run,
}

// Запуск самописного линтера по проверке os.Exit в main
func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if v, ok := n.(*ast.FuncDecl); ok && v.Name.Name == "main" {
			i:
				for _, p := range v.Body.List {
					if es, ok := p.(*ast.ExprStmt); ok {
						if ce, ok := es.X.(*ast.CallExpr); ok {
							if se, ok := ce.Fun.(*ast.SelectorExpr); ok {
								if x, ok := se.X.(*ast.Ident); !ok && x.Name != "os" {
									continue i
								}
								if se.Sel.Name != "Exit" {
									continue i
								}
								pass.Reportf(v.Pos(), "there is os.Exit in the main function")
							}
						}
					}
				}
			}
			return true
		})

	}
	return nil, nil
}
