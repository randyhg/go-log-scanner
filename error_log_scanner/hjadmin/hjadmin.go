package hjadmin

import (
	"bufio"
	"errors"
	"fmt"
	model "go-log-scanner/error_log_scanner/log_model"
	"go-log-scanner/util"
	"hj_common/dbmodel"
	milog "hj_common/log"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func IsScanned(fileName string, db *gorm.DB) (scanned bool, err error) {
	sql := `CREATE TABLE t_hj_admin_scanned_logs (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		file_name TEXT NOT NULL, 
		scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	if !db.Migrator().HasTable("t_hj_admin_scanned_logs") {
		db.Exec(sql)
	}
	var existing model.HjAdminScannedLogs
	if err := db.Model(model.HjAdminScannedLogs{}).Where("file_name = ?", fileName).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error querying database:", err)
			return false, err
		}
	}
	if existing.FileName != "" {
		return true, nil
	}
	scanned_log := model.HjAdminScannedLogs{
		FileName: fileName,
	}
	if err := db.Create(&scanned_log).Error; err != nil {
		fmt.Println("Error creating record in database:", err)
		return false, err
	}
	return false, nil
}

func HjAdminLogScanner(logURL string, db *gorm.DB) {
	//db.AutoMigrate(&model.HjAdminErrors{})
	tableName := dbmodel.GetMonthTableName(model.HjAdminErrors{})
	if err := util.CreateMonthTable(util.Master(), model.HjAppServerErrors{}, tableName); err != nil {
		milog.Error(err)
	}

	resp, err := http.Get(logURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var stackFlag = false
	var stackTraces []string
	var currentTrace strings.Builder
	var traceHeadMessage string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") && !strings.Contains(line, "goroutine") && !stackFlag {
			message := model.HjAdminTimestampRegex.ReplaceAllString(line, "")
			match := model.HjAdminTimestampRegex.FindStringSubmatch(line)
			if len(match) > 0 {
				timestamp := match[0]
				hash := model.HjAdminSha256(line)
				logError := model.HjAdminErrors{
					Message:  message,
					FailedAt: timestamp,
					Hash:     &hash,
				}
				//db.Create(&logError)
				db.Table(tableName).Create(&logError)
			}
		} else if strings.Contains(line, "goroutine") && !stackFlag {
			traceHeadMessage = line
			stackFlag = true
			currentTrace.WriteString(line + "\n")
		} else if stackFlag && !model.HjAdminTimestampRegex.MatchString(line) {
			currentTrace.WriteString(line + "\n")
		} else if stackFlag && model.HjAdminTimestampRegex.MatchString(line) {
			stackFlag = false
			stackTraces = append(stackTraces, currentTrace.String())
			index := len(stackTraces)
			stackTrace := stackTraces[index-1]
			currentTrace.Reset()

			message := model.HjAdminTimestampRegex.ReplaceAllString(traceHeadMessage, "")
			match := model.HjAdminTimestampRegex.FindStringSubmatch(line)
			if len(match) > 0 {
				timestamp := match[0]
				hash := model.HjAdminSha256(line)
				logError := model.HjAdminErrors{
					Message:    message,
					StackTrace: &stackTrace,
					FailedAt:   timestamp,
					Hash:       &hash,
				}
				//db.Create(&logError)
				db.Table(tableName).Create(&logError)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}
	fmt.Println(logURL, "successfully scanned")
}
