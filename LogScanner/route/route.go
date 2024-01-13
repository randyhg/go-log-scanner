package route

import (
	"go-log-scanner/LogScanner/cache"
	"go-log-scanner/LogScanner/controller"
	"go-log-scanner/config"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func RegisterRoutes(app *iris.Application) {
	tmpl := iris.HTML("./LogScanner/views", ".html")
	app.RegisterView(tmpl)
	apiTotal(app.Party("/total"))

	opts := basicauth.Options{
		Allow: basicauth.AllowUsers(map[string]string{
			config.Instance.Username: config.Instance.Password,
		}),
		Realm:        "Authorization Required",
		ErrorHandler: basicauth.DefaultErrorHandler,
	}
	auth := basicauth.New(opts)
	app.Use(auth)

	app.Get("/", func(ctx iris.Context) {
		data := iris.Map{
			"Title": "HJ Errors Log",
		}
		ctx.View("part/header", data)
		ctx.View("index")
		ctx.View("part/footer")
	})

	apiHjm3u8(app.Party("/hjm3u8"))
	apiChatServer(app.Party("/chat_server"))
	apiHjAppServer(app.Party("/hjapp_server"))
	apiHjApi(app.Party("/hjapi"))
	apiHjQueue(app.Party("/hjqueue"))
	apiHjAdmin(app.Party("/hjadmin"))

}

func apiTotal(router router.Party) {
	router.Get("/chat_server", cache.HandleCache(controller.ErrorLogController.GetChatServerTotalRecord))
	router.Get("/hjadmin", cache.HandleCache(controller.ErrorLogController.GetHjAdminTotalRecord))
	router.Get("/hjapi", cache.HandleCache(controller.ErrorLogController.GetHjApiTotalRecord))
	router.Get("/hjapp_server", cache.HandleCache(controller.ErrorLogController.GetHjAppServerTotalRecord))
	router.Get("/hjm3u8", cache.HandleCache(controller.ErrorLogController.GetHjM3u8TotalRecord))
	router.Get("/hjqueue", cache.HandleCache(controller.ErrorLogController.GetHjQueueTotalRecord))
}

func apiHjm3u8(router router.Party) {
	data := iris.Map{
		"Title": "HJ M3u8 Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/table", data)
		ctx.View("part/footer", data)
		ctx.View("hjm3u8", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/error_table", data)
		ctx.View("part/footer", data)
		ctx.View("hjm3u8_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetHjm3u8ErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjm3u8AllErrorByHash)
}

func apiChatServer(router router.Party) {
	data := iris.Map{
		"Title": "Chat Server Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/table", data)
		ctx.View("part/footer", data)
		ctx.View("chat_server", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/error_table", data)
		ctx.View("part/footer", data)
		ctx.View("chat_server_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetChatServerErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetChatServerAllErrorByHash)
}

func apiHjAppServer(router router.Party) {
	data := iris.Map{
		"Title": "HJ App Server Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/table", data)
		ctx.View("part/footer", data)
		ctx.View("hjapp_server", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/error_table", data)
		ctx.View("part/footer", data)
		ctx.View("hjapp_server_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetHjAppServerErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjAppServerAllErrorByHash)
}

func apiHjApi(router router.Party) {
	data := iris.Map{
		"Title": "HJ Api Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/table", data)
		ctx.View("part/footer", data)
		ctx.View("hjapi", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/error_table", data)
		ctx.View("part/footer", data)
		ctx.View("hjapi_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetHjApiErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjApiAllErrorByHash)
}

func apiHjQueue(router router.Party) {
	data := iris.Map{
		"Title": "HJ Queue Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("hjqueue", data)
		ctx.View("part/footer", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("hjqueue_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetHjQueueErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjQueueAllErrorByHash)
}

func apiHjAdmin(router router.Party) {
	data := iris.Map{
		"Title": "HJ Admin Error Log",
	}
	router.Get("/", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/table", data)
		ctx.View("part/footer", data)
		ctx.View("hjadmin", data)
	})
	router.Get("/errors", func(ctx iris.Context) {
		ctx.View("part/header", data)
		ctx.View("part/error_table", data)
		ctx.View("part/footer", data)
		ctx.View("hjadmin_errors", data)
	})
	router.Get("/api/list", cache.HandleCache(controller.ErrorLogController.GetHjAdminErrorList))
	router.Get("/api/{errorHash}", controller.ErrorLogController.GetHjAdminAllErrorByHash)
}
