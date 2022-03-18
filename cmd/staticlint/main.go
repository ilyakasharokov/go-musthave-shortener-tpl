// Статический анализатор кода.
// Запуск из корня проекта:
// ./cmd/staticlint/staticlint ./...
package main

import (
	"encoding/json"
	"ilyakasharokov/pkg/myanalyzer"
	"os"
	"path/filepath"

	gocr "github.com/go-critic/go-critic/checkers/analyzer"
	"github.com/gostaticanalysis/nilerr"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

const Config = `config.json`

// ConfigData описывает структуру файла конфигурации.
type ConfigData struct {
	Staticcheck []string
}

func main() {
	appfile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(appfile), Config))
	if err != nil {
		panic(err)
	}
	var cfg ConfigData
	if err = json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	var mychecks []*analysis.Analyzer
	mychecks = []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		bools.Analyzer,
		assign.Analyzer,
		lostcancel.Analyzer,
		copylock.Analyzer,
		myanalyzer.OsExitAnalyzer,
		gocr.Analyzer, // Вопрос ментору: данный анализатор находит ошибки в автогенерируемых моках в /internal/app/mocks, как заставить игнорировать фаилы?
		nilerr.Analyzer,
	}
	checks := make(map[string]bool)
	for _, v := range cfg.Staticcheck {
		checks[v] = true
	}
	// добавляем анализаторы из staticcheck, которые указаны в файле конфигурации
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	multichecker.Main(
		mychecks...,
	)

}
