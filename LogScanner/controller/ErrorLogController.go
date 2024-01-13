package controller

import (
	"fmt"
	"go-log-scanner/LogScanner/model"
	"go-log-scanner/LogScanner/response"
	"go-log-scanner/LogScanner/service"
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

func (c *errorLogController) GetHjm3u8ErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.Hjm3u8LogErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetHjm3u8AllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.Hjm3u8LogErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.Hjm3u8LogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetChatServerErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.ChatServerLogErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetChatServerAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.ChatServerLogErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.ChatServerLogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAppServerErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.HjAppServerErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetHjAppServerAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.HjAppServerErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjAppServerErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjApiErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.HjApiErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetHjApiAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.HjApiErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjApiErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAdminErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.HjAdminErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetHjAdminAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.HjAdminErrors{}, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjAdminErrors{}, errorHash)
	// response.OkWithData(results, ctx)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjQueueErrorList(ctx iris.Context) response.JsonResult {
	results := service.ErrorLogService.GetAll(model.QueueLogErrors{})

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

	return response.JsonResult{
		Data: results,
	}
}

func (c *errorLogController) GetHjQueueAllErrorByHash(ctx iris.Context) {
	draw, _ := strconv.Atoi(ctx.URLParam("draw"))
	start, _ := strconv.Atoi(ctx.URLParam("start"))
	length, _ := strconv.Atoi(ctx.URLParam("length"))
	errorHash := ctx.Params().GetString("errorHash")
	results := service.ErrorLogService.GetAllErrorsV2(model.QueueLogErrors{}, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.QueueLogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetChatServerTotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.ChatServerLogErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}

func (c *errorLogController) GetHjAdminTotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.HjAdminErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}

func (c *errorLogController) GetHjApiTotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.HjApiErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}

func (c *errorLogController) GetHjAppServerTotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.HjAppServerErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}

func (c *errorLogController) GetHjM3u8TotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.Hjm3u8LogErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}

func (c *errorLogController) GetHjQueueTotalRecord(ctx iris.Context) response.JsonResult {
	total := service.ErrorLogService.GetTotalRecordService(model.QueueLogErrors{})
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(total, number.Scale(scale))
	p := message.NewPrinter(language.English)
	formattedTotal := p.Sprintf("%v", dec)
	return response.JsonResult{
		Data: formattedTotal,
	}
}
