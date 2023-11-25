package route

import (
	"go-log-scanner/LogScanner/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func RegisterRoutes(app *iris.Application) {
	tmpl := iris.HTML("./LogScanner/views", ".html")
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index")
	})

	apiHjm3u8(app.Party("/hjm3u8"))
	apiChatServer(app.Party("/chat_server"))
	apiHjAppServer(app.Party("/hjapp_server"))
	apiHjApi(app.Party("/hjapi"))
	apiHjQueue(app.Party("/hjqueue"))
	apiHjAdmin(app.Party("/hjadmin"))

}

func apiHjm3u8(router router.Party) {
	data := iris.Map{
		"Title": "HJ M3u8 Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjm3u8", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjm3u8_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjm3u8ErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjm3u8AllErrorByHash)
}

func apiChatServer(router router.Party) {
	data := iris.Map{
		"Title": "Chat Server Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("chat_server", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("chat_server_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetChatServerErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetChatServerAllErrorByHash)
}

func apiHjAppServer(router router.Party) {
	data := iris.Map{
		"Title": "HJ App Server Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjapp_server", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjapp_server_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjAppServerErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjAppServerAllErrorByHash)
}

func apiHjApi(router router.Party) {
	data := iris.Map{
		"Title": "HJ Api Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjapi", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjapi_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjApiErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjApiAllErrorByHash)
}

func apiHjQueue(router router.Party) {
	data := iris.Map{
		"Title": "HJ Queue Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjqueue", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjqueue_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjQueueErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjQueueAllErrorByHash)
}

func apiHjAdmin(router router.Party) {
	data := iris.Map{
		"Title": "HJ Admin Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjadmin", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjadmin_errors", data)
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjAdminErrorList)
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjAdminAllErrorByHash)
}
