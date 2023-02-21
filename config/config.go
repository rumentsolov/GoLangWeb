package config

import (
	"log"
	"text/template"
)

// * holds the application configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
