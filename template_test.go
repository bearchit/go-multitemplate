package multitemplate

import (
	"html/template"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsage(t *testing.T) {
	mt := New()

	layouts, _ := filepath.Glob("testdata" + "/layouts/*.gohtml")
	viewsPath := "testdata" + "/views"
	views, _ := filepath.Glob(viewsPath + "/**/*")
	for _, view := range views {
		for _, layout := range layouts {
			files := append([]string{layout}, view)
			templateName, _ := filepath.Rel(viewsPath, view)
			mt.Add(templateName, template.Must(template.ParseFiles(files...)))
		}
	}

	var (
		tmpl *template.Template
		ok   bool
	)

	tmpl, ok = mt.Get("a/index.gohtml")
	assert.True(t, ok)
	assert.NotNil(t, tmpl)

	tmpl, ok = mt.Get("b/index.gohtml")
	assert.True(t, ok)
	assert.NotNil(t, tmpl)
}
