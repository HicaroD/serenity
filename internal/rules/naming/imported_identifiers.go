package naming

import (
	"go/ast"
	"regexp"

	"github.com/serenitysz/serenity/internal/rules"
)

type ImportedIdentifiersRule struct{}

func (r *ImportedIdentifiersRule) Name() string {
	return "imported-identifiers"
}

func (r *ImportedIdentifiersRule) Targets() []ast.Node {
	return []ast.Node{(*ast.ImportSpec)(nil)}
}

func (r *ImportedIdentifiersRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	naming := runner.Cfg.Linter.Rules.Naming

	if naming == nil || (naming.Use != nil && !*naming.Use) || naming.ImportedIdentifiers == nil {
		return
	}

	maxIssues := rules.GetMaxIssues(runner.Cfg)

	if maxIssues > 0 && *runner.IssuesCount >= maxIssues {
		return
	}

	re, _ := regexp.Compile(*naming.ImportedIdentifiers.Pattern)

	spec := node.(*ast.ImportSpec)
	name := spec.Name

	if name != nil && !re.MatchString(name.Name)  {
		*runner.IssuesCount++

		*runner.Issues = append(*runner.Issues, rules.Issue{
			ArgStr1:  name.Name,
			ID:       rules.ImportedIdentifiersID,
			Pos:      runner.Fset.Position(spec.Path.ValuePos),
			Severity: rules.ParseSeverity(naming.ImportedIdentifiers.Severity),
		})
	}
}
