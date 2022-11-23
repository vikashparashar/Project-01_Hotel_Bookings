package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig hold's application config

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLogger    *log.Logger
	InProduction  bool // true in production mode else false in development mode
	Session       *scs.SessionManager
}
