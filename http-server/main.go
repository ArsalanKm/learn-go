package main

import "github.com/labstack/echo/v4"

func helloHandler(c echo.Context) error{
	return nil
}

func main(){
	app:=echo.New()
	app.GET("/hello",helloHandler)
}