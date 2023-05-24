package staticlint

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var ErrCheckAnalyzer = &analysis.Analyzer{
	Name: "mainexit",
	Doc:  "check for in main no os.exit",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if file.Name.Name == "main" {
				//for _, f := range file.Decls{
				//if f.
				//}
			}
			return true
		})

	}
	return nil, nil
}

//var errorType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
//
//func run(pass *analysis.Pass) (interface{}, error) {
//	expr := func(x *ast.ExprStmt) {
//		// проверяем, что выражение представляет собой вызов функции,
//		// у которой возвращаемая ошибка никак не обрабатывается
//		if call, ok := x.X.(*ast.CallExpr); ok {
//			if isReturnError(pass, call) {
//				pass.Reportf(x.Pos(), "expression returns unchecked error")
//			}
//		}
//	}
//	tuplefunc := func(x *ast.AssignStmt) {
//		// рассматриваем присваивание, при котором
//		// вместо получения ошибок используется '_'
//		// a, b, _ := tuplefunc()
//		// проверяем, что это вызов функции
//		if call, ok := x.Rhs[0].(*ast.CallExpr); ok {
//			results := resultErrors(pass, call)
//			for i := 0; i < len(x.Lhs); i++ {
//				// перебираем все идентификаторы слева от присваивания
//				if id, ok := x.Lhs[i].(*ast.Ident); ok && id.Name == "_" && results[i] {
//					pass.Reportf(id.NamePos, "assignment with unchecked error")
//				}
//			}
//		}
//	}
//	errfunc := func(x *ast.AssignStmt) {
//		// множественное присваивание: a, _ := b, myfunc()
//		// ищем ситуацию, когда функция справа возвращает ошибку,
//		// а соответствующий идентификатор слева равен '_'
//		for i := 0; i < len(x.Lhs); i++ {
//			if id, ok := x.Lhs[i].(*ast.Ident); ok {
//				// вызов функции справа
//				if call, ok := x.Rhs[i].(*ast.CallExpr); ok {
//					if id.Name == "_" && isReturnError(pass, call) {
//						pass.Reportf(id.NamePos, "assignment with unchecked error")
//					}
//				}
//			}
//		}
//	}
//	for _, file := range pass.Files {
//		// функцией ast.Inspect проходим по всем узлам AST
//		ast.Inspect(file, func(node ast.Node) bool {
//			switch x := node.(type) {
//			case *ast.ExprStmt: // выражение
//				expr(x)
//			case *ast.AssignStmt: // оператор присваивания
//				// справа одно выражение x,y := myfunc()
//				if len(x.Rhs) == 1 {
//					tuplefunc(x)
//				} else {
//					// справа несколько выражений x,y := z,myfunc()
//					errfunc(x)
//				}
//			}
//			return true
//		})
//	}
//	return nil, nil
//}
//
//func isErrorType(t types.Type) bool {
//	return types.Implements(t, errorType)
//}
//
//func resultErrors(pass *analysis.Pass, call *ast.CallExpr) []bool {
//	switch t := pass.TypesInfo.Types[call].Type.(type) {
//	case *types.Named: // возвращается значение
//		return []bool{isErrorType(t)}
//	case *types.Pointer: // возвращается указатель
//		return []bool{isErrorType(t)}
//	case *types.Tuple: // возвращается несколько значений
//		s := make([]bool, t.Len())
//		for i := 0; i < t.Len(); i++ {
//			switch mt := t.At(i).Type().(type) {
//			case *types.Named:
//				s[i] = isErrorType(mt)
//			case *types.Pointer:
//				s[i] = isErrorType(mt)
//			}
//		}
//		return s
//	}
//	return []bool{false}
//}
//
//// isReturnError возвращает true, если среди возвращаемых значений есть ошибка.
//func isReturnError(pass *analysis.Pass, call *ast.CallExpr) bool {
//	for _, isError := range resultErrors(pass, call) {
//		if isError {
//			return true
//		}
//	}
//	return false
//}
