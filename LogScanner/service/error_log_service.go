package service

import (
	"go-log-scanner/util"
	milog "hj_common/log"
)

var ErrorLogService = newErrorLogService()

func newErrorLogService() *errorLogService {
	return &errorLogService{}
}

type errorLogService struct {
}

type ChatServerLogErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

func (c *errorLogService) GetErrorTotalService(hash string, model interface{}) (int64, error) {
	var total int64
	// debug.PrintStack()
	if err := util.Master().Model(model).Where("hash = ?", hash).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (c *errorLogService) GetErrorMessageService(model interface{}) (string, error) {
	return "", nil
}

type Results struct {
	Message  string `json:"message"`
	FailedAt string `json:"failed_at"`
	Hash     string `json:"hash"`
	Total    int64  `json:"total"`
	FileName string `json:"file_name"`
}

func (c *errorLogService) GetAll(model interface{}) []Results {
	var results []Results
	util.Master().Model(model).Select("MIN(message) as message, MAX(failed_at) AS failed_at, hash as hash, COUNT(*) as total").Group("hash").Scan(&results)
	return results
}

type ErrorResults struct {
	Id         int64  `json:"id"`
	Message    string `gorm:"message" json:"message"`
	FailedAt   string `gorm:"failed_at" json:"failed_at"`
	StackTrace string `gorm:"stack_trace" json:"stack_trace"`
	Hash       string `gorm:"hash" json:"hash"`
}

func (c *errorLogService) GetAllErrors(model interface{}, hash string, offset int) []ErrorResults {
	var results []ErrorResults
	if err := util.Master().Model(model).Where("hash = ?", hash).Order("id DESC").Scan(&results).Error; err != nil {
		// if err := util.Master().Model(model).Where("hash = ?", hash).Order("id DESC").Limit(20).Offset(offset).Scan(&results).Error; err != nil {
		milog.Error(err)
		return nil
	}
	return results
}

func (c *errorLogService) GetAllErrorsV2(model interface{}, hash string, start int, length int) []ErrorResults {
	var results []ErrorResults
	if err := util.Master().Model(model).Where("hash = ?", hash).Order("id DESC").Limit(length).Offset(start).Scan(&results).Error; err != nil {
		milog.Error(err)
		return nil
	}
	return results
}

func (c *errorLogService) GetAllErrorsTotalByHash(model interface{}, hash string) (total int64) {
	if err := util.Master().Model(model).Where("hash = ?", hash).Count(&total).Error; err != nil {
		milog.Error(err)
		return 0
	}
	return total
}
