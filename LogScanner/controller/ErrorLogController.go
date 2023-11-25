package controller

import (
	"fmt"
	"go-log-scanner/LogScanner/model"
	"go-log-scanner/LogScanner/response"
	"go-log-scanner/LogScanner/service"
	"regexp"
	"strconv"

	"github.com/kataras/iris/v12"
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

type Pagination struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
}

// func Index(ctx iris.Context) {
// 	ctx.View("hjm3u8")
// }

func Hjm3u8Errors(ctx iris.Context) {
	page := ctx.Params().GetIntDefault("page", 1)
	data := map[string]interface{}{
		"pagination": Pagination{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
		},
		"path": ctx.Path(),
	}
	ctx.View("hjm3u8_errors", data)
}

func (c *errorLogController) GetHjm3u8ErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.Hjm3u8LogErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.Hjm3u8LogErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.Hjm3u8LogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetChatServerErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.ChatServerLogErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.ChatServerLogErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.ChatServerLogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAppServerErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.HjAppServerErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.HjAppServerErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjAppServerErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjApiErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.HjApiErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.HjApiErrors{}, errorHash, start, length)

	// response.OkWithData(results, ctx)
	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjApiErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjAdminErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.HjAdminErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.HjAdminErrors{}, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.HjAdminErrors{}, errorHash)
	// response.OkWithData(results, ctx)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

func (c *errorLogController) GetHjQueueErrorList(ctx iris.Context) {
	results := service.ErrorLogService.GetAll(model.QueueLogErrors{})

	for i, result := range results {
		match := fileNameRegex.FindStringSubmatch(result.Message)
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
	results := service.ErrorLogService.GetAllErrorsV2(model.QueueLogErrors{}, errorHash, start, length)

	total := service.ErrorLogService.GetAllErrorsTotalByHash(model.QueueLogErrors{}, errorHash)
	response.OkWithData(results, int64(draw), total, total, ctx)
}

// func (c *errorLogController) GetPagination(ctx iris.Context) {
// 	page := ctx.Params().GetIntDefault("page", 1)
// 	offset := (page - 1) * 10
// 	errorHash := ctx.Params().GetString("errorHash")

// 	results := service.ErrorLogService.GetAllErrors(model.Hjm3u8LogErrors{}, errorHash, offset)
// 	fmt.Println(len(results))

// 	response.OkWithData(results, ctx)
// }

// // daksjdlakdl
// func (c *errorLogController) GetErrorMessage(ctx iris.Context) {
// 	hash := "1c27099b3b84b13d0e3fbd299ba93ae7853ec1d0d3a4e5daa89e68b7ad59d7cb"
// 	message, err := service.ErrorLogService.GetErrorMessageService(hash)
// 	if err != nil {
// 		return
// 	}
// 	log.Info(message)
// 	ctx.JSON(LogErrorResponse{
// 		Message: message,
// 		Total:   0,
// 	})
// }
