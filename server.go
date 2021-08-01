package main

import(
	"net/http"
	"github.com/labstack/echo"
	"github.com/joho/godotenv"
	"os"
	"myapp/controller"
	database "myapp/database"
)

func main(){
	godotenv.Load()
	port := os.Getenv("PORT")
	e := echo.New() 
	database.ConnectDB()
	e.GET("/", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Welcome to the CovidTracker Application\n1.Go to /covidcases to update Covid Cases In the Mongodb\n2.Go to /getCases to Get Covid Cases In region given by user")  
	})
	
	// Run Server
	e.GET("/covidcases", controller.GetCovidCases) // UpdateCovidData endpoint
	//updatedata()

	e.POST("/getCases",controller.GetCases) //GetCovidInCoordinates endpoint
	e.Logger.Fatal(e.Start(port))
}