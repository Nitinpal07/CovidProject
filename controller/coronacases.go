package controller
import (
  "encoding/json" 
  "fmt"
  "io/ioutil"
  "reflect"
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
  LastUpdatedAtApify time.Time `json:"lastUpdatedAtApify";bson:"lastUpdatedAtApify"`
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

// structToBsonDocument - method to convert golang struct to bson document
func structToBsonDocument(v interface{}) (doc *bson.D, err error) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
 
    err = bson.Unmarshal(data, &doc)
    return
}
 
// UpdateData - handler method for updating covid data in mongodb
func UpdateData(response Response){
  // now we need to define our database and collection
  quickstartDatabase := database.MI.Client.Database("covid") // maybe we need create this database before 
 
  // set collection
  covidCollection := quickstartDatabase.Collection("cases")
 
  // set insert timeout
  ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
 
  // Call the DeleteOne() method by passing BSON
  res, err := covidCollection.DeleteOne(ctx, bson.M{})
  fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))
  if err != nil {
    log.Fatal("DeleteOne() ERROR:", err)
  }
  // for each all covid cases
  for _, v := range response {
      // now we must use 'v' (current response case)
      // but we need set by hand every field :(
      // but if we define struc like mongodb document, can parse with library :)
      //fmt.Println("79")
      //fmt.Println(v)
      parsedStruct, parseStructToDocumentErr := structToBsonDocument(v)
      if parseStructToDocumentErr  != nil{
        log.Fatal(parseStructToDocumentErr)
      }
      //fmt.Println("84")
      //fmt.Println(parsedStruct) 
 
 
      covidCaseInsertResult, err := covidCollection.InsertOne(ctx, parsedStruct) // now this throw syntax error, we need parse before 
      if err  != nil{
        log.Fatal(err)
      }
      fmt.Printf("Inserted documents into cases collection!\n")
      fmt.Println(covidCaseInsertResult)
  }
}
// GetCovidCases - handler method for getting covid cases from API
func GetCovidCases(c echo.Context) (err error) {
  fmt.Println("Calling API...")
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://api.apify.com/v2/actor-tasks/81uAfgDHw6d8n1eNd/run-sync-get-dataset-items?token=p2CpMHQHSjXuayE8HPcdYaXjc", nil)
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
  //log.Println(resp.Body)
  bodyBytes, err := ioutil.ReadAll(resp.Body) //reading resonse
  if err != nil {
    fmt.Print(err.Error())
  }
  var responseObject Response //object of struct type Response
  json.Unmarshal(bodyBytes, &responseObject) //unmarshaling the response
  //fmt.Printf("API Response as struct %+v\n", responseObject)
  UpdateData(responseObject)
  return c.JSON(http.StatusOK,"CovidCases Updated in MongoDb") 
}

// GetCasesInState - handler method for getting covid cases in a state from database
func GetCasesInState(state string){

  quickstartDatabase := database.MI.Client.Database("covid") // maybe we need create this database before 
  // set collection
  covidCollection := quickstartDatabase.Collection("cases")
  ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

  var response bson.M
  var err error
  if err = covidCollection.FindOne(ctx, bson.M{}).Decode(&response); err != nil {
    log.Fatal(err)
  }
  fmt.Println("activecases")
  fmt.Println(response["activecases"])
  fmt.Println("-----------------")

  fmt.Println("lastupdatetime")
  fmt.Println(response["lastupdatedatapify"])
  fmt.Println("-----------------")
  
  fmt.Println("totalCases")
  fmt.Println(response["totalcases"])
  fmt.Println("-----------------")
  
  var CaseInState int32
  for _,v := range response["regiondata"].(primitive.A){
    //fmt.Println(k) 
    //fmt.Println(reflect.TypeOf(v))
    //fmt.Println(v) --> map
    //var currstate string
    //var currstatecases string
    // for x,y := range v.(primitive.M){
    //   //fmt.Println("inside loop")
    //   fmt.Println(x[])
    //   fmt.Println(y)
    //   // if(x==state){
    //   //   CaseInState = 
    //   // }
    // }
   // fmt.Println(v.(primitive.M)["region"])
    if(v.(primitive.M)["region"]==state){
      fmt.Println(state)
      CaseInState = v.(primitive.M)["totalinfected"].(int32)
      fmt.Println(CaseInState)
      break
    }
    //fmt.Println("---------")
  }

}

// GetCases - handler method for getting covid cases from gps coordinates provided by user
func GetCases(c echo.Context) error {
  latitude := c.FormValue("lat")
  longitude := c.FormValue("lng")
  //fmt.Println(latitude)
  //fmt.Println(longitude)
  state := getState(latitude,longitude)
  GetCasesInState(state)
  return c.String(http.StatusOK, state)
}