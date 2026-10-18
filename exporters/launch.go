package exporters

import (
	"html/template"
	"log"
	"bytes"
	"os"
	"path"
)

func ExportLaunch(baseDir string) {

	log.Print(baseDir)

	tpl := template.Must(template.ParseFiles(path.Join("templates", "launch.gohtml"), path.Join("templates", "shell.gohtml")))

	if f, err := os.Create(path.Join(baseDir, "launch.html")); err == nil {
		launchData := map[string]string {
		}
		var b bytes.Buffer
		if err := tpl.ExecuteTemplate(&b, "launch.gohtml", launchData); err == nil {
			data := map[string]any {
				"state": "launch",
				"content": template.HTML(b.String()),
			}
			if err := tpl.ExecuteTemplate(f, "shell.gohtml", data); err != nil {
				log.Fatalf("Error rendering launch shell: %w", err)
			}
		}
	} else {
		log.Fatalf("Failed to create launch.html: %w", err)
	}
}
