package main

import (
	"os"
	"log"
	"sync"
	"time"
	"github.com/dialangproject/content/exporters"
)

func main()  {

	contentDir := "./static-site/content"
	os.MkdirAll(contentDir, 0777)

	webDataDir := "../web/data-files"
	os.MkdirAll(webDataDir, 0777)

	start := time.Now()

	var wg sync.WaitGroup

	/*
	wg.Go(func() {
		exporters.ExportBasketPages(contentDir)
	})
	*/

	wg.Go(func() {
		//exporters.ExportKeyboardFragments(contentDir)
		exporters.ExportWebData(webDataDir)
		/*
		exporters.ExportALS(contentDir)
		exporters.ExportHelpDialogs(contentDir)
		exporters.ExportLegendPages(contentDir)
		exporters.ExportFlowchartPages(contentDir)
		exporters.ExportTLSPages(contentDir)
		exporters.ExportVSPTIntroPages(contentDir)
		*/
	})

	/*
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
	*/

	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Content build took %v", elapsed)
}
