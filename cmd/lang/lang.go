package lang

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/eujandergois/mc-control/pkg/config"
	"golang.org/x/text/language"
)

type Lang string

const (
	PT_BR Lang = "ptBR"
	EN_US Lang = "enUS"
)

type Window struct {
	Name string `yaml:"name"`
	Id   string `yaml:"id"`
}

type AppWindows struct {
	Home     Window `yaml:"home"`
	Settings Window `yaml:"settings"`
}

type LangSchema struct {
	AppName     string     `yaml:"app-name"`
	AppVersion  string     `yaml:"app-version"`
	AppProducer string     `yaml:"app-producer"`
	AppWindows  AppWindows `yaml:"app-windows"`
}

func GetLang(lang Lang) (*LangSchema, error) {
	var langYAML LangSchema

	if lang == PT_BR {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)
		absPath := filepath.Join(dir, "ptBR.yaml")

		err := config.LoadYAML(absPath, &langYAML)
		if err != nil {
			return nil, err
		}
	}

	if lang == EN_US {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)
		absPath := filepath.Join(dir, "enUS.yaml")

		err := config.LoadYAML(absPath, &langYAML)
		if err != nil {
			return nil, err
		}
	}

	return &langYAML, nil
}

func GetSystemLanguage() language.Tag {
	locale := os.Getenv("LANG")

	if locale == "" {
		locale = os.Getenv("LC_ALL")
	}
	if locale == "" {
		locale = os.Getenv("LC_MESSAGES")
	}

	// Removing any encoding suffix like .UTF-8
	locale = strings.Split(locale, ".")[0]

	tag, _ := language.Parse(locale)
	return tag
}
