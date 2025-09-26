package exporters

import (
	"os"
	"path"
	"log"
	"html/template"
	"github.com/dialangproject/common/db"
	"github.com/dialangproject/common/models"
)

func ExportVSPTIntroPages(baseDir string) {

	vsptintroDir := path.Join(baseDir, "vsptintro")
	if err := os.MkdirAll(vsptintroDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "vsptintro.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {

		data := map[string]any{
			"title": db.GetTranslation("Title_Placement", al),
			"text": template.HTML(db.GetTranslation("PlacementIntro_Text", al)),
			"warningText": template.HTML(db.GetTranslation("Dialogues_SkipPlacement", al)),
			"yes": db.GetTranslation("Caption_Yes", al),
			"no": db.GetTranslation("Caption_No", al),
			"nexttooltip": db.GetTranslation("Caption_StartPlacement", al),
			"backtooltip": db.GetTranslation("Caption_BacktoChooseTest", al),
			"skipforwardtooltip": db.GetTranslation("Caption_SkipPlacement", al),
		}

		if f, err := os.Create(path.Join(vsptintroDir, al + ".html")); err == nil {
			tpl.ExecuteTemplate(f, "vsptintro.gohtml", data)
		} else {
			log.Fatal(err)
		}

		if f, err := os.Create(path.Join(vsptintroDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", data)
		} else {
			log.Fatal(err)
		}
    }
}

func ExportVSPTPages(baseDir string) {

	vsptBaseDir := path.Join(baseDir, "vspt")
	if err := os.MkdirAll(vsptBaseDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "vspt.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	testLanguagesAndVSPT := map[string][]models.VSPTWord{}

	for _, tl := range db.GetTestLanguageCodes() {
        testLanguagesAndVSPT[tl.Locale] = db.GetVSPTWords(tl.Locale)
	}

	for _, al := range db.AdminLocales {
		vsptDir := path.Join(vsptBaseDir, al)
		if err := os.MkdirAll(vsptDir, 0777); err != nil {
			log.Fatal(err)
		}

		title := db.GetTranslation("Title_Placement", al)
		submit := db.GetTranslation("Caption_SubmitAnswers", al)
		yes := db.GetTranslation("Caption_Yes", al)
		no := db.GetTranslation("Caption_No", al)
		confirmSend := db.GetTranslation("Dialogues_Submit", al)
		skipforwardtooltip := db.GetTranslation("Caption_QuitPlacement", al)

		tipData := map[string]string{"nexttooltip": submit, "skipforwardtooltip": skipforwardtooltip}
		if f, err := os.Create(path.Join(vsptBaseDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData)
			f.Close()
		} else {
			log.Fatal(err)
		}

		warningText := db.GetTranslation("Dialogues_SkipPlacement", al)

		for tl, wordList := range testLanguagesAndVSPT {

			tabList := []map[string]any{}
			words := []map[string]string{}

			numWords := len(wordList)
			for i := 0; i < numWords; i++ {

				word1 := wordList[i]

				var validClass1, invalidClass1, validClass2, invalidClass2, validClass3, invalidClass3 string

				if word1.Valid  == 1 {
					validClass1 = "correct"
					invalidClass1 = "incorrect"
				} else {
					validClass1 = "incorrect"
					invalidClass1 = "correct"
				}
				words = append(words, map[string]string{"id": word1.WordId})

				i++

				word2 := wordList[i]

				if word2.Valid == 1 {
					validClass2 = "correct"
					invalidClass2 = "incorrect"
				} else {
					validClass2 = "incorrect"
					invalidClass2 = "correct"
				}
				words = append(words, map[string]string{"id": word2.WordId})

				i++

				word3 := wordList[i]

				if word3.Valid  == 1 {
					validClass3 = "correct"
					invalidClass3 = "incorrect"
				} else {
					validClass3 = "incorrect"
					invalidClass3 = "correct"
				}
				words = append(words, map[string]string{"id": word3.WordId})

				tabList = append(tabList, map[string]any{
					"word1": word1.Word,
					"id1": word1.WordId,
					"validClass1": validClass1,
					"invalidClass1": invalidClass1,
					"word2": word2.Word,
					"id2": word2.WordId,
					"validClass2": validClass2,
					"invalidClass2": invalidClass2,
					"word3": word3.Word,
					"id3": word3.WordId,
					"validClass3": validClass3,
					"invalidClass3": invalidClass3,
				})

				data := map[string]any{
					"al": al,
					"title": title,
					"tl": tl,
					"warningText": warningText,
					"yes": yes,
					"no": no,
					"confirmsendquestion": confirmSend,
					"submit": submit,
					"words": words,
					"tab": tabList,
				}

				if f, err := os.Create(path.Join(vsptDir, tl + ".html")); err == nil {
					if err := tpl.ExecuteTemplate(f, "vspt.gohtml", data); err != nil {
						log.Fatal(err)
					}
					f.Close()
				} else {
					log.Fatal(err)
				}
			}
		}
	}
}

func ExportVSPTFeedbackPages(baseDir string) {

	feedbackDir := path.Join(baseDir, "vsptfeedback")
	if err := os.MkdirAll(feedbackDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(path.Join("templates", "vsptfeedback.gohtml"), path.Join("templates", "toolbartooltips.gojson"))
	if err != nil {
		log.Fatal(err)
	}

	//levels := db.GetVSPLevels()

	for _, al := range db.AdminLocales {
		title := db.GetTranslation("Title_PlacementFeedback", al)
		yourscore := db.GetTranslation("PlacementFeedback_YourScore", al)
		backtooltip := db.GetTranslation("Caption_BacktoFeedback", al)
		nexttooltip := db.GetTranslation("Caption_GotoNext", al)

		tipData := map[string]string{"backtooltip": backtooltip, "nexttooltip": nexttooltip}
		if f, err := os.Create(path.Join(feedbackDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData)
			f.Close()
		} else {
			log.Fatal(err)
		}

		levelv6 := db.GetTranslation("PlacementFeedback_Text#V6", al)
		levelv5 := db.GetTranslation("PlacementFeedback_Text#V5", al)
		levelv4 := db.GetTranslation("PlacementFeedback_Text#V4", al)
		levelv3 := db.GetTranslation("PlacementFeedback_Text#V3", al)
		levelv2 := db.GetTranslation("PlacementFeedback_Text#V2", al)
		levelv1 := db.GetTranslation("PlacementFeedback_Text#V1", al)

		data := map[string]string{
			"al": al,
			"Title": title,
			"Yourscore": yourscore,
			"LevelV6": levelv6,
			"LevelV5": levelv5,
			"LevelV4": levelv4,
			"LevelV3": levelv3,
			"LevelV2": levelv2,
			"LevelV1": levelv1,
			"nexttooltip": nexttooltip,
		}

		if f, err := os.Create(path.Join(feedbackDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "vsptfeedback.gohtml", data); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
    }
}

