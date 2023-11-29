package cmd

import (
	"bufio"
	"fmt"
	"go-log-scanner/config"
	"go-log-scanner/error_log_scanner/chatserver"
	"go-log-scanner/error_log_scanner/hjadmin"
	"go-log-scanner/error_log_scanner/hjapi"
	"go-log-scanner/error_log_scanner/hjappserver"
	"go-log-scanner/error_log_scanner/hjm3u8"
	"go-log-scanner/error_log_scanner/hjqueue"
	"go-log-scanner/util"
	milog "hj_common/log"
	"log"
	"net/http"
	"os"
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

	if len(args) == 0 {
		scanAllLogs(util.Master())
	} else if len(args) == 1 {
		logUrl := args[0]
		singleLogScanner(logUrl, util.Master())
		fmt.Println("Log printed to ./scanned_logs/")
	} else if len(args) == 2 { // go run . log-scanner "url" gz
		logUrl := args[1]
		if err := singleGzLogScanner(logUrl, util.Master()); err != nil {
			milog.Error(err)
		}
	}
}

func scanAllLogs(db *gorm.DB) {
	baseUrl := config.Instance.BaseURL
	var wg sync.WaitGroup
	wg.Add(29)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-10/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-11/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-22/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-33/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-44/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-51/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-52/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-53/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-71/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-72/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-73/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-74/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-75/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjapi/hj-api-log-76/", baseUrl), db, &wg)

	// hjm3u8 scanner
	go multipleUrlScanner(fmt.Sprintf("%s/hjm3u8/m3u801/", baseUrl), db, &wg)

	// hjappserver scanner
	go multipleUrlScanner(fmt.Sprintf("%s/hjappserver/hj-appserver-1/", baseUrl), db, &wg)
	go multipleUrlScanner(fmt.Sprintf("%s/hjappserver/hj-appserver-2/", baseUrl), db, &wg)

	// chatserver scanner
	go multipleUrlScanner(fmt.Sprintf("%s/chatserver/", baseUrl), db, &wg)

	// =============================================================================

	// hjqueue scanner
	yesterday := time.Now().AddDate(0, 0, -1)
	defaultStart := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.UTC)
	defaultEnd := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 0, 0, 0, time.UTC)
	hjqueueUrl := fmt.Sprintf("%s/hjqueue/", baseUrl)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "other-revenue", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "topic-buy-stats", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "topic-revenue", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "update-topic-count", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "video-add-view-count", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "video-incr", defaultStart, defaultEnd, db, &wg)
	go hjqueue.PatternedLogScanner(hjqueueUrl, "video-revenue", defaultStart, defaultEnd, db, &wg)

	// =============================================================================

	// hjadmin scanner
	go hjAdminValidation(fmt.Sprintf("%s/hjadmin/2023-11-03.log", baseUrl), db, &wg)
	go hjAdminValidation(fmt.Sprintf("%s/hjadmin/2023-11-17.log", baseUrl), db, &wg)
	go hjAdminValidation(fmt.Sprintf("%s/hjadmin/2023-11-20.log", baseUrl), db, &wg)
	go hjAdminValidation(fmt.Sprintf("%s/hjadmin/2023-11-23.log", baseUrl), db, &wg)

	wg.Wait()
}

func multipleUrlScanner(directoryURL string, db *gorm.DB, wg *sync.WaitGroup) error {
	defer wg.Done()
	if err := scanGzFiles(directoryURL, db); err != nil {
		milog.Errorf("Scan %s error: %s", directoryURL, err)
		return err
	}
	return nil
}

func singleLogScanner(logUrl string, db *gorm.DB) {
	resp, err := http.Get(logUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("scanned_logs/scanned.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") {
			log.Print(line)
		}
	}
}

func singleGzLogScanner(logUrl string, db *gorm.DB) error {
	if chatServerRegex.MatchString(logUrl) {
		scanned, err := chatserver.IsScanned(logUrl, db)
		if err != nil {
			fmt.Println("Error processing file:", err)
			return err
		}
		if scanned {
			err := fmt.Errorf("%s has been scanned", logUrl)
			return err
		} else {
			if err := chatserver.GzippedLogFileReader(logUrl, db); err != nil {
				fmt.Println("Error creating record in database:", err)
				return err
			}
		}
	} else if hjAppServerRegex.MatchString(logUrl) {
		scanned, err := hjappserver.IsScanned(logUrl, db)
		if err != nil {
			fmt.Println("Error processing file:", err)
			return err
		}
		if scanned {
			err := fmt.Errorf("%s has been scanned", logUrl)
			return err
		} else {
			if err := hjappserver.GzippedLogFileReader(logUrl, db); err != nil {
				fmt.Println("Error creating record in database:", err)
				return err
			}
		}
	} else if hjm3u8Regex.MatchString(logUrl) {
		scanned, err := hjm3u8.IsScanned(logUrl, db)
		if err != nil {
			fmt.Println("Error processing file:", err)
			return err
		}
		if scanned {
			err := fmt.Errorf("%s has been scanned", logUrl)
			return err
		} else {
			if err := hjm3u8.GzippedLogFileReader(logUrl, db); err != nil {
				fmt.Println("Error creating record in database:", err)
				return err
			}
		}
	} else if hjapiRegex.MatchString(logUrl) {
		scanned, err := hjapi.IsScanned(logUrl, db)
		if err != nil {
			fmt.Println("Error processing file:", err)
			return err
		}
		if scanned {
			err := fmt.Errorf("%s has been scanned", logUrl)
			return err
		} else {
			if err := hjapi.GzippedLogFileReader(logUrl, db); err != nil {
				fmt.Println("Error creating record in database:", err)
				return err
			}
		}
	}
	return nil
}

func hjAdminValidation(directoryURL string, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	scanned, err := hjadmin.IsScanned(directoryURL, db)
	if err != nil {
		milog.Error(err)
		return
	}
	if !scanned {
		hjadmin.HjAdminLogScanner(directoryURL, db)
	}
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
