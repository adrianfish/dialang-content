package exporters

import (
	"os"
	"log"
	"path"
	"strings"
	"html/template"
	"github.com/adrianfish/dialang-content/db"
)

func ExportSAIntroPages(baseDir string) {

	saintroDir := path.Join(baseDir, "saintro")
	if err := os.MkdirAll(saintroDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "saintro.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {
		alDir := path.Join(saintroDir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatal(err)
		}

		nexttooltip := db.GetTranslation("Caption_StartSelfAssess",al)
		skipforwardtooltip := db.GetTranslation("Caption_SkipSelfAssess",al)
		text := template.HTML(db.GetTranslation("SelfAssessIntro_Text", al))

		// Confirmation dialog texts.
		warningText := db.GetTranslation("Dialogues_SkipSelfAssess",al)
		yes := db.GetTranslation("Caption_Yes",al)
		no := db.GetTranslation("Caption_No",al)

		tipData := map[string]string{"nexttooltip": nexttooltip,"skipforwardtooltip": skipforwardtooltip}
		if f, err := os.Create(path.Join(saintroDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData)
			f.Close()
		} else {
			log.Fatal(err)
		}

		for _, skill := range db.SaSkills {

			data := map[string]any {
				"skill": strings.ToLower(skill),
				"title": db.GetTranslation("Title_SelfAssess#" + skill, al),
				"text": text,
				"warningText": warningText,
				"yes": yes,
				"no": no,
			}

			if f, err := os.Create(path.Join(alDir, strings.ToLower(skill) + ".html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "saintro.gohtml", data); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		}
	}
}

func ExportSAPages(baseDir string) {

	saDir := path.Join(baseDir, "sa")

	tpl, err := template.ParseFiles(path.Join("templates", "sa.gohtml"), path.Join("templates", "toolbartooltips.gojson"))
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		alDir := path.Join(saDir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatal(err)
		}

		submit := db.GetTranslation("Caption_SubmitAnswers",al)
		skipforwardtooltip := db.GetTranslation("Caption_QuitSelfAssess",al)
		yes := db.GetTranslation("Caption_Yes",al)
		no := db.GetTranslation("Caption_No",al)
		confirmSend := db.GetTranslation("Dialogues_Submit",al)

		// Confirmation dialog texts.
		warningText := db.GetTranslation("Dialogues_SkipSelfAssess",al)

		tipData := map[string]string{"nexttooltip": submit,"skipforwardtooltip": skipforwardtooltip}
		if f, err := os.Create(path.Join(saDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData)
			f.Close()
		} else {
			log.Fatal(err)
		}

		for _, skill := range db.SaSkills {

			statements := db.GetSAStatements(al, strings.ToLower(skill))

			title := db.GetTranslation("Title_SelfAssess#" + skill, al)
			data := map[string]any{
				"al": al,
				"title": title,
				"warningText": warningText,
				"submit": submit,
				"yes": yes,
				"no": no,
				"confirmsendquestion": confirmSend,
				"statements": statements,
			}
			if f, err := os.Create(path.Join(alDir, strings.ToLower(skill) + ".html")); err == nil {
				tpl.ExecuteTemplate(f, "sa.gohtml", data)
				f.Close()
			} else {
				log.Fatal(err)
			}
      	}
    }
}

func ExportSAFeedbackPages(baseDir string) {

	saFeedbackDir := path.Join(baseDir, "safeedback")
	if err := os.MkdirAll(saFeedbackDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(path.Join("templates", "safeedback.gohtml"), path.Join("templates", "toolbartooltips.gojson"))
	if err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		alDir := path.Join(saFeedbackDir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatal(err)
		}

		backtooltip := db.GetTranslation("Caption_BacktoFeedback", al)
		if f, err := os.Create(path.Join(saFeedbackDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", map[string]string{"backtooltip": backtooltip})
			f.Close()
		} else {
			log.Fatal(err)
		}

		title := db.GetTranslation("Title_SelfAssessFeedback",al)
		aboutSAText := db.GetTranslation("FeedbackOption_AboutSelfAssess",al)
		overEst := db.GetTranslation("SelfAssessFeedback_OverEst_Par2",al)
		accurate := db.GetTranslation("SelfAssessFeedback_Match_Par2",al)
		underEst := db.GetTranslation("SelfAssessFeedback_UnderEst_Par2",al)

		for _, itemLevel := range db.ItemLevels {

			itemLevelDir := path.Join(alDir, itemLevel)
			if err := os.MkdirAll(itemLevelDir, 0777); err != nil {
				log.Fatal(err)
			}

			for _, saLevel := range db.ItemLevels {
				var partOne string
				if saLevel != itemLevel {
					partOne = db.GetTranslationLike("SelfAssessFeedback%Par1#" + itemLevel + "#" + saLevel, al)
              	} else {
                	partOne = db.GetTranslation("SelfAssessFeedback_Match_Par1#" + itemLevel, al)
				}

          		var partTwo string
              	if saLevel > itemLevel {
                	partTwo = overEst
              	} else if saLevel == itemLevel {
                	partTwo = accurate
              	} else {
                	partTwo = underEst
              	}

				data := map[string]string{
					"al": al,
					"title": title,
					"partOne": partOne,
					"partTwo": partTwo,
					"aboutSAText": aboutSAText,
				}

				if f, err := os.Create(path.Join(itemLevelDir, saLevel + ".html")); err == nil {
					tpl.ExecuteTemplate(f, "safeedback.gohtml", data)
					f.Close()
				} else {
					log.Fatal(err)
				}
			}
        }
	}
}


