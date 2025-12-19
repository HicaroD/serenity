package complexity

import (
	"go/ast"
	"go/token"

	"github.com/serenitysz/serenity/internal/rules"
)

func CheckMaxFuncLinesNode(
	n ast.Node,
	fset *token.FileSet,
	cfg *rules.LinterOptions,
) []rules.Issue {
	complexity := cfg.Linter.Rules.Complexity

	if complexity != nil && complexity.Use != nil && !*complexity.Use {
		return nil
	}

	fn, ok := n.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	salve := fn.Body.List

	return nil
}
