package rules

import (
	"go/ast"
	"go/token"
)

func CheckMaxParamsNode(
	n ast.Node,
	fset *token.FileSet,
	out []Issue,
	cfg *LinterOptions,
) []Issue {
	bestPractices := cfg.Linter.Rules.BestPractices

	if bestPractices == nil {
		return nil
	}

	if bestPractices.Use != nil && !*bestPractices.Use {
		return nil
	}

	var limit int8 = 5
	if err := VerifyIssues(cfg, out); err != nil {
		return nil
	}

	if bestPractices.MaxParams != nil &&
		bestPractices.MaxParams.Quantity != nil {
		limit = *bestPractices.MaxParams.Quantity
	}

	fn, ok := n.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	params := fn.Type.Params
	if params == nil {
		return nil
	}

	count := 0

	for _, field := range params.List {
		count += len(field.Names)

		if len(field.Names) == 0 {
			count++
		}
	}

	if int8(count) <= limit {
		return nil
	}

	out = append(out, Issue{
		Pos:     fset.Position(fn.Pos()),
		Message: "functions exceed the maximum parameter limit",
		Fix: func() {
			// Unsafe
			params.List = params.List[:limit]
		},
	})

	return out
}
