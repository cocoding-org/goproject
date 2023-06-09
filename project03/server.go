package project03

import (
	project03 "cocoding-org/goproject/project03/page"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
)

func Server() {
	http.HandleFunc("/", project03.IndexPage)
	http.HandleFunc("/hello", project03.HelloPage)
	http.HandleFunc("/html", project03.HtmlPage)
	http.HandleFunc("/post", project03.PostAction)
	log.Println("start server on port:", 8099)
	log.Fatal(http.ListenAndServe("localhost:8099", nil))
}

func Iris() {
	app := iris.Default()

	app.Use(myMiddleware)

	app.RegisterView(iris.HTML("./project03/views", ".html"))

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		//输出JSON
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.Post("/hello", func(ctx iris.Context) {
		//输出字符串
		ctx.WriteString("hi," + ctx.FormValue("username"))
	})
	app.Handle("GET", "/", func(ctx iris.Context) {
		// 设置消息
		ctx.ViewData("message", "hello golang")
		//加载模板
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8098"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Run iris before %s", ctx.Path())
	ctx.Next()
}
