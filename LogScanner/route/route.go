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
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjm3u8")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjm3u8_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjm3u8ErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetHjm3u8AllErrorByHash)
}

func apiChatServer(router router.Party) {
	router.Get("/", func(ctx iris.Context) {
		ctx.View("chat_server")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("chat_server_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetChatServerErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetChatServerAllErrorByHash)
}

func apiHjAppServer(router router.Party) {
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjapp_server")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjapp_server_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjAppServerErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetHjAppServerAllErrorByHash)
}

func apiHjApi(router router.Party) {
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjapi")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjapi_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjApiErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetHjApiAllErrorByHash)
}

func apiHjQueue(router router.Party) {
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjqueue")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjqueue_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjQueueErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetHjQueueAllErrorByHash)
}

func apiHjAdmin(router router.Party) {
	router.Get("/", func(ctx iris.Context) {
		ctx.View("hjadmin")
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("hjadmin_errors")
	})
	router.Get("/api/list", controller.ErrorLogController.GetHjAdminErrorList)
	router.Get("/api/{errorHash}/{page}", controller.ErrorLogController.GetHjAdminAllErrorByHash)
}
