package exporters

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"github.com/adrianfish/dialang/content/db"
	"github.com/adrianfish/dialang/content/models"
)

func exportVSPTData(baseDir string) error {

	tlWords := []map[string]any{}
	for _, tl := range db.GetTestLanguageCodes() {
		tlWords = append(tlWords, map[string]any{"key": "data:vspt-words:" + tl.Locale, "value": convertToString(db.GetVSPTWords(tl.Locale))})
	}
	createEncoder(baseDir, "vspt-words.json").Encode(tlWords)

	tlBands := map[string][]models.VSPTBand{}
	for _, band := range db.GetVSPTBands() {
		key := "data:vspt-bands:" + band.Locale
		if val, ok := tlBands[key]; ok {
			tlBands[key] = append(val, band)
		} else {
			tlBands[key] = []models.VSPTBand{band}
		}
	}

	tlBandsSlice := []map[string]any{}
	for key, value := range tlBands {
		tlBandsSlice = append(tlBandsSlice, map[string]any{"key": key, "value": convertToString(value)})
	}
	createEncoder(baseDir, "vspt-bands.json").Encode(tlBandsSlice)

	return nil
}

func exportSAData(baseDir string) error {

	saWeights := map[string]map[string]int{}
	for _, weight := range db.GetSAWeights() {
		key := "data:sa-weights:" + weight.Skill
		if _, ok := saWeights[key]; ok {
			saWeights[key][weight.WordId] = weight.Weight
		} else {
			saWeights[key] = map[string]int{weight.WordId: weight.Weight}
		}
	}

	saWeightsSlice := []map[string]any{}
	for key, value := range saWeights {
		saWeightsSlice = append(saWeightsSlice, map[string]any{"key": key, "value": convertToString(value)})
	}
	createEncoder(baseDir, "sa-weights.json").Encode(saWeightsSlice)

	saGrades := []map[string]any{}
	for _, grade := range db.GetSAGrades() {
		key := "data:sa-grades:" + grade.Skill + ":" + strconv.Itoa(grade.Rsc)
		saGrades = append(saGrades, map[string]any{"key": key, "value": convertToString(grade)})
	}
	createEncoder(baseDir, "sa-grading.json").Encode(saGrades)

	return nil
}

func exportPreestData(baseDir string) error {

	preestWeights := []map[string]any{}
	for key, weight := range db.GetPreestWeights() {
		preestWeights = append(preestWeights, map[string]any{"key": "data:preest-weights:" + key, "value": convertToString(weight)})
	}
	createEncoder(baseDir, "preest-weights.json").Encode(preestWeights)


	preestAssignments := []map[string]any{}
	for key, assignments := range db.GetPreestAssignments() {
		assignmentsSlice := []models.PreestAssignment{}
		for _, assignment := range assignments {
			assignmentsSlice = append(assignmentsSlice, assignment)
		}
		preestAssignments = append(preestAssignments, map[string]any{"key": "data:preest-assignments:" + key, "value": convertToString(assignmentsSlice)})
	}
	createEncoder(baseDir, "preest-assignments.json").Encode(preestAssignments)

	return nil
}

func exportBookletData(baseDir string) error {

	lengths := []map[string]any{}
	for _, bookletId := range db.GetBookletIds() {
		var total int
		for _, basket := range db.GetBasketsForBooklet(bookletId) {
			if basket.Type == "tabbedpane" {
				for _, childBasket := range db.GetChildBasketsForBasket(basket.Id) {
					numItems := db.GetNumItemsForBasket(childBasket.Id)
					total += numItems
				}
			} else {
				total += db.GetNumItemsForBasket(basket.Id)
			}
		}
		lengths = append(lengths, map[string]any{"key": "data:booklet-lengths:" + strconv.Itoa(bookletId), "value": strconv.Itoa(total)})
	}
	createEncoder(baseDir, "booklet-lengths.json").Encode(lengths)

	bookletBaskets := []map[string]any{}
	for bookletId, baskets := range db.GetBookletBaskets() {
		bookletBaskets = append(bookletBaskets, map[string]any{"key": "data:booklet-baskets:" + strconv.Itoa(bookletId), "value": convertToString(baskets)});
	}
	createEncoder(baseDir, "booklet-baskets.json").Encode(bookletBaskets)

	return nil
}

func exportItemAnswers(baseDir string) {

	itemAnswers:= map[int][]models.Answer{}
	for _, answer := range db.GetAnswers() {
		if answers, ok := itemAnswers[answer.ItemId]; ok {
			itemAnswers[answer.ItemId] = append(answers, answer)
		} else {
			itemAnswers[answer.ItemId] = []models.Answer{answer}
		}
	}

	itemAnswersSlice := []map[string]any{}
	for key, value := range itemAnswers {
		itemAnswersSlice = append(itemAnswersSlice, map[string]any{"key": key, "value": convertToString(value)})
	}
	createEncoder(baseDir, "item-answers.json").Encode(itemAnswersSlice)
}

func exportAnswers(baseDir string) {

	answers:= []map[string]any{}
	for _, answer := range db.GetAnswers() {
		answers = append(answers, map[string]any{"key": "data:answers:" + strconv.Itoa(answer.Id), "value": convertToString(answer)})
	}
	createEncoder(baseDir, "answers.json").Encode(answers)
}

func exportItems(baseDir string) {

	items:= []map[string]any{}
	if dbItems, err := db.GetItems(); err == nil {
		for _, item := range dbItems {
			items = append(items, map[string]any{"key": "data:items:" + strconv.Itoa(item.Id), "value": convertToString(item)})
		}
	} else {
		log.Fatal(err)
	}
	createEncoder(baseDir, "items.json").Encode(items)
}

func exportItemGrades(baseDir string) {

	gradesMap := map[string]map[int]models.ItemGrade{}
	if grades, err := db.GetItemGrades(); err == nil {
		for _, grade := range grades {
			key := fmt.Sprintf("%v#%v#%v", grade.Tl, grade.Skill, grade.BookletId)
			if _, ok := gradesMap[key]; !ok {
				gradesMap[key] = map[int]models.ItemGrade{}
			}
			gradesMap[key][grade.RawScore] = grade
		}
	} else {
		log.Fatal(err)
	}
	gradesSlice := []map[string]any{}
	for gradeKey, value := range gradesMap {
		for rawScore, grade := range value {
			key := "data:item-grades:" + gradeKey + ":" + strconv.Itoa(rawScore)
			gradesSlice = append(gradesSlice, map[string]any{"key": key, "value": convertToString(grade)})
		}
	}
	createEncoder(baseDir, "item-grades.json").Encode(gradesSlice)
}

func exportPunctuation(baseDir string) {

	punctuation := []map[string]any{
		map[string]any{"key": "data:punctuation", "value": convertToString(db.GetPunctuationCharacters())},
	}
	createEncoder(baseDir, "punctuation.json").Encode(punctuation)
}

func ExportWebData(baseDir string) {

	exportVSPTData(baseDir)
	exportSAData(baseDir)
	exportPreestData(baseDir)
	exportBookletData(baseDir)
	exportItemAnswers(baseDir)
	exportAnswers(baseDir)
	exportItems(baseDir)
	exportPunctuation(baseDir)
	exportItemGrades(baseDir)
}

func createEncoder(baseDir string, name string) *json.Encoder {

	f, err := os.Create(path.Join(baseDir, name))
	if err != nil {
		log.Fatalf("Failed to create %s: %s\n", name, err)
	}
	jsonEncoder := json.NewEncoder(f)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder
}

func convertToString(thing any) string {

	b, err := json.Marshal(thing)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
