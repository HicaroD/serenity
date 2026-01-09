package bestpractices

import (
	"go/ast"
	"strings"

	"github.com/serenitysz/serenity/internal/rules"
)

type GetMustReturnValueRule struct{}

func (r *GetMustReturnValueRule) Name() string {
	return "get-must-return-value"
}

func (r *GetMustReturnValueRule) Targets() []ast.Node {
	return []ast.Node{(*ast.FuncDecl)(nil)}
}

func (r *GetMustReturnValueRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	if max := runner.Cfg.GetMaxIssues(); max > 0 && *runner.IssuesCount >= max {
		return
	}

	cfg := runner.Cfg.Linter.Rules.BestPractices

	if cfg == nil || !cfg.Use || cfg.GetMustReturnValue == nil {
		return
	}

	fn := node.(*ast.FuncDecl)

	if fn.Name == nil || !strings.HasPrefix(fn.Name.Name, "Get") {
		return
	}

	results := fn.Type.Results

	if results == nil || len(results.List) == 0 {
		r.report(runner, fn)
		return
	}

	nonErrorReturns := 0

	for _, field := range results.List {
		if isErrorType(field.Type) {
			continue
		}

		count := max(1, len(field.Names))

		nonErrorReturns += count

		if nonErrorReturns > 0 {
			return
		}
	}

	r.report(runner, fn)
}

func (r *GetMustReturnValueRule) report(runner *rules.Runner, fn *ast.FuncDecl) {
	*runner.IssuesCount++

	*runner.Issues = append(*runner.Issues, rules.Issue{
		ID:  rules.GetMustReturnValueID,
		Pos: runner.Fset.Position(fn.Name.Pos()),
		Severity: rules.ParseSeverity(
			runner.Cfg.Linter.Rules.BestPractices.GetMustReturnValue.Severity,
		),
	})
}

func isErrorType(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name == "error"
	case *ast.SelectorExpr:
		return false
	}

	return false
}
