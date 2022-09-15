package config

import "html/template"

// holds the app wide config
type AppCondig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}
