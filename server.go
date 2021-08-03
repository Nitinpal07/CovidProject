package main

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"path"
	"myapp/controller"
	database "myapp/database"
	_ "myapp/docs/echosimple"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main(){
	fmt.Println("Go Program")
	server := echo.New()
	database.ConnectDB()
	server.GET(path.Join("/"), homepage)
	server.GET("/covidcases", controller.GetCovidCases) // UpdateCovidData endpoint
	//updatedata()
	server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.GET("/getCases",controller.GetCases) //GetCovidInCoordinates endpoint
	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)
}

func homepage(context echo.Context) error {
	return context.HTML(http.StatusOK,"<strong>Welcome to the Covid Tracker  API</strong><ol><li>Go to /covidcases to update Covid Cases In the Mongodb Database</li>  <li>Go to /getCases?lat=xx&amp;lng=xx to Get Covid Cases In state where the gps coordinates(lat,lng) lies</li></ol>")
}

