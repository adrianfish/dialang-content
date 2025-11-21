package exporters

import (
	"html/template"
	"log"
	"os"
	"path"
	"github.com/adrianfish/dialang-content/db"
)

func ExportFlowchartPages(baseDir string) {

	flowchartDir := path.Join(baseDir, "flowchart")

	if err := os.MkdirAll(flowchartDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl:= template.Must(template.ParseFiles(path.Join("templates", "flowchart.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {
		welcomeTitle := db.GetTranslation("Title_WelcomeDIALANG",al)
		welcomeText := db.GetTranslation("Welcome_Intro_Text",al)
		procedureTitle := db.GetTranslation("Title_ProcedureCAPS",al)
		procedureText := db.GetTranslation("Welcome_Procedure_Text",al)
		backtooltip := db.GetTranslation("Caption_BacktoWelcome",al)
		nexttooltip := db.GetTranslation("Caption_GotoChooseTest",al)
		stage1Title := db.GetTranslation("Title_ChooseTest",al)
		stage1 := db.GetTranslation("Welcome_Chart_ChooseTest_Text",al)
		stage2Title := db.GetTranslation("Title_Placement",al)
		stage2 := db.GetTranslation("Welcome_Chart_Placement_Text",al)
		stage3Title := db.GetTranslation("Title_SelfAssess",al)
		stage3 := db.GetTranslation("Welcome_Chart_SelfAssess_Text",al)
		stage4Title := db.GetTranslation("Title_LangTest",al)
		stage4 := db.GetTranslation("Welcome_Chart_LangTest_Text",al)
		stage5Title := db.GetTranslation("Title_FeedbackResultsAdvice",al)
		stage5 := db.GetTranslation("Welcome_Chart_Feedback_Text",al)
		data := map[string]string{
			"al": al,
			"welcomeTitle": welcomeTitle,
			"welcomeText": welcomeText,
			"procedureTitle": procedureTitle,
			"procedureText": procedureText,
			"stage1Title": stage1Title,
			"stage1": stage1,
			"stage2Title": stage2Title,
			"stage2": stage2,
			"stage3Title": stage3Title,
			"stage3": stage3,
			"stage4Title": stage4Title,
			"stage4": stage4,
			"stage5Title": stage5Title,
			"stage5": stage5,
			"backtooltip": backtooltip,
			"nexttooltip": nexttooltip,
		}

		if f, err := os.Create(path.Join(flowchartDir, al + ".html")); err == nil {
			tpl.ExecuteTemplate(f, "flowchart.gohtml", data)
		} else {
			log.Fatal(err)
		}

		if f, err := os.Create(path.Join(flowchartDir, al + "-toolbarTooltips.json")); err == nil {
			tpl.ExecuteTemplate(f, "toolbartooltips.gojson", data)
		} else {
			log.Fatal(err)
		}
	}
}
