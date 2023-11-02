package main

import (
	"fmt"
	"log-scanner/gzscanner"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/production_error_log?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// hjapi scanner
	// var wg sync.WaitGroup
	// wg.Add(13)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-11/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-22/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-33/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-44/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-51/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-52/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-53/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-71/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-72/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-73/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-74/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-75/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-76/", db, &wg)
	// wg.Wait()

	// hjappserver scanner
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go multipleUrlScanner("https://log.hjpfef.com/hjappserver/hj-appserver-1/", db, &wg)
	// go multipleUrlScanner("https://log.hjpfef.com/hjappserver/hj-appserver-2/", db, &wg)
	// wg.Wait()

	// hjqueue scanner
	// var wg sync.WaitGroup
	// wg.Add(8)
	// defaultStart := time.Date(2023, time.October, 27, 10, 0, 0, 0, time.UTC)
	// defaultEnd := time.Date(2023, time.October, 31, 23, 0, 0, 0, time.UTC)
	// baseURL := "https://log.hjpfef.com/hjqueue/"
	// go hjqueue.PatternedLogScanner(baseURL, "other-revenue", defaultStart, defaultEnd, db, &wg)        // successful
	// go hjqueue.PatternedLogScanner(baseURL, "topic-buy-stats", defaultStart, defaultEnd, db, &wg)      // successful
	// go hjqueue.PatternedLogScanner(baseURL, "topic-revenue", defaultStart, defaultEnd, db, &wg)        // successful
	// go hjqueue.PatternedLogScanner(baseURL, "update-topic-count", defaultStart, defaultEnd, db, &wg)   // successful
	// go hjqueue.PatternedLogScanner(baseURL, "video-add-view-count", defaultStart, defaultEnd, db, &wg) // successful
	// go hjqueue.PatternedLogScanner(baseURL, "video-incr", defaultStart, defaultEnd, db, &wg)           // successful
	// go hjqueue.PatternedLogScanner(baseURL, "video-revenue", defaultStart, defaultEnd, db, &wg)        // successful
	// wg.Wait()

	// chatserver scanner
	// directoryURL := "https://log.hjpfef.com/chatserver/"
	// err = gzscanner.ScanGzFiles(directoryURL, db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// hjm3u8 scanner
	directoryURL := "https://log.hjpfef.com/hjm3u8/m3u801/"
	err = gzscanner.ScanGzFiles(directoryURL, db)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func multipleUrlScanner(directoryURL string, db *gorm.DB, wg *sync.WaitGroup) error {
	defer wg.Done()
	if err := gzscanner.ScanGzFiles(directoryURL, db); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
