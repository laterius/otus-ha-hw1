package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   = "views/layouts/"
	TemplateExt = ".html"
	TemplateDir = "views/"
)

// addTemplatePath takes slice of strings for file paths of templates and prepends
// the templateDir dir to each string in the slice:
// EG: the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func _(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes slice of strings to append the TemplateExt
// extension to each string in the slice
// EG: The input {"home"} would result the output {"home.html"} if the TemplateExt == ".html"
func _(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

// View Struct
type View struct {
	Template *template.Template
	Layout   string
}

func _() []string {
	files, err := filepath.Glob(layoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
