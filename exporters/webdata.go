package exporters

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"github.com/adrianfish/dialang-content/db"
	"github.com/adrianfish/dialang-common/models"
)

func exportVSPTData(baseDir string) error {

	wordsFile, err := os.OpenFile(baseDir + "/vspt-words.csv", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	wordsWriter := csv.NewWriter(wordsFile)
	if err := wordsWriter.Write([]string{"test_language", "word_id", "word", "valid", "weight"}); err != nil {
		return err
	}

	for _, tl := range db.GetTestLanguageCodes() {
		for _, word := range db.GetVSPTWords(tl.Locale) {
			if err := wordsWriter.Write([]string{tl.Locale, word.WordId, word.Word, strconv.Itoa(word.Valid), strconv.Itoa(word.Weight)}); err != nil {
				return err
			}
		}
	}

	wordsWriter.Flush()

	bandsFile, err := os.OpenFile(baseDir + "/vspt-bands.csv", os.O_RDWR | os.O_CREATE, 0666)

	bandsWriter := csv.NewWriter(bandsFile)

	if err := bandsWriter.Write([]string{"test_language", "level", "low", "high"}); err != nil {
		return err
	}

	for _, band := range db.GetVSPTBands() {
		if err := bandsWriter.Write([]string{band.Locale, band.Level, strconv.Itoa(band.Low), strconv.Itoa(band.High)}); err != nil {
			return err
		}
	}

	bandsWriter.Flush()

	return nil
}

func exportSAData(baseDir string) error {

	weightsFile, err := os.Create(path.Join(baseDir, "sa-weights.csv"))
	if err != nil {
		return err
	}

	weightsWriter := csv.NewWriter(weightsFile)
	if err := weightsWriter.Write([]string{"skill", "wid", "weight"}); err != nil {
		return err
	}

	for _, weight := range db.GetSAWeights() {
		weightsWriter.Write([]string{weight.Skill, weight.WordId, strconv.Itoa(weight.Weight)})
	}

	weightsWriter.Flush()

	gradesFile, err := os.Create(path.Join(baseDir, "sa-grading.csv"))
	if err != nil {
		return err
	}

	gradesWriter := csv.NewWriter(gradesFile)
	if err := gradesWriter.Write([]string{"skill", "rsc", "ppe", "se", "grade"}); err != nil {
		return err
	}

	for _, grade := range db.GetSAGrades() {
		gradesWriter.Write([]string{grade.Skill, strconv.Itoa(grade.Rsc), strconv.FormatFloat(grade.Ppe, 'g', -1, 64), strconv.FormatFloat(grade.Se, 'g', -1, 64), strconv.Itoa(grade.Grade)})
	}

	gradesWriter.Flush()

	return nil
}

func exportPreestData(baseDir string) error {

	weightsFile, err := os.Create(path.Join(baseDir, "preest-weights.csv"))
	if err != nil {
		return err
	}

	weightsWriter := csv.NewWriter(weightsFile)
	if err := weightsWriter.Write([]string{"key", "sa", "vspt", "coe"}); err != nil {
		return err
	}

	for key, weight := range db.GetPreestWeights() {
		weightsWriter.Write([]string{
			key,
			strconv.FormatFloat(weight.Sa, 'g', -1, 64),
			strconv.FormatFloat(weight.Vspt, 'g', -1, 64),
			strconv.FormatFloat(weight.Coe, 'g', -1, 64),
		})
	}

	weightsWriter.Flush()

	assignmentsFile, err := os.Create(path.Join(baseDir, "preest-assignments.csv"))
	if err != nil {
		return err
	}

	assignmentsWriter := csv.NewWriter(assignmentsFile)
	if err := assignmentsWriter.Write([]string{"key", "pe", "booklet_id"}); err != nil {
		return err
	}

	for key, assignments := range db.GetPreestAssignments() {
		for _, assignment := range assignments {
			assignmentsWriter.Write([]string{
				key,
				strconv.FormatFloat(assignment.Pe, 'g', -1, 64),
				strconv.Itoa(assignment.BookletId),
			})
		}
	}

	assignmentsWriter.Flush()

	return nil
}

func exportBookletData(baseDir string) error {

	lengthsFile, err := os.Create(path.Join(baseDir, "booklet-lengths.csv"))
	if err != nil {
		return err
	}

	lengthsWriter := csv.NewWriter(lengthsFile)
	if err := lengthsWriter.Write([]string{"booklet_id", "length"}); err != nil {
		return err
	}

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

		lengthsWriter.Write([]string{strconv.Itoa(bookletId), strconv.Itoa(total)})
	}

	lengthsWriter.Flush()

	basketsFile, err := os.Create(path.Join(baseDir, "booklet-baskets.csv"))
	if err != nil {
		return err
	}

	basketsWriter := csv.NewWriter(basketsFile)
	if err := basketsWriter.Write([]string{"booklet_id", "basket_ids"}); err != nil {
		return err
	}

	for bookletId, baskets := range db.GetBookletBaskets() {
		basketsStrings := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(baskets)), ","), "[]")
		basketsWriter.Write([]string{strconv.Itoa(bookletId), basketsStrings})
	}

	basketsWriter.Flush()

	return nil
}

func exportItemAnswers(baseDir string) {

	itemAnswers:= map[int][]models.Answer{}
	for _, answer := range db.GetAnswers() {
		if answers, ok := itemAnswers[answer.ItemId]; ok {
			itemAnswers[answer.ItemId] = append(answers, answer)
		} else {
			itemAnswers[answer.ItemId] = []models.Answer{}
		}
	}

	if f, err := os.Create(path.Join(baseDir, "item-answers.json")); err != nil {
		log.Fatalf("Failed to create item-answers.json: %s\n", err)
	} else {
		defer f.Close()
		json.NewEncoder(f).Encode(itemAnswers)
	}
}

func exportAnswers(baseDir string) {

	answers:= map[int]models.Answer{}
	for _, answer := range db.GetAnswers() {
		answers[answer.Id] = answer
	}

	if f, err := os.Create(path.Join(baseDir, "answers.json")); err != nil {
		log.Fatalf("Failed to create answers.json: %s\n", err)
	} else {
		defer f.Close()
		json.NewEncoder(f).Encode(answers)
	}
}

func exportItems(baseDir string) {

	items:= map[int]models.Item{}
	if dbItems, err := db.GetItems(); err == nil {
		for _, item := range dbItems {
			items[item.Id] = item
		}
	} else {
		log.Fatal(err)
	}

	if f, err := os.Create(path.Join(baseDir, "items.json")); err != nil {
		log.Fatalf("Failed to create items.json: %s\n", err)
	} else {
		defer f.Close()
		json.NewEncoder(f).Encode(items)
	}
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

	if f, err := os.Create(path.Join(baseDir, "item-grades.json")); err != nil {
		log.Fatalf("Failed to create items.json: %s\n", err)
	} else {
		defer f.Close()
		json.NewEncoder(f).Encode(gradesMap)
	}
}

func exportPunctuation(baseDir string) {

	chars := db.GetPunctuationCharacters()
	if f, err := os.Create(path.Join(baseDir, "punctuation.json")); err != nil {
		log.Fatalf("Failed to create punctuation.json: %s\n", err)
	} else {
		defer f.Close()
		json.NewEncoder(f).Encode(chars)
	}
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
