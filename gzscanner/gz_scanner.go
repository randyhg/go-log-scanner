package gzscanner

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"log-scanner/chatserver"
	"log-scanner/hjapi"
	"log-scanner/hjappserver"
	"log-scanner/hjm3u8"

	"golang.org/x/net/html"
	"gorm.io/gorm"
)

var chatServerRegex = regexp.MustCompile("chatserver")
var hjAppServerRegex = regexp.MustCompile("hjappserver")
var hjm3u8Regex = regexp.MustCompile("hjm3u8")
var hjapiRegex = regexp.MustCompile("hjapi")

func ScanGzFiles(directoryURL string, db *gorm.DB) error {
	// get url
	resp, err := http.Get(directoryURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// parsing resp.body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	// mencari semua file .log.gz
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

	// validasi, scan, dan input ke db
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
