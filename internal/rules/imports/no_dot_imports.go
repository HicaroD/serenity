package imports

import (
	"go/ast"
	"go/token"

	"github.com/serenitysz/serenity/internal/rules"
)

func CheckNoDotImports(
	f *ast.File,
	fset *token.FileSet,
	out []rules.Issue,
	cfg *rules.LinterOptions,
) []rules.Issue {
	imports := cfg.Linter.Rules.Imports

	if cfg.Linter.Use != nil && !*cfg.Linter.Use {
		return out
	}

	if err := rules.VerifyIssues(cfg, out); err != nil {
		return out
	}

	if imports == nil ||
		(imports.Use != nil && !*imports.Use) ||
		imports.NoDotImports == nil {
		return out
	}

	for _, i := range f.Imports {
		if i.Name != nil && i.Name.Name == "." {
			out = append(out, rules.Issue{
				Pos:     fset.Position(i.Name.NamePos),
				Message: "Imports should not be named with '.' ",
				Fix: func() {
					i.Name = nil
				},
			})
		}
	}

	return out
}
