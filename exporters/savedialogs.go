package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/adrianfish/dialang/content/db"
)

func ExportSaveDialogs(baseDir string) {

	saveDir := path.Join(baseDir, "save")

	if err := os.MkdirAll(saveDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles("templates/savedialog.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		savedInstruction := db.GetTranslation("Instruction_Saved", al)

		data := map[string]any{
			"savedInstruction": savedInstruction,
		}

		if f, err := os.Create(path.Join(saveDir, al + ".html")); err == nil {
      		tpl.ExecuteTemplate(f, "savedialog.gohtml", data)
		} else {
			log.Fatal(err)
		}
    }
}
