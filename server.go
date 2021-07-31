package main

import(
	"net/http"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"myapp/controller"
	database "myapp/database"
)

func main(){
	// client,ctx := connectMongo()
	e := echo.New() 
	// // Middleware
	// e.Use(middleware.Logger()) // Logger 
	// e.Use(middleware.Recover()) // Recover
	// // CORS
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// AllowOrigins: []string{"*"},
	// AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH,    echo.POST, echo.DELETE},
	// }))
	database.ConnectDB()
	e.GET("/", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Hello, World!\n")  
	})
	
	// Run Server
	e.GET("/covidcases", controller.GetCovidCases) // Price endpoint
	//updatedata()

	e.POST("/getCases",controller.GetCases)
	e.Logger.Fatal(e.Start(":8000"))
}