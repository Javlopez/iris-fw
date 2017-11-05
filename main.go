package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/middleware/logger"
	//"github.com/kataras/iris/middleware/recover"
)


func main(){
	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context){

		ctx.ViewData("message", "Hello World!")

		ctx.View("hello.html")
	})

	app.Get("/user/{id:long}", func(ctx iris.Context){

		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)

	})

	//app.Use(recover.New())
	//app.Use(logger.New())



	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}
