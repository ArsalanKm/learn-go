package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct{
	Name string
	Family string
	Verb int
}

func helloHandler(c echo.Context) error{
	return c.JSON(http.StatusOK,Response{
		Name: "arsalan",
		Family: "Karimzad",
		Verb: 10,
	})
}

func main(){
	app:=echo.New()
	app.GET("/hello",helloHandler)
	if err :=app.Start(":1379");err != nil{
		log.Println(err)
	}
}