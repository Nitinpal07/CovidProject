package model

import()

type AutoGenerated []struct {
	ActiveCases      int    `json:"activeCases"`
	ActiveCasesNew   int    `json:"activeCasesNew"`
	Recovered        int    `json:"recovered"`
	RecoveredNew     int    `json:"recoveredNew"`
	Deaths           int    `json:"deaths"`
	DeathsNew        int    `json:"deathsNew"`
	PreviousDayTests int    `json:"previousDayTests"`
	TotalCases       int    `json:"totalCases"`
	SourceURL        string `json:"sourceUrl"`
	ReadMe           string `json:"readMe"`
	RegionData       []struct {
		Region        string `json:"region"`
		ActiveCases   int    `json:"activeCases"`
		NewInfected   int    `json:"newInfected"`
		Recovered     int    `json:"recovered"`
		NewRecovered  int    `json:"newRecovered"`
		Deceased      int    `json:"deceased"`
		NewDeceased   int    `json:"newDeceased"`
		TotalInfected int    `json:"totalInfected"`
	} `json:"regionData"`
}