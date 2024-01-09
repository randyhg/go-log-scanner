package hjqueue

import (
	"bufio"
	"fmt"
	model "go-log-scanner/error_log_scanner/log_model"
	"net/http"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

func getURL(baseURL, logName string, currentTime time.Time) string {
	logURL := fmt.Sprintf("%s%s-%s.log", baseURL, logName, currentTime.Format("2006-01-02-15"))
	return logURL
}

func PatternedLogScanner(baseURL string, logName string, start time.Time, end time.Time, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	db.AutoMigrate(&model.QueueLogErrors{})
	for currentTime := start; currentTime.Before(end) || currentTime.Equal(end); currentTime = currentTime.Add(time.Hour) {
		logURL := getURL(baseURL, logName, currentTime)
		// fmt.Println(logURL)

		resp, err := http.Get(logURL)
		if err != nil {
			fmt.Println("Get log failed:", err)
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "|err") {
				fmt.Printf("%s: %s\n", logURL, line)
				message := model.QueueTimestamp.ReplaceAllString(line, "")
				match := model.QueueTimestamp.FindStringSubmatch(line)
				if len(match) > 0 {
					timestamp := match[0]
					hash := model.Sha256(message)
					queueLog := model.QueueLogErrors{
						Message:  message,
						FailedAt: timestamp,
						Hash:     &hash,
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
