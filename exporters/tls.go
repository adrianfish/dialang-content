package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/dialangproject/common/db"
)

func ExportTLSPages(baseDir string) {

	tlsDir := path.Join(baseDir, "tls")

	if err := os.MkdirAll(tlsDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(path.Join("templates", "tls.gohtml"), path.Join("templates", "toolbartooltips.gojson"))
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {

		data := map[string]any{
			"tlsTitle": db.GetTranslation("Title_ChooseTest", al),
			"testrows": db.GetTestLanguagePrompts(al),
			"skipbacktooltip": db.GetTranslation("Caption_BacktoALS", al),
			"backtooltip": db.GetTranslation("Caption_BacktoWelcome", al),
			"listeningTooltip": db.GetTranslation("ChooseTest_Skill#Listening", al),
			"writingTooltip": db.GetTranslation("ChooseTest_Skill#Writing", al),
			"disclaimerTitle": db.GetTranslation("Title_UseMisuse", al), 
			"disclaimer": template.HTML(db.GetTranslation("Disclaimer_UseMisuse", al)),
			"testChosen": db.GetTranslation("Dialogues_TestSelected", al),
			"yes": db.GetTranslation("Caption_Yes", al),
			"no": db.GetTranslation("Caption_No", al),
			"ok": db.GetTranslation("Caption_OK", al),
			"done": db.GetTranslation("Caption_Done", al),
			"available": db.GetTranslation("Caption_Available", al),
			"notavailable": db.GetTranslation("Caption_NotAvailable", al),
			"readingTooltip": db.GetTranslation("ChooseTest_Skill#Reading", al),
			"structuresTooltip": db.GetTranslation("ChooseTest_Skill#Structures", al),
			"vocabularyTooltip": db.GetTranslation("ChooseTest_Skill#Vocabulary", al),
		}

		if f, err := os.Create(path.Join(tlsDir, al + ".html")); err == nil {
			tpl.ExecuteTemplate(f, "tls.gohtml", data)
		} else {
			log.Fatal(err)
		}

		if f, err := os.Create(path.Join(tlsDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", data)
		} else {
			log.Fatal(err)
		}
	}
}
