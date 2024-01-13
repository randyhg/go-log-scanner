package response

import (
	"github.com/kataras/iris/v12"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type JsonResult struct {
	IsEncrypted bool        `json:"isEncrypted"`
	ErrorCode   int         `json:"errorCode"`
	Message     string      `json:"message"`
	Success     bool        `json:"success"`
	Data        interface{} `json:"data"`
}

func Result(code int, data interface{}, msg string, draw, recordsTotal, recordsFiltered int64, ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code":            code,
		"data":            data,
		"msg":             msg,
		"draw":            draw,
		"recordsTotal":    recordsTotal,
		"recordsFiltered": recordsFiltered,
	})
}

func ResultV2(code int, data interface{}, msg string, ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func OkWithDataV2(data interface{}, ctx iris.Context) {
	ResultV2(SUCCESS, data, "Successful operation", ctx)
}

func OkWithData(data interface{}, draw, recordsTotal, recordsFiltered int64, ctx iris.Context) {
	Result(SUCCESS, data, "Successful operation", draw, recordsTotal, recordsFiltered, ctx)
}

// func OkWithMessage(message string, ctx iris.Context) {
// 	Result(SUCCESS, map[string]interface{}{}, message, ctx)
// }

// func OkWithDetailed(data interface{}, message string, ctx iris.Context) {
// 	Result(SUCCESS, data, message, ctx)
// }

// func FailWithMessage(message string, ctx iris.Context) {
// 	Result(ERROR, map[string]interface{}{}, message, ctx)
// }
