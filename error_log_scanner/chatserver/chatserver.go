package chatserver

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	model "go-log-scanner/error_log_scanner/log_model"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func IsScanned(fileName string, db *gorm.DB) (scanned bool, err error) {
	sql := `CREATE TABLE t_chat_server_scanned_logs (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		file_name TEXT NOT NULL, 
		scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	if !db.Migrator().HasTable("t_chat_server_scanned_logs") {
		db.Exec(sql)
	}
	var existing model.ChatServerScannedLogs
	if err := db.Model(model.ChatServerScannedLogs{}).Where("file_name = ?", fileName).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error querying database:", err)
			return false, err
		}
	}
	if existing.FileName != "" {
		return true, nil
	}
	scanned_log := model.ChatServerScannedLogs{
		FileName: fileName,
	}
	if err := db.Create(&scanned_log).Error; err != nil {
		fmt.Println("Error creating record in database:", err)
		return false, err
	}
	return false, nil
}

func GzippedLogFileReader(logURL string, db *gorm.DB) error {
	db.AutoMigrate(&model.ChatServerLogErrors{})
	resp, err := http.Get(logURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	var stackFlag = false
	var stackTraces []string
	var currentTrace strings.Builder
	var traceHeadMessage string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") && !strings.Contains(line, "goroutine") && !stackFlag {
			message := model.TimestampRegex.ReplaceAllString(line, "")
			match := model.TimestampRegex.FindStringSubmatch(line)
			if len(match) > 0 {
				timestamp := match[0]
				hash := model.Sha256(message)
				logError := model.ChatServerLogErrors{
					Message:  message,
					FailedAt: timestamp,
					Hash:     &hash,
				}
				db.Create(&logError)
			}
		} else if strings.Contains(line, "goroutine") && !stackFlag {
			traceHeadMessage = line
			stackFlag = true
			currentTrace.WriteString(line + "\n")
		} else if stackFlag && !model.TimestampRegex.MatchString(line) {
			currentTrace.WriteString(line + "\n")
		} else if stackFlag && model.TimestampRegex.MatchString(line) {
			stackFlag = false
			stackTraces = append(stackTraces, currentTrace.String())
			index := len(stackTraces)
			stackTrace := stackTraces[index-1]
			currentTrace.Reset()

			message := model.TimestampRegex.ReplaceAllString(traceHeadMessage, "")
			match := model.TimestampRegex.FindStringSubmatch(line)
			if len(match) > 0 {
				timestamp := match[0]
				hash := model.Sha256(message)
				logError := model.ChatServerLogErrors{
					Message:    message,
					StackTrace: &stackTrace,
					FailedAt:   timestamp,
					Hash:       &hash,
				}
				db.Create(&logError)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(logURL, "successfully scanned")
	return nil
}
