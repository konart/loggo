package main

//import (
//	"net/http"
//
//	"github.com/labstack/echo"
//	"fmt"
//	"github.com/labstack/echo/middleware"
//	"github.com/labstack/gommon/log"
//)
//
//func main() {
//	e := echo.New()
//	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//		Skipper: middleware.DefaultSkipper,
//		Format: "method=${method}, uri=${uri}, status=${status}\n",
//	}))
//
//	//e.Use(middleware.Logger())
//	e.Logger.SetLevel(log.ERROR)
//	e.GET("/", func(c echo.Context) error {
//		return c.String(http.StatusOK, "Hello, World!")
//	})
//	e.GET("/users/:id", getUser)
//	e.GET("/input", showForm)
//	e.POST("/input", handleForm)
//	e.Logger.Fatal(e.Start(":1323"))
//
//}
//
//func getUser(c echo.Context) error {
//	return c.HTML(http.StatusOK, "<html><body><h1>test</h1></body></html>")
//}
//
//func showForm(c echo.Context) error  {
//	return c.HTML(http.StatusOK, "<html><body><form action=\"/input\" method=\"post\">input <input type=\"text\" name=\"input\"></form></body></html>")
//}
//
//func handleForm(c echo.Context) error {
//	input := c.FormValue("input")
//	return c.HTML(http.StatusOK, fmt.Sprintf("<html><body><h1>%s</h1></body></html>", input))
//}
