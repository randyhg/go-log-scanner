package cmd

import (
	"fmt"
	"go-log-scanner/config"
	"go-log-scanner/error_log_scanner/chatserver"
	"go-log-scanner/error_log_scanner/hjapi"
	"go-log-scanner/error_log_scanner/hjappserver"
	"go-log-scanner/error_log_scanner/hjm3u8"
	"go-log-scanner/error_log_scanner/hjqueue"
	"go-log-scanner/util"
	milog "hj_common/log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"gorm.io/gorm"
)

var logScanner = &cobra.Command{
	Use:   "log-scanner",
	Short: "Error log scanner",
	Run:   errorLogScanner,
}

var chatServerRegex = regexp.MustCompile("chatserver")
var hjAppServerRegex = regexp.MustCompile("hjappserver")
var hjm3u8Regex = regexp.MustCompile("hjm3u8")
var hjapiRegex = regexp.MustCompile("hjapi")

func init() {
	rootCmd.AddCommand(logScanner)
}

func errorLogScanner(cmd *cobra.Command, args []string) {
	config.Init()
	util.Init()
	db := util.Master()

	var wg sync.WaitGroup
	wg.Add(25)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-10/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-11/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-22/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-33/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-44/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-51/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-52/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-53/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-71/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-72/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-73/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-74/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-75/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjapi/hj-api-log-76/", db, &wg)

	// hjm3u8 scanner
	go multipleUrlScanner("https://log.hjpfef.com/hjm3u8/m3u801/", db, &wg)

	// hjappserver scanner
	go multipleUrlScanner("https://log.hjpfef.com/hjappserver/hj-appserver-1/", db, &wg)
	go multipleUrlScanner("https://log.hjpfef.com/hjappserver/hj-appserver-2/", db, &wg)

	// chatserver scanner
	go multipleUrlScanner("https://log.hjpfef.com/chatserver/", db, &wg)

	// =============================================================================

	// hjqueue scanner
	defaultStart := time.Date(2023, time.November, 22, 0, 0, 0, 0, time.UTC)
	defaultEnd := time.Date(2023, time.November, 22, 23, 0, 0, 0, time.UTC)
	baseURL := "https://log.hjpfef.com/hjqueue/"
	go hjqueue.PatternedLogScanner(baseURL, "other-revenue", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "topic-buy-stats", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "topic-revenue", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "update-topic-count", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "video-add-view-count", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "video-incr", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(baseURL, "video-revenue", defaultStart, defaultEnd, db, &wg)

	wg.Wait()
}

func multipleUrlScanner(directoryURL string, db *gorm.DB, wg *sync.WaitGroup) error {
	defer wg.Done()
	if err := scanGzFiles(directoryURL, db); err != nil {
		milog.Error(err)
		return err
	}
	return nil
}

func scanGzFiles(directoryURL string, db *gorm.DB) error {
	resp, err := http.Get(directoryURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var fileList []string
	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.HasSuffix(attr.Val, ".log.gz") {
					fileList = append(fileList, directoryURL+attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}
	extractLinks(doc)

	for _, fileName := range fileList {
		if chatServerRegex.MatchString(directoryURL) {
			scanned, err := chatserver.IsScanned(fileName, db)
			if err != nil {
				fmt.Println("Error processing file:", err)
				return err
			}
			if scanned {
				continue
			} else {
				if err := chatserver.GzippedLogFileReader(fileName, db); err != nil {
					fmt.Println("Error creating record in database:", err)
					return err
				}
			}
		} else if hjAppServerRegex.MatchString(directoryURL) {
			scanned, err := hjappserver.IsScanned(fileName, db)
			if err != nil {
				fmt.Println("Error processing file:", err)
				return err
			}
			if scanned {
				continue
			} else {
				if err := hjappserver.GzippedLogFileReader(fileName, db); err != nil {
					fmt.Println("Error creating record in database:", err)
					return err
				}
			}
		} else if hjm3u8Regex.MatchString(directoryURL) {
			scanned, err := hjm3u8.IsScanned(fileName, db)
			if err != nil {
				fmt.Println("Error processing file:", err)
				return err
			}
			if scanned {
				continue
			} else {
				if err := hjm3u8.GzippedLogFileReader(fileName, db); err != nil {
					fmt.Println("Error creating record in database:", err)
					return err
				}
			}
		} else if hjapiRegex.MatchString(directoryURL) {
			scanned, err := hjapi.IsScanned(fileName, db)
			if err != nil {
				fmt.Println("Error processing file:", err)
				return err
			}
			if scanned {
				continue
			} else {
				if err := hjapi.GzippedLogFileReader(fileName, db); err != nil {
					fmt.Println("Error creating record in database:", err)
					return err
				}
			}
		} else {
			fmt.Printf("Handler for %s not found\n", directoryURL)
			return err
		}
	}
	return nil
}
