package controller
import (
  "encoding/json" 
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
  "github.com/labstack/echo"
  "go.mongodb.org/mongo-driver/bson"
  "context"
  "time"
  "go.mongodb.org/mongo-driver/bson/primitive"
  database "myapp/database"
)
 
 
// this is array of cases
// we must add bson keyword for mongodb library autodetect fields
type Response []struct {
  ID               primitive.ObjectID `json:"id";bson:"_id"`
  ActiveCases      int    `json:"activeCases";bson:"activeCases"`
  ActiveCasesNew   int    `json:"activeCasesNew";bson:"activeCasesNew"`
  Recovered        int    `json:"recovered";bson:"recovered"`
  RecoveredNew     int    `json:"recoveredNew";bson:"recoveredNew"`
  Deaths           int    `json:"deaths";bson:"deaths"`
  DeathsNew        int    `json:"deathsNew";bson:"deathsNew"`
  PreviousDayTests int    `json:"previousDayTests";bson:"previousDayTests"`
  TotalCases       int    `json:"totalCases";bson:"totalCases"`
  SourceURL        string `json:"sourceUrl";bson:"sourceUrl"`
  ReadMe           string `json:"readMe";bson:"readMe"`
  RegionData       []struct {
    Region        string `json:"region";bson:"region"`
    ActiveCases   int    `json:"activeCases";bson:"activeCases"`
    NewInfected   int    `json:"newInfected";bson:"newInfected"`
    Recovered     int    `json:"recovered";bson:"recovered"`
    NewRecovered  int    `json:"newRecovered";bson:"newRecovered"`
    Deceased      int    `json:"deceased";bson:"deceased"`
    NewDeceased   int    `json:"newDeceased";bson:"newDeceased"`
    TotalInfected int    `json:"totalInfected";bson:"totalInfected"`
  } `json:"regionData";bson:"regionData"`
}
 
func structToBsonDocument(v interface{}) (doc *bson.D, err error) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
 
    err = bson.Unmarshal(data, &doc)
    return
}
 
 
func updatedata(response Response){
  // now we need to define our database and collection
  quickstartDatabase := database.MI.Client.Database("covid") // maybe we need create this database before 
 
  // set collection
  articleCollection := quickstartDatabase.Collection("cases")
 
    // set insert timeout
  ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
 
 
  // for each all covid cases
  for _, v := range response {
      // now we must use 'v' (current response case)
      // but we need set by hand every field :(
      // but if we define struc like mongodb document, can parse with libreary :)
      parsedStruct, parseStructToDocumentErr := structToBsonDocument(v)
      if parseStructToDocumentErr  != nil{
        log.Fatal(parseStructToDocumentErr)
      }
      fmt.Println(parsedStruct) 
 
 
      covidCaseInsertResult, err := articleCollection.InsertOne(ctx, parsedStruct) // now this throw syntax error, we need parse before 
 
      /*  
      articleResult, err := articleCollection.InsertOne(ctx, bson.D{
        {Key: "title", Value: "The Polyglot Developer Podcast"},
        {Key: "author", Value: "Nic Raboy"},
      })
 
      */
      if err  != nil{
        log.Fatal(err)
      }
      fmt.Printf("Inserted documents into episode collection!\n")
      // // lo.Printf(len(articleResult))
      // objectID := articleResult.InsertedID.(primitive.ObjectID)
      fmt.Println(covidCaseInsertResult)
  }
}
// GetCovidCases - handler method for getting covid cases from API
func GetCovidCases(c echo.Context) (err error) {
  fmt.Println("Calling API...")
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://api.apify.com/v2/datasets/6isWRcvjRrcmlSaYw/items", nil)
  if err != nil {
    fmt.Print(err.Error())
  }
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/json")
  resp, err := client.Do(req) //making request
  if err != nil {
    fmt.Print(err.Error())
  }
  defer resp.Body.Close()
  log.Println(resp.Body)
  bodyBytes, err := ioutil.ReadAll(resp.Body) //reading resonse
  if err != nil {
    fmt.Print(err.Error())
  }
  var responseObject Response //object of struct type Response
  json.Unmarshal(bodyBytes, &responseObject) //unmarshaling the response
  //fmt.Printf("API Response as struct %+v\n", responseObject)
  updatedata(responseObject)
  return c.JSON(http.StatusOK,responseObject) 
}


func GetCases(c echo.Context) error {
  lat := c.FormValue("lat")
  lng := c.FormValue("lng")
  fmt.Println(lat)
  fmt.Println(lng)
  return c.String(http.StatusOK, "hi")
}