package main

import (
	"os"
	"fmt"
	"log"
	"strings"
)

var iso_to_bcp = map[string]string {
	"nld_nl": "nl-NL",
	"spa_es": "es-ES",
	"fin_fi": "fi-FI",
	"ell_gr": "el-GR",
	"eng_gb": "en-GB",
	"por_pt": "pt-PT",
	"nor_no": "no-NO",
	"dan_dk": "da-DK",
	"ita_it": "it-IT",
	"gle_ie": "ga-IE",
	"deu_de": "de-DE",
	"cmn_cn": "zh-CN",
	"isl_is": "is-IS",
	"fra_fr": "fr-FR",
	"swe_se": "sv-SE",
	"ind_id": "id-ID",
}

func main()  {

	files, err := os.ReadDir("./admin-texts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := "./admin-texts/" + file.Name()
		input, err := os.ReadFile(name)
		if err != nil {
		}

		replaced := replaceAll(string(input))
		fmt.Println(replaced)
		os.WriteFile(name, []byte(replaced), 0666)
	}
}

func replaceAll(input string) string {

	for k, v := range iso_to_bcp {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}
