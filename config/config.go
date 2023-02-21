package config

import "text/template"

//* holds the application configurations
type AppConfig struct {
	
	TemplateCache map[string]*template.Template
}