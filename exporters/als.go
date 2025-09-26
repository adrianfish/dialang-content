package exporters

import (
	"html/template"
	"log"
	"bytes"
	"os"
	"path"
	"github.com/dialangproject/common/db"
	"github.com/dialangproject/common/models"
)

func ExportALS(baseDir string) {

	tpl := template.Must(template.ParseFiles(path.Join("templates", "als.gohtml"), path.Join("templates", "shell.gohtml")))

	if f, err := os.Create(path.Join(baseDir, "als.html")); err == nil {
		funderMessage := "The original DIALANG Project was carried out with the support of the commission of the European Communities within the framework of the SOCRATES programme, LINGUA 2"
		data := struct {
			Languages []models.AdminLanguage
			Fundermessage string
		}{
			Languages: db.GetAdminLanguages(),
			Fundermessage: funderMessage,
		}
		var b bytes.Buffer
		if err := tpl.ExecuteTemplate(&b, "als.gohtml", data); err == nil {
			if err := tpl.ExecuteTemplate(f, "shell.gohtml", map[string]any{"state": "als", "content": template.HTML(b.String())}); err != nil {
				log.Fatalf("Error rendering als shell: %w", err)
			}
		}
	} else {
		log.Fatalf("Failed to create als.html: %w", err)
	}
}
