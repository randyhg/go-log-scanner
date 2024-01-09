package hjqueue

import (
	"bufio"
	"errors"
	"fmt"
	model "go-log-scanner/error_log_scanner/log_model"
	"go-log-scanner/util"
	"hj_common/log"
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

		scanned, err := isScanned(logURL, util.Master())
		if err != nil {
			log.Error("Error while check isScanned", err)
			return
		}

		if scanned {
			continue
		}

		resp, err := http.Get(logURL)
		if err != nil {
			log.Error("Get log failed:", err)
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

func isScanned(fileName string, db *gorm.DB) (scanned bool, err error) {
	sql := `CREATE TABLE t_queue_scanned_logs (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		file_name TEXT NOT NULL, 
		scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if !db.Migrator().HasTable("t_queue_scanned_logs") {
		db.Exec(sql)
	}

	var existing model.QueueScannedLogs
	if err := db.Model(model.QueueScannedLogs{}).Where("file_name = ?", fileName).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error querying database:", err)
			return false, err
		}
	}
	if existing.FileName != "" {
		return true, nil
	}
	scanned_log := model.QueueScannedLogs{
		FileName: fileName,
	}
	if err := db.Create(&scanned_log).Error; err != nil {
		fmt.Println("Error creating record in database:", err)
		return false, err
	}
	return false, nil
}
