package hjm3u8

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

type THjm3u8LogError struct {
	ID       int    `gorm:"primaryKey"`
	Message  string `gorm:"type:text"`
	FailedAt string
}

type THjm3u8ScannedLog struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

func IsScanned(fileName string, db *gorm.DB) (scanned bool, err error) {
	sql := `CREATE TABLE t_hjm3u8_scanned_logs (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		file_name TEXT NOT NULL, 
		scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	if !db.Migrator().HasTable("t_hjm3u8_scanned_logs") {
		db.Exec(sql)
	}
	var existing THjm3u8ScannedLog
	if err := db.Model(THjm3u8ScannedLog{}).Where("file_name = ?", fileName).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error querying database:", err)
			return false, err
		}
	}
	if existing.FileName != "" {
		return true, nil
	}
	scanned_log := THjm3u8ScannedLog{
		FileName: fileName,
	}
	if err := db.Create(&scanned_log).Error; err != nil {
		fmt.Println("Error creating record in database:", err)
		return false, err
	}
	return false, nil
}

func GzippedLogFileReader(logURL string, db *gorm.DB) error {
	db.AutoMigrate(&THjm3u8LogError{})
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
				hjm3u8Log := THjm3u8LogError{
					Message:  message,
					FailedAt: timestamp,
				}
				db.Create(&hjm3u8Log)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(logURL, "successfully scanned")
	return nil
}
