package testg

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func IrisWebTest() {
	app := iris.New()
	app.Get("/", func(ctx context.Context){})
	app.Run(iris.Addr(":8080"))
}