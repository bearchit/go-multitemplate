package multitemplate

import "html/template"

type Templates map[string]*template.Template

type Engine struct {
	templates Templates
}

func New() *Engine {
	return &Engine{
		templates: make(Templates),
	}
}

func (e *Engine) Add(name string, t *template.Template) {
	e.templates[name] = t
}

func (e *Engine) Get(name string) (*template.Template, bool) {
	v, ok := e.templates[name]
	return v, ok
}
