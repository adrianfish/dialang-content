package exporters

import (
	"html/template"
	"log"
	"os"
	"strings"
	"path"
	"github.com/dialangproject/common/db"
	"github.com/pariz/gountries"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var tags []language.Tag
var dicts map[string]*display.Dictionary

func init() {

	tags = []language.Tag{
		language.English,
		language.German,
		language.Greek,
		language.Spanish,
		language.Finnish,
		language.French,
		language.Icelandic,
		language.Italian,
		language.Dutch,
		language.Norwegian,
		language.Portuguese,
		language.Swedish,
		language.Danish,
		language.Chinese,
		language.Indonesian,
	}

	dicts = map[string]*display.Dictionary{
		"deu_de": display.German,
		"ell_gr": display.Greek,
		"eng_gb": display.English,
		"spa_es": display.Spanish,
		"fin_fi": display.Finnish,
		"fra_fr": display.French,
		"isl_is": display.Icelandic,
		"ita_it": display.Italian,
		"nld_nl": display.Dutch,
		"nor_no": display.Norwegian,
		"por_pt": display.Portuguese,
		"swe_se": display.Swedish,
		"dan_dk": display.Danish,
		"cmn_cn": display.Chinese,
		"ind_id": display.Indonesian,
	}
}

func ExportQuestionnairePages(baseDir string) {

	questionnaireDir := path.Join(baseDir, "questionnaire")
	if err := os.MkdirAll(questionnaireDir, 0777); err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles(
		path.Join("templates", "questionnaire.gohtml"),
		path.Join("templates", "toolbartooltips.gojson")))

	allCountries := gountries.New().FindAllCountries()

	for _, al := range db.AdminLocales {
		// Split the code into language and country
		parts := strings.Split(al, "_")

		languages := []map[string]string{}
		n := dicts[al].Languages()
		for _, al := range db.AdminLocales {
			languages = append(languages, map[string]string{"name": n.Name(language.Make(al)), "isocode": al})
		}

		languageCode := strings.ToUpper(parts[0])
		var countries []map[string]string
		for code, country := range allCountries {
			name := country.Translations[languageCode].Common
			if name == "" {
				name = country.Name.Common
			}
			countries = append(countries, map[string]string{"name": name, "isocode": strings.ToLower(code)})
		}
		//languageMatches := language.MatchStrings(matcher, al)
      	//languages = availableLanguages.map(l => Map("name" -> l.getDisplayLanguage(currentLocale), "isocode" -> l.getLanguage))
		data := map[string]any{
			"title": db.GetTranslation("Questionnaire_Title", al),
			"information": template.HTML(db.GetTranslation("Questionnaire_Information", al)),
			"age": db.GetTranslation("Questionnaire_Age", al),
			"gender": db.GetTranslation("Questionnaire_Gender", al),
			"gendermale": db.GetTranslation("Questionnaire_Gender_Male", al),
			"genderfemale": db.GetTranslation("Questionnaire_Gender_Female", al),
			"prefernottosay": db.GetTranslation("Questionnaire_PreferNot", al),
			"other": db.GetTranslation("Questionnaire_Other", al),
			"firstlanguage": db.GetTranslation("Questionnaire_FirstLanguage", al),
			"nationality": db.GetTranslation("Questionnaire_Nationality", al),
			"languages": languages,
			"countries": countries,
			"institution": db.GetTranslation("Questionnaire_Institution", al),
			"reason": db.GetTranslation("Questionnaire_Reason", al),
			"reasonplacement": db.GetTranslation("Questionnaire_Reason_Placement", al),
			"reasonimprovement": db.GetTranslation("Questionnaire_Reason_Improvement", al),
			"reasonresearchproject": db.GetTranslation("Questionnaire_Reason_ResearchProject", al),
			"reasondialangresearch": db.GetTranslation("Questionnaire_Reason_DialangResearch", al),
			"reasonemployer": db.GetTranslation("Questionnaire_Reason_Employer", al),
			"reasonother": db.GetTranslation("Questionnaire_Reason_Other", al),
			"accuracy": db.GetTranslation("Questionnaire_Accuracy", al),
			"accuracyterrible": db.GetTranslation("Questionnaire_Accuracy_Terrible", al),
			"accuracybad": db.GetTranslation("Questionnaire_Accuracy_Bad", al),
			"accuracyokay": db.GetTranslation("Questionnaire_Accuracy_Okay", al),
			"accuracygood": db.GetTranslation("Questionnaire_Accuracy_Good", al),
			"accuracyexcellent": db.GetTranslation("Questionnaire_Accuracy_Excellent", al),
			"comments": db.GetTranslation("Questionnaire_Comments", al),
			"email": db.GetTranslation("Questionnaire_Email", al),
		}

		if f, err := os.Create(path.Join(questionnaireDir, al + ".html")); err == nil {
			if err := tpl.ExecuteTemplate(f, "questionnaire.gohtml", data); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}

		tipData := map[string]string{
			"backtooltip": db.GetTranslation("Caption_BacktoFeedback", al),
			"skipforwardtooltip": db.GetTranslation("Caption_SkipQuestionnaire", al),
			"nexttooltip": db.GetTranslation("Caption_SubmitQuestionnaire", al),
		}

		if f, err := os.Create(path.Join(questionnaireDir, al + "toolbarTooltips.json")); err == nil {
			if err := tpl.ExecuteTemplate(f, "toolbartooltips.gojson", tipData); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}

/*
func getLanguageDisplayName(languageCode string) string {

	displayName := display.LanguageDisplayName(languageCode)
	if displayName == "" {
		displayName = languageCode
	}
	return displayName
}
*/
