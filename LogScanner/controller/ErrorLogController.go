package controller

import (
	"fmt"
	//"go-log-scanner/LogScanner/model"
	"go-log-scanner/LogScanner/response"
	"go-log-scanner/LogScanner/service"
	model "go-log-scanner/error_log_scanner/log_model"
	"hj_common/dbmodel"
	"regexp"
	"strconv"

	"github.com/kataras/iris/v12"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type LogErrorResponse struct {
	Message string `json:"message"`
	Total   int64  `json:"total"`
}

var ErrorLogController = newErrorLogController()
var fileNameRegex = regexp.MustCompile(`(\w+\.go):(\d+)`)

func newErrorLogController() *errorLogController {
	return &errorLogController{}
}

type errorLogController struct {
}

func (c *errorLogController) GetHjm3u8ErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.Hjm3u8LogErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetHjm3u8AllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.Hjm3u8LogErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetChatServerErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.ChatServerLogErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetChatServerAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.ChatServerLogErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAppServerErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.HjAppServerErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetHjAppServerAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.HjAppServerErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjApiErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.HjApiErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetHjApiAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.HjApiErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAdminErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.HjAdminErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetHjAdminAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.HjAdminErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	// response.OkWithData(results, ctx)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjQueueErrorList(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.QueueLogErrors{})
	results := service.ErrorLogService.GetAll(tableName)

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
		limitedMessage := service.ErrorLogService.LimitString(result.Message, 100)
		results[i].Message = limitedMessage
		if len(match) >= 3 {
			fileName := match[1]
			lineNumber := match[2]
			results[i].FileName = fmt.Sprintf("%s:%s", fileName, lineNumber)
		}
	}

	response.OkWithDataV2(results, ctx)
}

func (c *errorLogController) GetHjQueueAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	tableName := dbmodel.GetMonthTableName(model.QueueLogErrors{})
	results := service.ErrorLogService.GetAllErrorsV2(tableName, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(tableName, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetChatServerTotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.ChatServerLogErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}

func (c *errorLogController) GetHjAdminTotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.ChatServerLogErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}

func (c *errorLogController) GetHjApiTotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.HjApiErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}

func (c *errorLogController) GetHjAppServerTotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.HjAppServerErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}

func (c *errorLogController) GetHjM3u8TotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.Hjm3u8LogErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}

func (c *errorLogController) GetHjQueueTotalRecord(ctx iris.Context) {
	tableName := dbmodel.GetMonthTableName(model.QueueLogErrors{})
	total := service.ErrorLogService.GetTotalRecordService(tableName)
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	response.OkWithDataV2(formattedTotal, ctx)
}
