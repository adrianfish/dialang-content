package exporters

import (
	"html/template"
	"bytes"
	"strings"
	"log"
	"regexp"
	"os"
	"path"
	"strconv"
	"github.com/adrianfish/dialang-content/db"
)

func ExportTestIntroPages(baseDir string) {

	testIntroDir := path.Join(baseDir, "testintro")
	if err := os.MkdirAll(testIntroDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "testintro.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {

		tipData := map[string]string{
			"nexttooltip": db.GetTranslation("Caption_StartTest", al), 
			"skipforwardtooltip": db.GetTranslation("Caption_SkipLangTest", al),
		}

		if f, err := os.Create(path.Join(testIntroDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}

		data := map[string]any{
			"al": al,
			"title": db.GetTranslation("Title_DIALANGLangTest", al),
			"text": template.HTML(db.GetTranslation("LangTestIntro_Text", al)),
			"warningText": template.HTML(db.GetTranslation("Dialogues_SkipLangTest", al)),
			"yes": db.GetTranslation("Caption_Yes", al),
			"no": db.GetTranslation("Caption_No",al),
			"feedback": db.GetTranslation("Caption_InstantFeedback", al),
			"instantfeedbackontooltip": db.GetTranslation("Caption_InstantFeedbackOn", al), 
			"instantfeedbackofftooltip": db.GetTranslation("Caption_InstantFeedbackOff", al),
		}

		if f, err := os.Create(path.Join(testIntroDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "testintro.gohtml", data); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
    }
}

func ExportBasketPages(baseDir string) {

	tpl := template.Must(template.ParseFiles(
		path.Join("templates", "basket.gohtml"),
		path.Join("templates", "audiomedia.gohtml"),
		path.Join("templates", "imagemedia.gohtml"),
		path.Join("templates", "textmedia.gohtml"),
		path.Join("templates", "mcqresponse.gohtml"),
		path.Join("templates", "saresponse.gohtml"),
		path.Join("templates", "gtresponse.gohtml"),
		path.Join("templates", "gdresponse.gohtml"),
		path.Join("templates", "tabbedpaneresponse.gohtml"),
		path.Join("templates", "toolbartooltips.gojson"),
	))

	typeMap := map[string]string{
		"gapdrop": "GapDrop",
		"gaptext": "GapText",
		"mcq": "MCQ",
		"shortanswer": "ShortAnswer",
		"tabbedpane": "TabbedMCQ",
	}

	basketDir := path.Join(baseDir, "baskets")
	if err := os.MkdirAll(basketDir, 0777); err != nil {
		log.Fatal(err)
	}

	for _, al := range db.AdminLocales {
		tipData := map[string]string{
			"nexttooltip": db.GetTranslation("Caption_Next", al),
			"skipforwardtooltip": db.GetTranslation("Caption_QuitLangTest", al),
		}

		if f, err := os.Create(path.Join(basketDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
    }

	itemPlaceholderPattern := regexp.MustCompile("<(\\d*)>")

	for _, basket := range db.GetBaskets() {

		// Render the media markup independently of al
		var mediaMarkupBuffer bytes.Buffer
		var mediaMarkup string
		var rubricMediaType string
		switch basket.MediaType {
			case "text/html":
				if err := tpl.ExecuteTemplate(&mediaMarkupBuffer, "textmedia.gohtml", map[string]any{"markup": template.HTML(basket.TextMedia)}); err != nil {
					log.Fatal(err)
				}
				mediaMarkup = mediaMarkupBuffer.String()
				rubricMediaType = "Text"
			case "audio/mpeg":
				if err := tpl.ExecuteTemplate(&mediaMarkupBuffer, "audiomedia.gohtml", map[string]string{"filename": basket.FileMedia}); err != nil {
					log.Fatal(err)
				}
				mediaMarkup = mediaMarkupBuffer.String()
				rubricMediaType = "Sound"
		  	case "image/jpeg":
				if err := tpl.ExecuteTemplate(&mediaMarkupBuffer, "imagemedia.gohtml", map[string]string{"filename": basket.FileMedia}); err != nil {
					log.Fatal(err)
				}
				mediaMarkup = mediaMarkupBuffer.String()
				rubricMediaType = "Image"
			case "none":
				mediaMarkup = ""
			  	rubricMediaType = "NoMedia"
		}

		basketType := basket.Type
		basketId := basket.Id
		basketPrompt := basket.Prompt

		var basketMarkupBuffer bytes.Buffer
      	var responseMarkup string
		var numberOfItems int
		switch basket.Type {
			case "mcq":
				item := db.GetItemsForBasket(basketId)[0]
				answers := db.GetAnswersForItem(item.Id)
				data := map[string]any{
					"itemText": template.HTML(item.Text),
					"itemId": item.Id,
					"positionInBasket": item.Position,
					"answers": answers,
				}
				if err := tpl.ExecuteTemplate(&basketMarkupBuffer, "mcqresponse.gohtml", data); err != nil {
					log.Fatal(err)
				}
				numberOfItems = 1
		  	case "shortanswer":
				items := db.GetItemsForBasket(basketId)
				data := map[string]any{"basketPrompt": template.HTML(basketPrompt), "items": items}
				if err := tpl.ExecuteTemplate(&basketMarkupBuffer, "saresponse.gohtml", data); err != nil {
					log.Fatal(err)
				}
				numberOfItems = len(items)
			case "gaptext":
				gapText := basket.GapText
				gapMarkup := gapText
				items := db.GetItemsForBasket(basketId)

				matches := itemPlaceholderPattern.FindAllString(gapText, -1)
				if matches == nil {
					log.Fatal("No matches found for pattern: " + itemPlaceholderPattern.String())
				}

				for i, itemNumber := range matches {
					item := items[i]
				  	gapMarkup = strings.Replace(gapMarkup, itemNumber,"<input type=\"text\" name=\"" + strconv.Itoa(item.Id) + "-response\" />", 1)
				}

				data := map[string]any{"basketPrompt": basketPrompt, "items": items, "markup": template.HTML(gapMarkup)}
				if err := tpl.ExecuteTemplate(&basketMarkupBuffer, "gtresponse.gohtml", data); err != nil {
					log.Fatal(err)
				}
				numberOfItems = len(items)
			case "gapdrop":
				gapText := basket.GapText
				gapMarkup := gapText

				items := db.GetItemsForBasket(basketId)

				matches := itemPlaceholderPattern.FindAllString(gapText, -1)
				if matches == nil {
					log.Fatal("No matches found for pattern: " + itemPlaceholderPattern.String())
				}

				for i, itemNumber := range matches {
					item := items[i]
					answers := db.GetAnswersForItem(item.Id)
					selectMarkup := "<select name=\"" + strconv.Itoa(item.Id) + "-response\">"
              		selectMarkup += "<option></option>"
					for _, answer := range answers {
                		selectMarkup += "<option value=\"" + strconv.Itoa(answer.Id) + "\">" + answer.Text + "</option>"
					}
					selectMarkup += "</select>"
					gapMarkup = strings.Replace(gapMarkup, itemNumber, selectMarkup, 1)
            	}

				data := map[string]any{"basketPrompt": basketPrompt, "items": items, "markup": template.HTML(gapMarkup)}
				if err := tpl.ExecuteTemplate(&basketMarkupBuffer, "gdresponse.gohtml", data); err != nil {
					log.Fatal(err)
				}
				numberOfItems = len(items)
			case "tabbedpane":
				// This is a testlet and will contain child baskets
				childBaskets := []map[string]any{}
				for _, childBasket := range db.GetChildBasketsForBasket(basketId) {
					childBasketId := childBasket.Id
					// NOTE: In the future there may be multi item baskets in a testlet. At the moment
					// there are mainly MCQ baskets and a gap drop with one item.
					item := db.GetItemsForBasket(childBasketId)[0]
					childBaskets = append(childBaskets, map[string]any{
						"basketId": strconv.Itoa(childBasketId),
						"itemId": strconv.Itoa(item.Id),
						"itemText": template.HTML(item.Text),
						"positionInBasket": childBasket.ParentTestletPosition,
						"answers": db.GetAnswersForItem(item.Id),
				  	})
				}
				data := map[string]any{"childBaskets": childBaskets}
				if err := tpl.ExecuteTemplate(&basketMarkupBuffer, "tabbedpaneresponse.gohtml", data); err != nil {
					log.Fatal(err)
				}
				numberOfItems = len(childBaskets)
        }
		responseMarkup = basketMarkupBuffer.String()

		for _, al := range db.AdminLocales {

			alDir := path.Join(basketDir, al)
			if err := os.MkdirAll(alDir, 0777); err != nil {
			  log.Fatal(err)
			}

			rubricText := db.GetTranslation("LangTest_Rubric#" + typeMap[basketType] + "#" + rubricMediaType,al)

			// Confirmation dialog texts.
			warningText := db.GetTranslation("Dialogues_SkipLangTest",al)
			yes := db.GetTranslation("Caption_Yes",al)
			no := db.GetTranslation("Caption_No",al)
			yourAnswerTitle := db.GetTranslation("LangTest_ItemFeedback_YourAnswer",al)
			correctAnswerTitle := db.GetTranslation("LangTest_ItemFeedback_CorrectAnswer",al)
			correctAnswersTitle := db.GetTranslation("LangTest_ItemFeedback_CorrectAnswers",al)

			data := map[string]any{
				"basketId": strconv.Itoa(basketId),
				"basketType": basketType,
				"rubricText": template.HTML(rubricText),
				"warningText": warningText,
				"yes": yes,
				"no": no,
				"yourAnswerTitle": yourAnswerTitle,
				"correctAnswerTitle": correctAnswerTitle,
				"correctAnswersTitle": correctAnswersTitle,
				"mediaMarkup": template.HTML(mediaMarkup),
				"responseMarkup": template.HTML(responseMarkup),
				"numberOfItems": numberOfItems,
			}

			if f, err := os.Create(path.Join(alDir, strconv.Itoa(basketId) + ".html")); err == nil {
				if err := tpl.ExecuteTemplate(f, "basket.gohtml", data); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		} // admin language iterator
    }
}

func ExportEndOfTestPages(baseDir string) {

	endOfTestDir := path.Join(baseDir, "endoftest")
	if err := os.MkdirAll(endOfTestDir, 0777); err != nil {
		log.Fatalf("Failed to create %v: %w", endOfTestDir, err)
	}

	tpl := template.Must(template.ParseFiles(path.Join("templates", "endoftest.gohtml"), path.Join("templates", "toolbartooltips.gojson")))

	for _, al := range db.AdminLocales {

		tipData := map[string]string{"nexttooltip": db.GetTranslation("Caption_Feedback", al)}
		if f, err := os.Create(path.Join(endOfTestDir, al + "-toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatalf("Failed to render tooltips: %w", err)
			}
		} else {
			log.Fatalf("Failed to create endoftest tooltip file for al %v: %w", al, err)
		}

		data := map[string]any{
			"title": db.GetTranslation("Title_LangTestEnd", al),
			"text": template.HTML(db.GetTranslation("LangTestEnd_Text", al)),
		}
		if f, err := os.Create(path.Join(endOfTestDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "endoftest.gohtml", data); err != nil {
				log.Fatalf("Failed to render endoftest template: %w", err)
			}
		} else {
			log.Fatalf("Failed to create endoftest file for al %v: %w", al, err)
		}
    }
}

