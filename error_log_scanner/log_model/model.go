package log_model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
)

var TimestampRegex = regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2})`)
var HjAdminTimestampRegex = regexp.MustCompile(`\[\w+-ADMIN\](\d{4}/\d{2}/\d{2} - \d{2}:\d{2}:\d{2}.\d{3})`)
var QueueTimestamp = regexp.MustCompile(`\b\w{3} \d{2} \d{2}:\d{2}:\d{2}\b`)
var fileNameRegex = regexp.MustCompile(`(\w+\.go):(\d+)`)
var hjAdminRegex = regexp.MustCompile(`error\s+(\p{Han}+.*?)(.{1})`)

type ChatServerLogErrors struct {
	ID         int    `gorm:"primaryKey"`
	Message    string `gorm:"type:text"`
	StackTrace *string
	FailedAt   string
	Hash       *string
}

type ChatServerScannedLogs struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

type HjApiErrors struct {
	ID         int    `gorm:"primaryKey"`
	Message    string `gorm:"type:text"`
	StackTrace *string
	FailedAt   string
	Hash       *string
}

type HjApiScannedLogs struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

type HjAppServerErrors struct {
	ID         int    `gorm:"primaryKey"`
	Message    string `gorm:"type:text"`
	StackTrace *string
	FailedAt   string  `gorm:"type:text"`
	Hash       *string `gorm:"type:text"`
}

type HjAppServerScannedLogs struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

type Hjm3u8LogErrors struct {
	ID         int    `gorm:"primaryKey"`
	Message    string `gorm:"type:text"`
	StackTrace *string
	FailedAt   string
	Hash       *string
}

type Hjm3u8ScannedLogs struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

type HjAdminErrors struct {
	ID         int    `gorm:"primaryKey"`
	Message    string `gorm:"type:text"`
	StackTrace *string
	FailedAt   string  `gorm:"type:text"`
	Hash       *string `gorm:"type:text"`
}

type HjAdminScannedLogs struct {
	ID       int    `gorm:"primaryKey"`
	FileName string `gorm:"type:text"`
}

type QueueLogErrors struct {
	ID       int    `gorm:"primaryKey"`
	Message  string `gorm:"type:text"`
	FailedAt string
	Hash     *string
}

func Sha256(data string) string {
	match := fileNameRegex.FindStringSubmatch(data)
	if len(match) >= 3 {
		fileName := match[1]
		lineNumber := match[2]
		hashFormat := fmt.Sprintf("%s:%s", fileName, lineNumber)
		hash := sha256.Sum256([]byte(hashFormat))
		hashString := hex.EncodeToString(hash[:])
		return hashString
	}
	return ""
}

func HjAdminSha256(data string) string {
	match := hjAdminRegex.FindStringSubmatch(data)
	if len(match) >= 3 {
		hashName := match[0]
		hash := sha256.Sum256([]byte(hashName))
		hashString := hex.EncodeToString(hash[:])
		return hashString
	}
	return ""
}
