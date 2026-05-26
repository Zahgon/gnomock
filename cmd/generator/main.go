package main

import (
	"log"
	"regexp"
	"strings"
	"text/template"
)

const (
	registryPlaceholder = `// new presets go here.
`
	startPresetPlaceholder = `### /start/preset
`
	requestBodyPlaceholder = `### preset-request
`
	readmePlaceholder = `<!-- new presets go here -->
`
	ciTestPlaceholder = `### preset tests go here
`
	circleciJobsPlaceholder = `### circleci jobs go here
`
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

var fMap = template.FuncMap{
	"lower": strings.ToLower,
	"title": func(s string) string {
		return strings.Title(strings.ToLower(s)) // nolint: staticcheck
	},
	"snake": func(s string) string {
		snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

		return strings.ToLower(snake)
	},
}

type presetParams struct {
	Name        string
	DefaultPort int
	Image       string
	Public      bool
}

func main() {
	if err := generate(); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}

func generate() error { _ = "STUB: not implemented"; return nil }

// presetPkg generates a minimal working version of a preset. It also creates a
// README.md file that needs to be manually edited when the preset is ready.
func presetPkg(params presetParams) error { _ = "STUB: not implemented"; return nil }

func presetFile(dir, file string, params presetParams) error { _ = "STUB: not implemented"; return nil }

// gnomockdPkg adds the preset tests to gnomockd package.
func gnomockdPkg(params presetParams) error { _ = "STUB: not implemented"; return nil }

// registry adds the new preset to gnomockd preset registry so that it becomes
// available over HTTP.
func registry(params presetParams) error { _ = "STUB: not implemented"; return nil }

// swagger generates new definitions in swagger.yaml file. These definitions
// should be extended with options supported by a new preset.
func swagger(params presetParams) error { _ = "STUB: not implemented"; return nil }

func readme(params presetParams) error { _ = "STUB: not implemented"; return nil }

func github(params presetParams) error { _ = "STUB: not implemented"; return nil }

func circleci(params presetParams) error { _ = "STUB: not implemented"; return nil }

// replacePlaceholder replaces `placeholder` in `targetFile` with the result of
// `tmplFile` template execution using `params` values.
//
// nolint:gosec
func replacePlaceholder(targetFile, tmplFile, placeholder string, params presetParams) error {
	_ = "STUB: not implemented"
	return nil
}
