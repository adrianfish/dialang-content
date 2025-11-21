package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/adrianfish/dialang-content/db"
)

func ExportHelpDialogs(baseDir string) {

	helpDir := path.Join(baseDir, "help")

	if err := os.MkdirAll(helpDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles("templates/helpdialog.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		key := db.GetTranslation("Title_Key", al)
		next := db.GetTranslation("Caption_ContinueNext", al)
		back := db.GetTranslation("Caption_BackPrevious", al)
		skipf := db.GetTranslation("Caption_SkipNextSection", al)
		skipb := db.GetTranslation("Caption_SkipPreviousSection", al)
		yes := db.GetTranslation("Caption_Yes", al)
		no := db.GetTranslation("Caption_No", al)
		help := db.GetTranslation("Caption_Help", al)
		smiley := db.GetTranslation("CaptionInstantOnOff",al)
		keyboard := db.GetTranslation("Caption_AdditionalCharacters",al)
		speaker := db.GetTranslation("Caption_PlaySound",al)
		aboutTitle := db.GetTranslation("Title_AboutDIALANG",al)
		crDialang := db.GetTranslation("Bits_CopyrightDIALANG",al)
		crLancaster := db.GetTranslation("Bits_CopyrightLancaster",al)
		vsptTitle := db.GetTranslation("Title_Placement",al)
		vsptText := db.GetTranslation("Help_Texts_Placement",al)
		saTitle := db.GetTranslation("Title_SelfAssess",al)
		saText := template.HTML(db.GetTranslation("Help_Texts_SelfAssess",al))
		data := map[string]any{
			"key": key,
			"next": next,
			"back": back,
			"skipf": skipf,
			"skipb": skipb,
			"yes": yes,
			"no": no,
			"help": help,
			"smiley": smiley,
			"keyboard": keyboard,
			"speaker": speaker,
			"aboutTitle": aboutTitle,
			"crDialang": crDialang,
			"crLancaster": crLancaster,
			"vsptTitle": vsptTitle,
			"vsptText": vsptText,
			"saTitle": saTitle,
			"saText": saText,
		}

		if f, err := os.Create(path.Join(helpDir, al + ".html")); err == nil {
      		tpl.ExecuteTemplate(f, "helpdialog.gohtml", data)
		} else {
			log.Fatal(err)
		}
    }
}
