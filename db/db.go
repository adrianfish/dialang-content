package db

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"database/sql"
	"html/template"
	"github.com/dialangproject/common/models"
	_ "github.com/lib/pq"
	"github.com/magiconair/properties"
)

var db *sql.DB

var adminTexts = map[string]*properties.Properties{}
var AdminLocales = []string{}
var SaSkills = []string{"Reading","Writing","Listening"}
var TestSkills = []string{"Reading","Writing","Listening","Structures","Vocabulary"}
var AdvfbSkills = []string{"Reading","Writing","Listening"}
var ItemLevels = []string{"A1", "A2", "B1", "B2", "C1", "C2"}
var itemMap = map[int]models.Item{}
var answerMap = map[int]models.Answer{}
//var isoLangMappingsStmt *sql.Stmt

func init() {

	connStr := "postgres://dialangadmin:dialangadmin@localhost/dialang?sslmode=disable"
	thisDb, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = thisDb

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to dialang database")
	}

	/*
	if stmt, err := db.Prepare("SELECT locale, two_letter_lang FROM admin_languages"); err == nil {
		isoLangMappingsStmt = stmt
	} else {
		log.Fatal(err)
	}
	*/

	for _, lang := range GetAdminLanguages() {
		AdminLocales = append(AdminLocales, lang.Locale)
	}

	for _, al := range AdminLocales {
		adminTexts[al] = properties.MustLoadFile("admin-texts/admintexts_" + al + ".properties", properties.UTF8)
	}

}

func GetTestLanguageCodes() []models.TestLanguage {

	rows, err := db.Query("SELECT locale, two_letter_code FROM test_languages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	languages := []models.TestLanguage{}
	for rows.Next() {
		var language models.TestLanguage
		if err := rows.Scan(&language.Locale, &language.TwoLetterCode); err != nil {
			log.Fatalf("Failed to scan test language: %w", err)
		}
		languages = append(languages, language)
	}
	return languages
}

func GetTestLanguagePrompts(al string) []map[string]string {

	prompts := []map[string]string{}
	if p, ok := adminTexts[al]; ok {
		// Grab all the pairs matching the choose string
		if chooseProps, err := p.Filter("^ChooseTest_Language.*"); err == nil {
			for k, v := range chooseProps.Map() {
				prompts = append(prompts,
					map[string]string{"languageCode": strings.Split(k, "#")[1], "languageName": v})
			}
		} else {
			log.Fatal("No test choice props for lang " + al)
		}
    } else {
		log.Fatal("No texts for lang " + al)
	}

	return prompts
}
func GetVSPTLevels() []string {

	rows, err := db.Query("SELECT level FROM vsp_levels")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	levels := []string{}

	for rows.Next() {
		var level string
		if err := rows.Scan(&level); err != nil {
			log.Fatal(err)
		}
		levels = append(levels, level)

	}
	return levels
}

func GetVSPTWords(tl string) []models.VSPTWord {

	rows, err := db.Query("SELECT word,words.word_id AS id,valid, weight FROM vsp_test_word,words WHERE locale = $1 AND vsp_test_word.word_id = words.word_id", tl)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	words := []models.VSPTWord{}
	for rows.Next() {
		var word models.VSPTWord
		if err := rows.Scan(&word.Word, &word.WordId, &word.Valid, &word.Weight); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	return words
}

func GetVSPTBands() []models.VSPTBand {

	rows, err := db.Query("SELECT * from vsp_bands")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bands := []models.VSPTBand{}
	for rows.Next() {
		var band models.VSPTBand
		if err := rows.Scan(&band.Locale, &band.Level, &band.Low, &band.High); err != nil {
			log.Fatal(err)
		}
		bands = append(bands, band)
	}

	return bands
}

func GetSAWeights() []models.SAWeight {

	rows, err := db.Query("SELECT * from sa_weights")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	weights := []models.SAWeight{}
	for rows.Next() {
		var weight models.SAWeight
		if err := rows.Scan(&weight.Skill, &weight.WordId, &weight.Weight); err != nil {
			log.Fatal(err)
		}
		weights = append(weights, weight)
	}

	return weights
}

func GetSAGrades() []models.SAGrade {

	rows, err := db.Query("SELECT * from sa_grading")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	grades := []models.SAGrade{}
	for rows.Next() {
		var grade models.SAGrade
		if err := rows.Scan(&grade.Skill, &grade.Rsc, &grade.Ppe, &grade.Se, &grade.Grade); err != nil {
			log.Fatal(err)
		}
		grades = append(grades, grade)
	}

	return grades
}

func GetPreestWeights() map[string]models.PreestWeight {

	rows, err := db.Query("SELECT * from preest_weights")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	weights := map[string]models.PreestWeight{}

	for rows.Next() {
		var tl, skill string
		var saTaken, vsptTaken int
		var weight models.PreestWeight
		if err := rows.Scan(&tl, &skill, &saTaken, &vsptTaken, &weight.Sa, &weight.Vspt, &weight.Coe); err != nil {
			log.Fatal(err)
		}
		key := fmt.Sprintf("%v#%v#%v#%v", tl, skill, saTaken, vsptTaken)
		weights[key] = weight
	}

	return weights
}

func GetPreestAssignments() map[string][]models.PreestAssignment {

	rows, err := db.Query("SELECT * from preest_assignments")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	assignmentsMap := map[string][]models.PreestAssignment{}

	for rows.Next() {
		var tl, skill string
		var assignment models.PreestAssignment
		if err := rows.Scan(&tl, &skill, &assignment.Pe, &assignment.BookletId); err != nil {
			log.Fatal(err)
		}
		key := fmt.Sprintf("%v#%v", tl, skill)
		if assignments, ok := assignmentsMap[key]; !ok {
			assignmentsMap[key] = []models.PreestAssignment{assignment}
		} else {
			assignmentsMap[key] = append(assignments, assignment)
		}
	}

	return assignmentsMap
}

func GetBookletIds() []int {

	rows, err := db.Query("SELECT id from booklets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ids := []int{}

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	return ids
}

func GetBaskets() []models.Basket {

	rows, err := db.Query("SELECT * FROM baskets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return getBasketsForRows(rows)
}

func GetBasketsForBooklet(bookletId int) []models.Basket {

	rows, err := db.Query("SELECT baskets.* FROM baskets, booklet_basket WHERE booklet_basket.booklet_id = $1 AND booklet_basket.basket_id = baskets.id", bookletId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return getBasketsForRows(rows)
}

func GetChildBasketsForBasket(basketId int) []models.Basket {

	rows, err := db.Query("SELECT * FROM baskets WHERE parent_testlet_id = $1", basketId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return getBasketsForRows(rows)
}

func GetNumItemsForBasket(basketId int) int {

	rows, err := db.Query("SELECT count(*) as num_items FROM basket_item WHERE basket_id = $1", basketId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var numItems int
	if rows.Next() {
		if err := rows.Scan(&numItems); err != nil {
			log.Fatal(err)
		}
	}

	return numItems
}

func GetBookletBaskets() map[int][]int {

	rows, err := db.Query("SELECT * FROM booklet_basket")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bookletBaskets := map[int][]int{}

	for rows.Next() {
		var bookletId, basketId int
		if err := rows.Scan(&bookletId, &basketId); err != nil {
			log.Fatal(err)
		}
		baskets, ok := bookletBaskets[bookletId]
		if !ok {
			bookletBaskets[bookletId] = []int{basketId}
		} else {
			bookletBaskets[bookletId] = append(baskets, basketId)
		}
	}

	return bookletBaskets
}

func GetItemsForBasket(basketId int) []models.Item {

	rows, err := db.Query("SELECT i.*, bi.position FROM baskets b,basket_item bi,items i WHERE b.id = $1 AND b.id = bi.basket_id AND bi.item_id = i.id ORDER BY position", basketId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	items := []models.Item{}
	for rows.Next() {
		var item models.Item
		var skill sql.Null[string]
		var text sql.Null[string]
		if err := rows.Scan(&item.Id, &item.Type, &skill, &item.SubSkill, &text, &item.Weight, &item.Position); err != nil {
			log.Fatal(err)
		}
		if skill.Valid { item.Skill = skill.V }
		if text.Valid { item.Text = template.HTML(text.V) }
		items = append(items, item)
	}
	return items;
}

func GetAnswers() []models.Answer {

	rows, err := db.Query("SELECT * FROM answers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	answers := []models.Answer{}
	for rows.Next() {
		var answer models.Answer
		if err := rows.Scan(&answer.Id, &answer.ItemId, &answer.Text, &answer.Correct); err != nil {
			log.Fatal(err)
		}
		answers = append(answers, answer)
	}
	return answers
}

func GetAnswersForItem(itemId int) []models.Answer {

	rows, err := db.Query("SELECT * FROM answers WHERE item_id = $1", itemId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	answers := []models.Answer{}
	for rows.Next() {
		var answer models.Answer
		if err := rows.Scan(&answer.Id, &answer.ItemId, &answer.Text, &answer.Correct); err != nil {
			log.Fatal(err)
		}
		answers = append(answers, answer)
	}
	return answers
}

func GetAnswer(answerId int) (*models.Answer, error) {

	if len(answerMap) == 0 {
		if answers, err := getAnswers(); err == nil {
			for _, answer := range answers {
				answerMap[answer.Id] = answer
			}
		} else {
			log.Println("Failed to cache answers in answerMap")
			return nil, err
		}
	}

	if answer, ok := answerMap[answerId]; ok {
		return &answer, nil
	} else {
		return nil, errors.New("Failed to find answer in answerMap")
	}
}

func getAnswers() ([]models.Answer, error) {

	if rows, err := db.Query("SELECT * FROM answers"); err == nil {
		answers := []models.Answer{}
		for rows.Next() {
			var answer models.Answer
			if err := rows.Scan(&answer.Id, &answer.ItemId, &answer.Text, &answer.Correct); err != nil {
				return nil, err
			}
			answers = append(answers, answer)
		}
		return answers, nil
	} else {
		return nil, err
	}
}

func GetItem(itemId int) (*models.Item, error) {

	if len(itemMap) == 0 {
		if items, err := GetItems(); err == nil {
			for _, item := range items {
				itemMap[item.Id] = item
			}
		} else {
			log.Println("Failed to cache items in itemMap")
			return nil, err
		}
	}

	if item, ok := itemMap[itemId]; ok {
		return &item, nil
	} else {
		return nil, errors.New("Failed to find item in itemMap")
	}
}

func GetItems() ([]models.Item, error) {

	if rows, err := db.Query("SELECT * FROM items"); err == nil {
		items := []models.Item{}
		for rows.Next() {
			var item models.Item
			var text sql.Null[string]
			if err := rows.Scan(&item.Id, &item.Type, &item.Skill, &item.SubSkill, &text, &item.Weight); err != nil {
				return nil, err
			} else {
				if text.Valid { item.Text = template.HTML(text.V) }
				items = append(items, item)
			}
		}
		return items, nil
	} else {
		return nil, err
	}
}

func GetPunctuationCharacters() []string {

	chars := []string{}
	if rows, err := db.Query("SELECT * FROM punctuation"); err == nil {
		for rows.Next() {
			var char string
			if err := rows.Scan(&char); err != nil {
				log.Fatal(err)
			}
			chars = append(chars, char)
		}
	} else {
		log.Fatal(err)
	}
	return chars
}

func GetTranslation(key string, lang string) string {

	if props, ok := adminTexts[lang]; ok {
		if v, ok := props.Get(key); ok {
			return v
		}
		fmt.Sprintf("No translation for lang %v and key %v. Will fall back to eng_gb ...\n", lang, key)
	} else {
		fmt.Sprintf("No translation for lang %v and key %v. Will fall back to eng_gb ...\n", lang, key)
	}

	if props, ok := adminTexts["eng_gb"]; ok {
		if v, ok := props.Get(key); ok {
			return v
		}
	}
	fmt.Sprintf("No translation for eng_gb and key %v. Will return %v ...\n", lang, key, key)
	return key
}

func GetTranslationLike(keyExpression string, language string) string {

	parts := strings.Split(keyExpression, "%")

	langTexts, ok := adminTexts[language]
	if !ok {
        log.Fatalf("Language: '%v' not found", language)
	}

	matchedProps := langTexts.FilterFunc(func(k, v string) bool {
		return strings.HasPrefix(k, parts[0]) && strings.HasSuffix(k, parts[1])
	})

	return matchedProps.MustGet(matchedProps.Keys()[0])
}

func GetAdminLanguages() []models.AdminLanguage {

	rows, err := db.Query("SELECT locale, description FROM admin_languages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	languages := []models.AdminLanguage{}

	for rows.Next() {
		var language models.AdminLanguage
		if err := rows.Scan(&language.Locale, &language.Description); err != nil {
			log.Fatal(err)
		}
		languages = append(languages, language)
	}

	return languages
}

func GetSAStatements(al string, skill string) []models.SAStatement {

	rows, err := db.Query("SELECT * FROM sa_statements WHERE locale = $1 AND skill = $2 ORDER BY wid", al, skill)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	list := []models.SAStatement{}

    for rows.Next() {
		var s models.SAStatement
		if err := rows.Scan(&s.Locale, &s.Skill, &s.WordId, &s.Statement); err != nil {
			log.Fatal(err)
		}
		list = append(list, s)
    }
	return list
}

func GetItemGrades() ([]models.ItemGrade, error) {

	rows, err := db.Query("SELECT * FROM item_grading")
	if err != nil {
		log.Printf("Failed to get item grades: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	itemGrades := []models.ItemGrade{}

	for rows.Next() {
		var itemGrade models.ItemGrade
		if err := rows.Scan(&itemGrade.Tl, &itemGrade.Skill, &itemGrade.BookletId, &itemGrade.RawScore, &itemGrade.PPE, &itemGrade.SE, &itemGrade.Grade); err == nil {
			itemGrades = append(itemGrades, itemGrade)
		} else {
			log.Printf("Failed to scan item grade: %s", err)
			return nil, err
		}
	}

	return itemGrades, nil
}

func GetSubSkills(al string) []map[string]string {

	langTexts, ok := adminTexts[al]
	if !ok {
        log.Fatalf("Failed to get admin texts for al %v.", al)
	}

	matchedProps := langTexts.FilterFunc(func(k, v string) bool {
		return strings.HasPrefix(k, "Subskill#")
	})

	subskills := []map[string]string{}
	for k, v := range matchedProps.Map() {
		subskills = append(subskills, map[string]string{"code": strings.Split(k, "#")[1], "description": v})
	}

	return subskills
}

/*
func GetTwoLetterIsoLangLookup() {

	rows, err := isoLangMappingsStmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	map := map[string]string{}

	for rows.Next() {
		var locale, twoLetterLang string
		if err := rows.Scan(&locale, &twoLetterLang); err != nil {
			log.Fatal(err)
		}
		map[locale] = twoLetterLang
	}

    val rs = isoLangMappingsST.executeQuery
    val map = new HashMap[String, String]
    while (rs.next) {
      map += (rs.getString("locale") -> rs.getString("two_letter_lang"))
    }
    rs.close
    map.toMap
  }
  */

func getBasketsForRows(rows *sql.Rows) []models.Basket {

	baskets := []models.Basket{}

	for rows.Next() {
		var b models.Basket
		var testletId, testletPosition sql.Null[int]
		var skill, label, prompt, gapText, fileMedia, textMedia sql.Null[string]

		if err := rows.Scan(&b.Id, &b.Type, &testletId, &testletPosition, &skill, &label, &prompt, &gapText, &b.Weight, &b.MediaType, &textMedia, &fileMedia); err != nil {
			log.Fatal(err)
		}

		if testletId.Valid { b.ParentTestletId = testletId.V }
		if testletPosition.Valid { b.ParentTestletPosition = testletPosition.V }
		if skill.Valid { b.Skill = skill.V }
		if label.Valid { b.Label = label.V }
		if prompt.Valid { b.Prompt = prompt.V }
		if gapText.Valid { b.GapText = gapText.V }
		if fileMedia.Valid { b.FileMedia = fileMedia.V }
		if textMedia.Valid { b.TextMedia = textMedia.V }

		baskets = append(baskets, b)
	}
	return baskets;
}
