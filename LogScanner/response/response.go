package response

import (
	"github.com/kataras/iris/v12"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

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

// func FailWithMessage(message string, ctx iris.Context) {
// 	Result(ERROR, map[string]interface{}{}, message, ctx)
// }

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

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	ScrollId string      `json:"scrollId"`
}
