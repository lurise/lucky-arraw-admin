package main

import (
	_ "luckyarrow/boot"
	_ "luckyarrow/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
