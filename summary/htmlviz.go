// Génération HTML
package summary

import (
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
	}).Parse(htmlTemplate))

	f, err := os.Create(filename)
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

func DiffHtml(diff DiffSummary, filename string) {
	tmpl := template.Must(template.New("diff").Funcs(template.FuncMap{
		"toJson": func(v interface{}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		},
		"len": func(x []string) int { return len(x) },
		"sub": func(a, b int) int { return a - b },
	}).Parse(diffTemplate))

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmpl.Execute(f, diff)
}
