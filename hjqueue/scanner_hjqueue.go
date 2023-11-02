package hjqueue

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

type TQueueLogError struct {
	ID       int    `gorm:"primaryKey"`
	Message  string `gorm:"type:text"`
	FailedAt string
}

func getURL(baseURL, logName string, currentTime time.Time) string {
	logURL := fmt.Sprintf("%s%s-%s.log", baseURL, logName, currentTime.Format("2006-01-02-15"))
	return logURL
}

func PatternedLogScanner(baseURL string, logName string, start time.Time, end time.Time, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	db.AutoMigrate(&TQueueLogError{})
	for currentTime := start; currentTime.Before(end) || currentTime.Equal(end); currentTime = currentTime.Add(time.Hour) {
		logURL := getURL(baseURL, logName, currentTime)
		// fmt.Println(logURL)

		resp, err := http.Get(logURL)
		if err != nil {
			fmt.Println("Get log failed:", err)
			return
		}
		defer resp.Body.Close()

		timestamp := regexp.MustCompile(`\b\w{3} \d{2} \d{2}:\d{2}:\d{2}\b`)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "|err") {
				fmt.Printf("%s: %s\n", logURL, line)
				message := timestamp.ReplaceAllString(line, "")
				match := timestamp.FindStringSubmatch(line)
				if len(match) > 0 {
					timestamp := match[1]
					queueLog := TQueueLogError{
						Message:  message,
						FailedAt: timestamp,
					}
					db.Create(&queueLog)
					fmt.Println("Successfully input queue log")
				}
			} //else {
			// 	fmt.Println(line)
			// }
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Failed to read log:", err)
			return
		}

	}
	fmt.Println(logName, "successfully scanned")
}
