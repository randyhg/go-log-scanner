package hjapi

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type THjApiError struct {
	ID       int    `gorm:"primaryKey"`
	Message  string `gorm:"type:text"`
	FailedAt string
}

type THjApiScannedLog struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

func IsScanned(fileName string, db *gorm.DB) (scanned bool, err error) {
	sql := `CREATE TABLE t_hj_api_scanned_logs (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		file_name TEXT NOT NULL, 
		scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	if !db.Migrator().HasTable("t_hj_api_scanned_logs") {
		db.Exec(sql)
	}
	var existing THjApiScannedLog
	if err := db.Model(THjApiScannedLog{}).Where("file_name = ?", fileName).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error querying database:", err)
			return false, err
		}
	}
	if existing.FileName != "" {
		return true, nil
	}
	scanned_log := THjApiScannedLog{
		FileName: fileName,
	}
	if err := db.Create(&scanned_log).Error; err != nil {
		fmt.Println("Error creating record in database:", err)
		return false, err
	}
	return false, nil
}

func GzippedLogFileReader(logURL string, db *gorm.DB) error {
	db.AutoMigrate(&THjApiError{})
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

	timestampRegex := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2})`)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") || strings.Contains(line, "stack") {
			message := timestampRegex.ReplaceAllString(line, "")
			match := timestampRegex.FindStringSubmatch(line)
			if len(match) > 0 {
				timestamp := match[1]
				logError := THjApiError{
					Message:  message,
					FailedAt: timestamp,
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
