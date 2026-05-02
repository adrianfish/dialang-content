package exporters

import (
	"log"
	"os"
	"path"
	"encoding/json"
	"github.com/adrianfish/dialang/content/db"
)

func ExportLangMappings(baseDir string) {

	if f, err := os.Create(path.Join(baseDir, "iso_lang_mappings.json")); err != nil {
		log.Fatalf("Failed to create iso_lang_mappings.json: %s\n", err)
	} else {
		defer f.Close()
		mappings := db.GetIsoLangMappings()
		json.NewEncoder(f).Encode(mappings)
	}
}
