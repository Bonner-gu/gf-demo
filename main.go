package main

import (
	_ "gf_demo_api/boot"
	_ "gf_demo_api/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
