package exporters

import (
	"html/template"
	"log"
	"os"
	"io"
	"path"
	"strings"
	"github.com/dialangproject/common/db"
)

func ExportFeedbackMenuPages(baseDir string) {

	feedbackMenuDir := path.Join(baseDir, "feedbackmenu")
	if err := os.MkdirAll(feedbackMenuDir, 0777); err != nil {
		log.Fatalf("Failed to create feedback menu dir at %v: %w", feedbackMenuDir, err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "feedbackmenu.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {
    	title := template.HTML(db.GetTranslation("Title_FeedbackMenu",al))
		text := template.HTML(db.GetTranslation("FeedbackMenu_Text",al))
		resultsTitle := template.HTML(db.GetTranslation("Title_Results",al))
		yourLevelText := template.HTML(db.GetTranslation("FeedbackOption_Level",al))
		checkAnswersText := template.HTML(db.GetTranslation("FeedbackOption_CheckAnswers",al))
		placementTestText := template.HTML(db.GetTranslation("Title_Placement",al))
		saFeedbackText := template.HTML(db.GetTranslation("Title_SelfAssessFeedback",al))
		adviceTitle := template.HTML(db.GetTranslation("Title_Advice",al))
		adviceText := adviceTitle
		aboutSAText := template.HTML(db.GetTranslation("FeedbackOption_AboutSelfAssess", al))
		restartText := template.HTML(db.GetTranslation("Dialogues_QuitFeedback",al))
		yes := template.HTML(db.GetTranslation("Caption_Yes",al))
		no := template.HTML(db.GetTranslation("Caption_No",al))
		skipforwardtooltip := template.HTML(db.GetTranslation("Caption_ChooseAnotherTest",al))
		data := map[string]any{
			"al": al,
			"title": title,
			"text": text,
			"resultsTitle": resultsTitle,
			"yourLevelText": yourLevelText,
			"checkAnswersText": checkAnswersText,
			"placementTestText": placementTestText,
			"saFeedbackText": saFeedbackText,
			"adviceTitle": adviceTitle,
			"adviceText": adviceText,
			"aboutSAText": aboutSAText,
			"restartText": restartText,
			"yes": yes,
			"no": no,
		}

		if f, err := os.Create(path.Join(feedbackMenuDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "feedbackmenu.gohtml", data); err != nil {
				log.Fatalf("Failed to render feedbackmenu for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create feedbackmenu file for al %v: %w", al, err)
		}

		tipData := map[string]any{"skipbacktooltip": "", "skipforwardtooltip": skipforwardtooltip}
		if f, err := os.Create(path.Join(feedbackMenuDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create tooltip file for al %v: %w", al, err)
		}
    }
}

func ExportItemReviewPages(baseDir string) {

	itemreviewDir := path.Join(baseDir, "itemreview")
	if err := os.MkdirAll(itemreviewDir, 0777); err != nil {
		log.Fatalf("Failed to create itemreview dir at %v: %w", itemreviewDir, err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "itemreviewwrapper.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {

		title := template.HTML(db.GetTranslation("Title_ItemReview",al))
		text := template.HTML(db.GetTranslation("ItemReview_Text",al))
		backtooltip := template.HTML(db.GetTranslation("Caption_BacktoFeedback",al))
		data := map[string]any{"title": title, "text": template.HTML(text), "subskills": db.GetSubSkills(al)}

		if f, err := os.Create(path.Join(itemreviewDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "itemreviewwrapper.gohtml", data); err != nil {
				log.Fatalf("Failed to render itemreviewwrapper for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create itemreviewwrapper file for al %v: %w", al, err)
		}

	  	tipData := map[string]any{"backtooltip": backtooltip}
		if f, err := os.Create(path.Join(itemreviewDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create tooltip file for al %v: %w", al, err)
		}
    }
}

func ExportExplfbPages(baseDir string) {

	aboutSADir := path.Join(baseDir, "aboutsa")
	if err := os.MkdirAll(aboutSADir, 0777); err != nil {
		log.Fatalf("Failed to create aboutsa dir at %v: %w", aboutSADir, err)
	}

	tpl := template.Must(template.ParseFiles(
		path.Join("templates", "aboutsashell.gohtml"),
		path.Join("templates", "howoften.gohtml"),
		path.Join("templates", "onepart.gohtml"),
		path.Join("templates", "twoparts.gohtml"),
		path.Join("templates", "threeparts.gohtml"),
		path.Join("templates", "situations.gohtml"),
		path.Join("templates", "othertests.gohtml"),
		path.Join("templates", "differenttests.gohtml"),
		path.Join("templates", "reallife.gohtml"),
		path.Join("templates", "onepart.gohtml"),
		path.Join("templates", "otherreasons.gohtml"),
		path.Join("templates", "yourtargets.gohtml"),
		path.Join("templates", "toolbartooltips.gojson"),
	))

	for _, al := range db.AdminLocales {

		tipData := map[string]any{"backtooltip": template.HTML(db.GetTranslation("Caption_BacktoFeedback", al))}
		if f, err := os.Create(path.Join(aboutSADir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create tooltip file for al %v: %w", al, err)
		}

		alDir := path.Join(aboutSADir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatalf("Failed to create aboutsa dir for al %v: %w", al, err)
		}

		data := map[string]any{
			"title": template.HTML(db.GetTranslation("FeedbackOption_AboutSelfAssess", al)),
			"mainHeader": template.HTML(db.GetTranslation("Explanatory_Main_Header", al)),
			"subHeader": template.HTML(db.GetTranslation("Explanatory_Main_SubHeader", al)),
			"howOften": template.HTML(db.GetTranslation("Explanatory_Main_Menu_HowOften", al)),
			"how": template.HTML(db.GetTranslation("Explanatory_Main_Menu_Skills", al)),
			"situationsDiffer": template.HTML(db.GetTranslation("Explanatory_Main_Menu_Situations", al)),
			"otherLearners": template.HTML(db.GetTranslation("Explanatory_Main_Menu_OtherLearners", al)),
			"otherTests": template.HTML(db.GetTranslation("Explanatory_Main_Menu_OtherTests", al)),
			"yourTargets": template.HTML(db.GetTranslation("Explanatory_Main_Menu_Targets", al)),
			"realLife": template.HTML(db.GetTranslation("Explanatory_Main_Menu_RealLife", al)),
			"otherReasons": template.HTML(db.GetTranslation("Explanatory_Main_Menu_OtherReasons", al)),
		}

		if f, err := os.Create(path.Join(alDir, "index.html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "aboutsashell.gohtml", data); err != nil {
				log.Fatalf("Failed to render aboutsashell for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create aboutsashell file for al %v: %w", al, err)
		}

		if f, err := os.Create(path.Join(alDir, "main.html")); err == nil {
			main := db.GetTranslation("Explanatory_Main_Text", al)
			if _, err := io.WriteString(f, main); err != nil {
				log.Fatalf("Failed to write main.html for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create main.html file for al %v: %w", al, err)
		}

		doHowOften(al, alDir, tpl)
		doHow(al, alDir, tpl)
		doSituations(al, alDir, tpl)
		doOtherLearners(al, alDir, tpl)
		doOtherTests(al, alDir, tpl)
		doYourTargets(al, alDir, tpl)
		doRealLife(al, alDir, tpl)
		doOtherReasons(al, alDir, tpl)
    }
}

func doHowOften(al string, alDir string, tpl *template.Template) {

	infrequently := template.HTML(db.GetTranslation("ExpHowOften_Bullet1", al))
	longtime := template.HTML(db.GetTranslation("ExpHowOften_Bullet2", al))

	howOftenMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("Explanatory_Main_Menu_HowOften", al)),
		"howoften1": template.HTML(db.GetTranslation("ExpHowOften_Par1", al)),
		"bullet1": infrequently,
		"bullet2": longtime,
		"howoften2": template.HTML(db.GetTranslation("ExpHowOften_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "howoften.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "howoften.gohtml", howOftenMap); err != nil {
			log.Fatalf("Failed to render howoften for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create howoften file for al %v: %w", al, err)
	}

	infrequentlyMap := map[string]template.HTML{
		"title": infrequently,
		"part1": template.HTML(db.GetTranslation("ExpInfrequently_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpInfrequently_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "infrequently.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", infrequentlyMap); err != nil {
			log.Fatalf("Failed to render infrequently for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create infrequently file for al %v: %w", al, err)
	}

	longtimeMap := map[string]template.HTML{
		"title": longtime,
		"part1": template.HTML(db.GetTranslation("ExpLongTime_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpLongTime_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "longtime.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", longtimeMap); err != nil {
			log.Fatalf("Failed to render longtime for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create longtime file for al %v: %w", al, err)
	}
}

func doHow(al string, alDir string, tpl *template.Template) {

	howMap := map[string]any{
		"title": template.HTML(db.GetTranslation("ExpSomeSkills_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpSomeSkills_Par1", al)),
		"part2": template.HTML(template.HTML(db.GetTranslation("ExpSomeSkills_Par2", al))),
	}

	if f, err := os.Create(path.Join(alDir, "how.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", howMap); err != nil {
			log.Fatalf("Failed to render how for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create how file for al %v: %w", al, err)
	}

	overestimateMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpOverestimate_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpOverestimate_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpOverestimate_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "overestimate.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", overestimateMap); err != nil {
			log.Fatalf("Failed to render overestimate for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create overestimate file for al %v: %w", al, err)
	}

	underestimateMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpUnderestimate_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpUnderestimate_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "underestimate.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", underestimateMap); err != nil {
			log.Fatalf("Failed to render underestimate for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create underestimate file for al %v: %w", al, err)
	}
}

func doSituations(al string, alDir string, tpl *template.Template) {

	situtationsMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpSituation_Head", al)),
		"bullet1": template.HTML(db.GetTranslation("ExpSituation_Bullet1", al)),
		"bullet2": template.HTML(db.GetTranslation("ExpSituation_Bullet2", al)),
		"bullet3": template.HTML(db.GetTranslation("ExpSituation_Bullet3", al)),
	}

	if f, err := os.Create(path.Join(alDir, "situations.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "situations.gohtml", situtationsMap); err != nil {
			log.Fatalf("Failed to render situations for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create situations file for al %v: %w", al, err)
	}
}

func doOtherLearners(al string, alDir string, tpl *template.Template) {

	data := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpCompareOthers_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpCompareOthers_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpCompareOthers_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "otherlearners.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", data); err != nil {
			log.Fatalf("Failed to render otherlearners for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create otherlearners file for al %v: %w", al, err)
	}
}

func doOtherTests(al string, alDir string, tpl *template.Template) {

	data := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpOtherTests_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpOtherTests_Par1", al)),
		"bullet1": template.HTML(db.GetTranslation("ExpOtherTests_Bullet1", al)),
		"bullet2": template.HTML(db.GetTranslation("ExpOtherTests_Bullet2", al)),
		"bullet3": template.HTML(db.GetTranslation("ExpOtherTests_Bullet3", al)),
		"bullet4": template.HTML(db.GetTranslation("ExpOtherTests_Bullet4", al)),
		"part2": template.HTML(db.GetTranslation("ExpOtherTests_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "othertests.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "othertests.gohtml", data); err != nil {
			log.Fatalf("Failed to render othertests for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create othertests file for al %v: %w", al, err)
	}

	dtMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpDiffTests_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpDiffTests_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpDiffTests_Par2", al)),
		"part3": template.HTML(db.GetTranslation("ExpDiffTests_Par3", al)),
		"part4": template.HTML(db.GetTranslation("ExpDiffTests_Par4", al)),
		"part5": template.HTML(db.GetTranslation("ExpDiffTests_Par5", al)),
	}

	if f, err := os.Create(path.Join(alDir, "differenttests.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "differenttests.gohtml", dtMap); err != nil {
			log.Fatalf("Failed to render differenttests for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create differenttests file for al %v: %w", al, err)
	}

	sMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpSchoolTests_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpSchoolTests_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpSchoolTests_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "schooltests.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", sMap); err != nil {
			log.Fatalf("Failed to render schooltests for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create schooltests file for al %v: %w", al, err)
	}

	wMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpWorkTests_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpWorkTests_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "worktests.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", wMap); err != nil {
			log.Fatalf("Failed to render worktests for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create worktests file for al %v: %w", al, err)
	}

	iMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpIntTests_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpIntTests_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpIntTests_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "internationaltests.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", iMap); err != nil {
			log.Fatalf("Failed to render internationaltests for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create internationaltests file for al %v: %w", al, err)
	}
}

func doYourTargets(al string, alDir string, tpl *template.Template) {

	data := map[string]template.HTML{
		"title1": template.HTML(db.GetTranslation("ExpTargets_Head1", al)),
		"part1": template.HTML(db.GetTranslation("ExpTargets_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpTargets_Par2", al)),
		"title2": template.HTML(db.GetTranslation("ExpTargets_Head2", al)),
		"part3": template.HTML(db.GetTranslation("ExpTargets_Par3", al)),
		"part4": template.HTML(db.GetTranslation("ExpTargets_Par4", al)),
	}

	if f, err := os.Create(path.Join(alDir, "yourtargets.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "yourtargets.gohtml", data); err != nil {
			log.Fatalf("Failed to render yourtargets for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create yourtargets file for al %v: %w", al, err)
	}
}

func doRealLife(al string, alDir string, tpl *template.Template) {

	data := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpRealLife_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpRealLife_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpRealLife_Par2", al)),
		"bullet1": template.HTML(db.GetTranslation("ExpRealLife_Bullet1", al)),
		"bullet2": template.HTML(db.GetTranslation("ExpRealLife_Bullet2", al)),
		"bullet3": template.HTML(db.GetTranslation("ExpRealLife_Bullet3", al)),
		"bullet4": template.HTML(db.GetTranslation("ExpRealLife_Bullet4", al)),
		"bullet5": template.HTML(db.GetTranslation("ExpRealLife_Bullet5", al)),
		"bullet6": template.HTML(db.GetTranslation("ExpRealLife_Bullet6", al)),
	}

	if f, err := os.Create(path.Join(alDir, "reallife.html")); err == nil {
		if err	:= tpl.ExecuteTemplate(f, "reallife.gohtml", data); err != nil {
			log.Fatalf("Failed to render reallife for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create reallife file for al %v: %w", al, err)
	}		

	aMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpAnxiety_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpAnxiety_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpAnxiety_Par2", al)),
	}

	if f, err := os.Create(path.Join(alDir, "anxiety.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "twoparts.gohtml", aMap); err != nil {
			log.Fatalf("Failed to render anxiety for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create anxiety file for al %v: %w", al, err)
	}

	tMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpTimeAllowed_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpTimeAllowed_Par1", al)),
		"part2": template.HTML(db.GetTranslation("ExpTimeAllowed_Par2", al)),
		"part3": template.HTML(db.GetTranslation("ExpTimeAllowed_Par3", al)),
	}

	if f, err := os.Create(path.Join(alDir, "timeallowed.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "threeparts.gohtml", tMap); err != nil {
			log.Fatalf("Failed to render timeallowed for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create timeallowed file for al %v: %w", al, err)
	}

	sMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpSupport_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpSupport_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "support.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", sMap); err != nil {
			log.Fatalf("Failed to render support for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create support file for al %v: %w", al, err)
	}

	nMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpNumber_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpNumber_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "number.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", nMap); err != nil {
			log.Fatalf("Failed to render number for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create number file for al %v: %w", al, err)
	}

	fMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpFamiliarity_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpFamiliarity_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "familiarity.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", fMap); err != nil {
			log.Fatalf("Failed to render familiarity for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create familiarity file for al %v: %w", al, err)
	}

	mMap := map[string]template.HTML{
		"title": template.HTML(db.GetTranslation("ExpMedium_Head", al)),
		"part1": template.HTML(db.GetTranslation("ExpMedium_Par1", al)),
	}

	if f, err := os.Create(path.Join(alDir, "medium.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "onepart.gohtml", mMap); err != nil {
			log.Fatalf("Failed to render medium for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create medium file for al %v: %w", al, err)
	}		
}

func doOtherReasons(al string, alDir string, tpl *template.Template) {

	data := map[string]template.HTML{
		"title1": template.HTML(db.GetTranslation("ExpOtherReasons_Head1", al)),
		"part1": template.HTML(db.GetTranslation("ExpOtherReasons_Par1", al)),
		"bullet1": template.HTML(db.GetTranslation("ExpOtherReasons_Bullet1", al)),
		"bullet2": template.HTML(db.GetTranslation("ExpOtherReasons_Bullet2", al)),
		"part2": template.HTML(db.GetTranslation("ExpOtherReasons_Par2", al)),
		"title2": template.HTML(db.GetTranslation("ExpOtherReasons_Head2", al)),
		"part3": template.HTML(db.GetTranslation("ExpTargets_Par3", al)),
		"part4": template.HTML(db.GetTranslation("ExpTargets_Par4", al)),
	}

	if f, err := os.Create(path.Join(alDir, "otherreasons.html")); err == nil {
		if err := tpl.ExecuteTemplate(f, "otherreasons.gohtml", data); err != nil {
			log.Fatalf("Failed to render otherreasons for al %v: %w", al, err)
		}
	} else {
		log.Fatalf("Failed to create otherreasons file for al %v: %w", al, err)
	}
}

func ExportAdvfbPages(baseDir string) {

	tpl := template.Must(template.ParseFiles(
		path.Join("templates", "advfbshell.gohtml"),
		path.Join("templates", "advfba1.gohtml"),
		path.Join("templates", "advfba2.gohtml"),
		path.Join("templates", "advfbb1.gohtml"),
		path.Join("templates", "advfbb2.gohtml"),
		path.Join("templates", "advfbc1.gohtml"),
		path.Join("templates", "advfbc2.gohtml"),
		path.Join("templates", "advfb2itemadvice.gohtml"),
		path.Join("templates", "advfb3itemadvice.gohtml"),
		path.Join("templates", "advfb4itemadvice.gohtml"),
		path.Join("templates", "advfb5itemadvice.gohtml"),
		path.Join("templates", "advfb6itemadvice.gohtml"),
		path.Join("templates", "toolbartooltips.gojson"),
	))

	advfbDir := path.Join(baseDir, "advfb")
	if err := os.MkdirAll(advfbDir, 0777); err != nil {
		log.Fatalf("Failed to create advfb dir at %v: %w", advfbDir, err)
	}

	tls := db.GetTestLanguageCodes()

	for _, al := range db.AdminLocales {

		tipData := map[string]template.HTML{"backtooltip": template.HTML(db.GetTranslation("Caption_BacktoFeedback", al))}
		if f, err := os.Create(path.Join(advfbDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create tooltip file for al %v: %w", al, err)
		}

		alDir := path.Join(advfbDir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatalf("Failed to create advfb dir for al %v: %w", al, err)
		}

		shellData := map[string]template.HTML{
			"chooselevel": template.HTML(db.GetTranslation("Caption_ChooseLevel", al)),
			"howtoimprove": template.HTML(db.GetTranslation("Caption_HowtoImprove", al)),
		}
		if f, err := os.Create(path.Join(alDir, "index.html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "advfbshell.gohtml", shellData); err != nil {
				log.Fatalf("Failed to render advfbshell for al %v: %w", al, err)
			}
		} else {
			log.Fatalf("Failed to create advfbshell file for al %v: %w", al, err)
		}

		a1Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#A1", al))
		a2Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#A2", al))
		b1Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#B1", al))
		b2Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#B2", al))
		c1Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#C1", al))
		c2Header := template.HTML(db.GetTranslation("AdvisoryTable_Intro_#C2", al))
		
		for _, skill := range db.AdvfbSkills {

			skillDir := path.Join(alDir, strings.ToLower(skill))
			if err := os.MkdirAll(skillDir, 0777); err != nil {
				log.Fatalf("Failed to create advfb dir for al %v: %w", al, err)
			}

			a1Map := map[string]template.HTML{
				"header": a1Header,
				"row1heading": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al))),
				"row1texta1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#A1", al))),
				"row1texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#A2", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2texta1": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#A1", al)),
				"row2texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#A2", al)),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3texta1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#A1", al)),
				"row3texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#A2", al)),
			}

			if f, err := os.Create(path.Join(skillDir, "A1.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfba1.gohtml", a1Map); err != nil {
					log.Fatalf("Failed to render advfba1 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfba1 file for al %v: %w", al, err)
			}

			a2Map := map[string]any{
				"header": a2Header,
				"row1heading": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al)),
				"row1texta1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#A1", al))),
				"row1texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#A2", al)),
				"row1textb1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B1", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2texta1": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#A1", al)),
				"row2texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#A2", al)),
				"row2textb1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B1", al))),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3texta1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#A1", al)),
				"row3texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#A2", al)),
				"row3textb1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B1", al))),
			}

			if f, err := os.Create(path.Join(skillDir, "A2.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfba2.gohtml", a2Map); err != nil {
					log.Fatalf("Failed to render advfba2 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfba2 file for al %v: %w", al, err)
			}

			b1Map := map[string]any{
				"header": b1Header,
				"row1heading": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al)),
				"row1texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#A2", al)),
				"row1textb1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B1", al)),
				"row1textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B2", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#A2", al)),
				"row2textb1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B1", al))),
				"row2textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B2", al)),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3texta2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#A2", al)),
				"row3textb1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B1", al)),
				"row3textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B2", al)),
			}

			if f, err := os.Create(path.Join(skillDir, "B1.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfbb1.gohtml", b1Map); err != nil {
					log.Fatalf("Failed to render advfbb1 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfbb1 file for al %v: %w", al, err)
			}

			b2Map := map[string]template.HTML{
				"header": b2Header,
				"row1heading": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al)),
				"row1textb1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B1", al)),
				"row1textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B2", al)),
				"row1textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#C1", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2textb1": template.HTML(template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B1", al))),
				"row2textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B2", al)),
				"row2textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#C1", al)),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3textb1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B1", al)),
				"row3textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B2", al)),
				"row3textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#C1", al)),
			}

			if f, err := os.Create(path.Join(skillDir, "B2.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfbb2.gohtml", b2Map); err != nil {
					log.Fatalf("Failed to render advfbb2 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfbb2 file for al %v: %w", al, err)
			}

			c1Map := map[string]template.HTML{
				"header": c1Header,
				"row1heading": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al)),
				"row1textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#B2", al)),
				"row1textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#C1", al)),
				"row1textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#C2", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#B2", al)),
				"row2textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#C1", al)),
				"row2textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#C2", al)),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3textb2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#B2", al)),
				"row3textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#C1", al)),
				"row3textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#C2", al)),
			}

			if f, err := os.Create(path.Join(skillDir, "C1.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfbc1.gohtml", c1Map); err != nil {
					log.Fatalf("Failed to render advfbc1 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfbc1 file for al %v: %w", al, err)
			}

			c2Map := map[string]template.HTML{
				"header": c2Header,
				"row1heading": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Heading_#"+skill, al)),
				"row1textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#C1", al)),
				"row1textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row1_Text_#"+skill+"_#C2", al)),
				"row2heading": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Heading_#"+skill, al)),
				"row2textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#C1", al)),
				"row2textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row2_Text_#"+skill+"_#C2", al)),
				"row3heading": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Heading", al)),
				"row3textc1": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#C1", al)),
				"row3textc2": template.HTML(db.GetTranslation("AdvisoryTable_Row3_Text_#"+skill+"_#C2", al)),
			}

			if f, err := os.Create(path.Join(skillDir, "C2.html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "advfbc2.gohtml", c2Map); err != nil {
					log.Fatalf("Failed to render advfbc2 for al %v: %w", al, err)
				}
			} else {
				log.Fatalf("Failed to create advfbc2 file for al %v: %w", al, err)
			}

			for _, tl := range tls {

				tlDir := path.Join(skillDir, tl.Locale)
				if err := os.MkdirAll(tlDir, 0777); err != nil {
					log.Fatalf("Failed to create test language dir for tl %v: %w", tl.Locale, err)
				}

          		if skill == "Reading" {
					a1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item4", al)),
						"item5": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item5_#" + tl.TwoLetterCode, al)),
						"item6": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A1_#Item6_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "A1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb6itemadvice.gohtml", a1Map); err != nil {
							log.Fatalf("Failed to render advfb6 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb6 item advice file for al %v: %w", al, err)
					}

					a2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A2_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A2_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A2_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#A2_#Item4_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "A2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", a2Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					b1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B1_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B1_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B1_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B1_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", b1Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					b2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B2_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B2_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B2_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#B2_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", b2Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					c1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#C1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#C1_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#C1_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#C1_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#C1_#Item4_#" + tl.TwoLetterCode, al)),
						"item5": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Reading_#C1_#Item5_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "C1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb5itemadvice.gohtml", c1Map); err != nil {
							log.Fatalf("Failed to render advfb5 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb5 item advice file for al %v: %w", al, err)
					}
				} else if skill == "Writing" {
					a1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A1_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A1_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A1_#Item3_#" + tl.TwoLetterCode, al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A1_#Item4", al)),
						"item5": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A1_#Item5_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "A1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb5itemadvice.gohtml", a1Map); err != nil {
							log.Fatalf("Failed to render advfb5 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb5 item advice file for al %v: %w", al, err)
					}

					a2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A2_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A2_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A2_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A2_#Item4_#" + tl.TwoLetterCode, al)),
						"item5": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#A2_#Item5_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "A2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb5itemadvice.gohtml", a2Map); err != nil {
							log.Fatalf("Failed to render advfb5 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb5 item advice file for al %v: %w", al, err)
					}

					b1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B1_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B1_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B1_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B1_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", b1Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					b2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B2_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B2_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B2_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#B2_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", b2Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					c1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#C1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#C1_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#C1_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#C1_#Item3", al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#C1_#Item4", al)),
						"item5": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Writing_#C1_#Item5", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "C1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb5itemadvice.gohtml", c1Map); err != nil {	
							log.Fatalf("Failed to render advfb5 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb5 item advice file for al %v: %w", al, err)
					}
				} else if skill == "Listening" {
					dontUnderstand := "'" + db.GetTranslation("Utterances_DontUnderstand", tl.Locale) + "'"
					pleaseRepeat := "'" + db.GetTranslation("Utterances_PleaseRepeat",tl.Locale) + "'"
					sayAgain := "'" + db.GetTranslation("Utterances_SayAgain",tl.Locale) + "'"
					speakSlowly := "'" + db.GetTranslation("Utterances_SpeakSlowly",tl.Locale) + "'"
					item4 := strings.Replace(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A1_#Item4", al), "<utterance>Utterances_DontUnderstand</utterance>", dontUnderstand, 1)
					item4 = strings.Replace(item4, "<utterance>Utterances_PleaseRepeat</utterance>", pleaseRepeat, 1)
					item4 = strings.Replace(item4, "<utterance>Utterances_SayAgain</utterance>", sayAgain, 1)
					item4 = strings.Replace(item4, "<utterance>Utterances_SpeakSlowly</utterance>", speakSlowly, 1)
					a1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A1_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A1_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A1_#Item3", al)),
						"item4": template.HTML(item4),
					}
					if f, err := os.Create(path.Join(tlDir, "A1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", a1Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					a2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#A2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A2_#Item1", al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A2_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A2_#Item3_#" + tl.TwoLetterCode, al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#A2_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "A2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", a2Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					b1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B1_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B1_#Item2_#" + tl.TwoLetterCode, al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B1_#Item3_#" + tl.TwoLetterCode, al)),
						"item4": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B1_#Item4", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb4itemadvice.gohtml", b1Map); err != nil {
							log.Fatalf("Failed to render advfb4 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb4 item advice file for al %v: %w", al, err)
					}

					b2Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#B2", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B2_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B2_#Item2", al)),
						"item3": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#B2_#Item3", al)),
					}
					if f, err := os.Create(path.Join(tlDir, "B2.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb3itemadvice.gohtml", b2Map); err != nil {
							log.Fatalf("Failed to render advfb3 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb3 item advice file for al %v: %w", al, err)
					}

					c1Map := map[string]template.HTML{
						"header": template.HTML(db.GetTranslation("AdvisoryTips_Intro_#C1", al)),
						"item1": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#C1_#Item1_#" + tl.TwoLetterCode, al)),
						"item2": template.HTML(db.GetTranslation("AdvisoryTips_Bullet_#Listening_#C1_#Item2_#" + tl.TwoLetterCode, al)),
					}
					if f, err := os.Create(path.Join(tlDir, "C1.html")); err == nil {
						if err := tpl.ExecuteTemplate(f, "advfb2itemadvice.gohtml", c1Map); err != nil {
							log.Fatalf("Failed to render advfb2 item advice for al %v: %w", al, err)
						}
					} else {
						log.Fatalf("Failed to create advfb2 item advice file for al %v: %w", al, err)
					}
				} // Listening
			} // tl
		}
    }
}

func ExportTestResultPages(baseDir string) {

	testresultsDir := path.Join(baseDir, "testresults")
	if err := os.MkdirAll(testresultsDir, 0777); err != nil {
		log.Fatalf("Failed to create testresults dir at %v: %s", testresultsDir, err)
	}

	tpl := template.Must(template.ParseFiles(
		path.Join("templates", "testresults.gohtml"),
		path.Join("templates", "toolbartooltips.gojson"),
	))

	for _, al := range db.AdminLocales {

		tipData := map[string]string{"backtooltip": db.GetTranslation("Caption_BacktoWelcome", al)}
		if f, err := os.Create(path.Join(testresultsDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips for al %v: %s", al, err)
			}
		} else {
			log.Fatalf("Failed to create tooltip file for al %v: %s", al, err)
		}

		alDir := path.Join(testresultsDir, al)
		if err := os.MkdirAll(alDir, 0777); err != nil {
			log.Fatalf("Failed to create al dir at %v: %s", alDir, err)
		}

		title := db.GetTranslation("Title_DIALANGTestResults", al)

      	for _, skill := range db.TestSkills {

			skillDir := path.Join(alDir, strings.ToLower(skill))
			if err := os.MkdirAll(skillDir, 0777); err != nil {
				log.Fatalf("Failed to create skill dir at %v: %s", skillDir, err)
			}

			explanTexts := map[string]template.HTML{}
			for _, itemLevel := range db.ItemLevels {
				text := db.GetTranslation("TestResults_Text#" + skill + "#" + itemLevel, al)
				explanTexts[itemLevel + "Explanation"] = template.HTML(text)
			}

			for _, itemLevel := range db.ItemLevels {
				text := db.GetTranslation("TestResults_Text#" + skill + "#" + itemLevel, al)
				data := map[string]any{
					"al": al,
					"title": title,
					"text": text,
					"itemLevel": itemLevel,
				}
				for k, v := range explanTexts {
					data[k] = v
				}

				if f, err := os.Create(path.Join(skillDir, itemLevel + ".html")); err == nil {
					if err := tpl.ExecuteTemplate(f, "testresults.gohtml", data); err != nil {
						log.Fatalf("Failed to render testresults for al %v: %s", al, err)
					}
				} else {
					log.Fatalf("Failed to create testresults file for al %v: %s", al, err)
				}
        	}
      	}
    }
}

