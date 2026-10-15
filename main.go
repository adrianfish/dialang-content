package main

import (
	"os"
	"log"
	"sync"
	"time"
	"github.com/adrianfish/dialang/content/exporters"
)

func main()  {

	contentDir := "../dialang-web/static/content"
	os.MkdirAll(contentDir, 0777)

	webDataDir := "./data-files"
	os.MkdirAll(webDataDir, 0777)

	start := time.Now()

	baseAssetsUrl := "https://pub-c5c4a068a24045f68b85a172153bd668.r2.dev"

	var wg sync.WaitGroup

	wg.Go(func() {
		exporters.ExportBasketPages(contentDir, baseAssetsUrl)
	})

	wg.Go(func() {
		exporters.ExportWebData(webDataDir)
		exporters.ExportLangMappings(contentDir)
		exporters.ExportALS(contentDir, baseAssetsUrl)
		exporters.ExportLaunch(contentDir, baseAssetsUrl)
		exporters.ExportHelpDialogs(contentDir, baseAssetsUrl)
		exporters.ExportSaveDialogs(contentDir, baseAssetsUrl)
		exporters.ExportLegendPages(contentDir, baseAssetsUrl)
		exporters.ExportFlowchartPages(contentDir, baseAssetsUrl)
		exporters.ExportTLSPages(contentDir, baseAssetsUrl)
		exporters.ExportDeepLinkingMenuPages(contentDir, baseAssetsUrl)
		exporters.ExportVSPTIntroPages(contentDir, baseAssetsUrl)
	})

	wg.Go(func() {
		exporters.ExportVSPTPages(contentDir, baseAssetsUrl)
		exporters.ExportVSPTFeedbackPages(contentDir)
		exporters.ExportSAIntroPages(contentDir, baseAssetsUrl)
		exporters.ExportSAPages(contentDir, baseAssetsUrl)
		exporters.ExportSAFeedbackPages(contentDir)
		exporters.ExportTestIntroPages(contentDir, baseAssetsUrl)
		exporters.ExportEndOfTestPages(contentDir)
		exporters.ExportFeedbackMenuPages(contentDir, baseAssetsUrl)
		exporters.ExportItemReviewPages(contentDir, baseAssetsUrl)
		exporters.ExportExplfbPages(contentDir, baseAssetsUrl)
		exporters.ExportAdvfbPages(contentDir)
		exporters.ExportTestResultPages(contentDir, baseAssetsUrl)
	})

	exporters.ExportQuestionnairePages(contentDir)

	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Content build took %v", elapsed)
}
