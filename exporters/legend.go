package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/dialangproject/common/db"
)

func ExportLegendPages(baseDir string) {

	legendDir := path.Join(baseDir, "legend")

	if err := os.MkdirAll(legendDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(path.Join("templates", "legend.gohtml"), path.Join("templates", "toolbartooltips.gojson"))
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		key := db.GetTranslation("Title_Key",al)
		welcome := db.GetTranslation("Title_WelcomeDIALANG",al)
		next := db.GetTranslation("Caption_ContinueNext",al)
		back := db.GetTranslation("Caption_BackPrevious",al)
		skipf := db.GetTranslation("Caption_SkipNextSection",al)
		skipb := db.GetTranslation("Caption_SkipPreviousSection",al)
		yes := db.GetTranslation("Caption_Yes",al)
		no := db.GetTranslation("Caption_No",al)
		help := db.GetTranslation("Caption_Help",al)
		smiley := db.GetTranslation("CaptionInstantOnOff",al)
		keyboard := db.GetTranslation("Caption_AdditionalCharacters",al)
		speaker := db.GetTranslation("Caption_PlaySound",al)
		backtooltip := db.GetTranslation("Caption_BacktoALS",al)
		nexttooltip := db.GetTranslation("Caption_ContinueNext",al)
		data := map[string]string{
			"key": key,
			"next": next,
			"back": back,
			"welcome": welcome,
			"skipf": skipf,
			"skipb": skipb,
			"yes": yes,
			"no": no,
			"help": help,
			"smiley": smiley,
			"keyboard": keyboard,
			"speaker": speaker,
			"backtooltip": backtooltip,
			"nexttooltip": nexttooltip,
			"al": al,
		}

		if f, err := os.Create(path.Join(legendDir, al + ".html")); err == nil {
			tpl.ExecuteTemplate(f, "legend.gohtml", data)
		} else {
			log.Fatal(err)
		}

		if f, err := os.Create(path.Join(legendDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", data)
		} else {
			log.Fatal(err)
		}
    }
}
