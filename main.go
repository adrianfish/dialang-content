package main

import (
	"path/filepath"
	"io"
	"os"
	"log"
	"sync"
	"time"
	"github.com/adrianfish/dialang/content/exporters"
)

func main()  {

	contentDir := "../dialang-web/static/content"
	os.MkdirAll(contentDir, 0777)

	webDataDir := "../dialang-web/data-files"
	os.MkdirAll(webDataDir, 0777)

	namesFile, err := os.Create(filepath.Join(webDataDir, "language-names.json"));
	if err != nil {
	}

	srcNamesFile, err := os.Open("./data-files/language-names.json")
	_, erre := io.Copy(namesFile, srcNamesFile)

	if erre != nil {
		log.Println("Failed to copy language names");
	}

	start := time.Now()

	var wg sync.WaitGroup

	wg.Go(func() {
		exporters.ExportBasketPages(contentDir)
	})

	wg.Go(func() {
		exporters.ExportWebData(webDataDir)
		exporters.ExportLangMappings(contentDir)
		exporters.ExportALS(contentDir)
		exporters.ExportLaunch(contentDir)
		exporters.ExportHelpDialogs(contentDir)
		exporters.ExportLegendPages(contentDir)
		exporters.ExportFlowchartPages(contentDir)
		exporters.ExportTLSPages(contentDir)
		exporters.ExportDeepLinkingMenuPages(contentDir)
		exporters.ExportVSPTIntroPages(contentDir)
	})

	wg.Go(func() {
		exporters.ExportVSPTPages(contentDir)
		exporters.ExportVSPTFeedbackPages(contentDir)
		exporters.ExportSAIntroPages(contentDir)
		exporters.ExportSAPages(contentDir)
		exporters.ExportSAFeedbackPages(contentDir)
		exporters.ExportTestIntroPages(contentDir)
		exporters.ExportEndOfTestPages(contentDir)
		exporters.ExportFeedbackMenuPages(contentDir)
		exporters.ExportItemReviewPages(contentDir)
		exporters.ExportExplfbPages(contentDir)
		exporters.ExportAdvfbPages(contentDir)
		exporters.ExportTestResultPages(contentDir)
	})

	exporters.ExportQuestionnairePages(contentDir)

	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Content build took %v", elapsed)
}
