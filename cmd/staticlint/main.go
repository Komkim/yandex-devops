package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"yandex-devops/internal/staticlint"
)

func main() {
	checks := staticlint.NewAnalyzer()

	multichecker.Main(
		checks...,
	)
}
