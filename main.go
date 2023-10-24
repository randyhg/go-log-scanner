package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	defaultStart := time.Date(2023, time.August, 18, 20, 0, 0, 0, time.UTC)
	defaultEnd := time.Date(2023, time.October, 24, 8, 0, 0, 0, time.UTC)

	var wg sync.WaitGroup
	wg.Add(7)

	baseURL := "https://log.hjpfef.com/hjqueue/"

	go logScanner(baseURL, "other-revenue", defaultStart, defaultEnd, &wg)        // successful
	go logScanner(baseURL, "topic-buy-stats", defaultStart, defaultEnd, &wg)      // successful
	go logScanner(baseURL, "topic-revenue", defaultStart, defaultEnd, &wg)        // successful
	go logScanner(baseURL, "update-topic-count", defaultStart, defaultEnd, &wg)   // successful
	go logScanner(baseURL, "video-add-view-count", defaultStart, defaultEnd, &wg) // successful
	go logScanner(baseURL, "video-incr", defaultStart, defaultEnd, &wg)           // successful
	go logScanner(baseURL, "video-revenue", defaultStart, defaultEnd, &wg)        // successful

	wg.Wait()

}

func getURL(baseURL, logName string, currentTime time.Time) string {
	logURL := fmt.Sprintf("%s%s-%s.log", baseURL, logName, currentTime.Format("2006-01-02-15"))
	return logURL
}

func logScanner(baseURL string, logName string, start time.Time, end time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for currentTime := start; currentTime.Before(end) || currentTime.Equal(end); currentTime = currentTime.Add(time.Hour) {
		logURL := getURL(baseURL, logName, currentTime)

		resp, err := http.Get(logURL)
		if err != nil {
			fmt.Println("Get log failed:", err)
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "|err") {
				fmt.Printf("%s: %s\n", logURL, line)
			} //else {
			// 	fmt.Println(logURL, "Passed")
			// }
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Failed to read log:", err)
			return
		}

	}
	fmt.Println(logName, "successfully scanned")
}
