package main

import(
	"net/http"
	"github.com/labstack/echo"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"path"
	"myapp/controller"
	database "myapp/database"
)

// func main(){
// 	godotenv.Load()
// 	port := os.Getenv("PORT")
// 	e := echo.New() 
// 	database.ConnectDB()
// 	e.GET("/", func(c echo.Context) error {       
// 	return c.String(http.StatusOK, "Welcome to the CovidTracker Application\n1.Go to /covidcases to update Covid Cases In the Mongodb\n2.Go to /getCases to Get Covid Cases In region given by user")  
// 	})
	
// 	// Run Server
// 	e.GET("/covidcases", controller.GetCovidCases) // UpdateCovidData endpoint
// 	//updatedata()

// 	e.POST("/getCases",controller.GetCases) //GetCovidInCoordinates endpoint
// 	e.Logger.Fatal(e.Start(port))
// }

func main(){
	fmt.Println("Go Program")
	server := echo.New()
	database.ConnectDB()
	server.GET(path.Join("/"), Version)
	server.GET("/covidcases", controller.GetCovidCases) // UpdateCovidData endpoint
	//updatedata()

	server.GET("/getCases/:lat/:lng",controller.GetCases) //GetCovidInCoordinates endpoint
	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)
}

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, "Welcome to the CovidTracker Application\n1.Go to /covidcases to update Covid Cases In the Mongodb\n2.Go to /getCases to Get Covid Cases In region given by user")
}
