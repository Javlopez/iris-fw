package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)


func main(){
	app := iris.New()

	app.Controller("/helloWorld", new(HelloWorldController))

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}

type HelloWorldController struct {
	mvc.C

	// [ Your fields here ]
	// Request lifecycle data
	// Models
	// Database
	// Global properties
}

func (c *HelloWorldController) Get() string {
	return "This is my default action... GET"
}

func (c *HelloWorldController) GetBy(name string) string {
	return "Hello " + name
}

func (c *HelloWorldController) GetWelcome() (string, int) {
	return "This is the GetWelcome action func...", iris.StatusOK
}

// GET: /helloworld/welcome/{name:string}/{numTimes:int}
func (c *HelloWorldController) GetWelcomeBy(name string, numTimes int) {
	// Access to the low-level Context,
	// output arguments are optional of course so we don't have to use them here.
	c.Ctx.Writef("Hello %s NumTimes is: %d", name, numTimes)
}