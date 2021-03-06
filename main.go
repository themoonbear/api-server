package main

import (
	"flag"

	"api-server/controller"
)

func main() {
	port := flag.String("p", "1324", "http listen port")
	flag.Parse()
	echo := controller.Init()
	echo.Logger.Fatal(echo.Start(":" + *port))
}
