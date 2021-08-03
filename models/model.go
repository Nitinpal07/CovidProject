package model

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
type Result struct{
  State string `json:"state"`
  ActiveCases int32 `json:"activecases"`
  LastUpdatedAtApify primitive.DateTime `json:"lastUpdatedAtApify"`
  TotalCasesInIndia int32 `json:"totalcasesinindia"`
}