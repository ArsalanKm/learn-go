package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error{
	return c.JSON(http.StatusOK,"Hello World From Glolang")
}

func main(){
	app:=echo.New()
	app.GET("/hello",helloHandler)
	if err :=app.Start(":1379");err != nil{
		log.Println(err)
	}
}