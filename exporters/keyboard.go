package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"strings"
	"github.com/magiconair/properties"
)

func ExportKeyboardFragments(baseDir string) {

	keyboardDir := path.Join(baseDir, "keyboards")
	if err := os.MkdirAll(keyboardDir, 0755); err != nil {
		log.Fatalf("Failed to create keyboards directory: %w", err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "keyboard.gohtml")))

    // Load the special characters file
	specialCharLists := properties.MustLoadFile("special_chars.properties", properties.UTF8)

	for locale, csv := range specialCharLists.Map() {

		chars := strings.Split(csv, ",")

		if f, err := os.Create(path.Join(keyboardDir, locale + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "keyboard.gohtml", map[string][]string{"chars": chars}); err != nil {
				log.Fatalf("Failed to render keyboards file for locale %v: %w", locale, err)
			}
			f.Close()
		} else {
			log.Fatalf("Failed to create keyboards file for locale %v: %w", locale, err)
		}
    }
}
