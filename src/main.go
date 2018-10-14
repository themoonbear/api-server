package main

import (
	"controller"
	"flag"
)

func main() {
	port := flag.String("p", "1323", "http listen port")
	flag.Parse()
	echo := controller.Init()
	echo.Logger.Fatal(echo.Start(":" + *port))
}
