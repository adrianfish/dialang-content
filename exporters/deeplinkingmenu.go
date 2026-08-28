package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/adrianfish/dialang/content/db"
)

func ExportDeepLinkingMenuPages(baseDir string, baseAssetsUrl string) {

	deepLinkingMenuDir := path.Join(baseDir, "deep-linking-menu")

	if err := os.MkdirAll(deepLinkingMenuDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(path.Join("templates", "deep-linking-menu.gohtml"))
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {

		data := map[string]any{
			"baseAssetsUrl": baseAssetsUrl,
			"al": al,
			"tlsTitle": db.GetTranslation("Title_ChooseTest", al),
			"testrows": db.GetTestLanguagePrompts(al),
			"listeningTooltip": db.GetTranslation("ChooseTest_Skill#Listening", al),
			"writingTooltip": db.GetTranslation("ChooseTest_Skill#Writing", al),
			"readingTooltip": db.GetTranslation("ChooseTest_Skill#Reading", al),
			"structuresTooltip": db.GetTranslation("ChooseTest_Skill#Structures", al),
			"vocabularyTooltip": db.GetTranslation("ChooseTest_Skill#Vocabulary", al),
		}

		if f, err := os.Create(path.Join(deepLinkingMenuDir, al + ".html")); err == nil {
			tpl.ExecuteTemplate(f, "deep-linking-menu.gohtml", data)
		} else {
			log.Fatal(err)
		}
	}
}
