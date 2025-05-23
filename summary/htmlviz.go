// Génération HTML
package summary

import (
	"embed"
	"encoding/json"
	"html/template"
	"os"
)

func SummaryHtml(summary SummaryFile, filename string) {

	tmpl := template.Must(template.New("rapport").Funcs(template.FuncMap{
		"toJson": func(v interface{}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		},
	}).ParseFS(templateFS, "template.html"))

	f, err := os.Create("rapport.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmpl.Execute(f, summary)

}

func ToJSON(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

//go:embed template.html
var templateFS embed.FS
