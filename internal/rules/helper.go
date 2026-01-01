package rules

func CanAutoFix(cfg *LinterOptions) bool {
	return cfg.Assistance != nil &&
		cfg.Assistance.Use != nil && *cfg.Assistance.Use &&
		cfg.Assistance.AutoFix != nil && *cfg.Assistance.AutoFix
}

// TODO: Change to uint16 (unsigned)
func GetMaxIssues(cfg *LinterOptions) int16 {
	if cfg.Linter.Issues != nil && cfg.Linter.Issues.Max != nil {
		return *cfg.Linter.Issues.Max
	}

	return 0
}
