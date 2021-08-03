package util

import(

)

type Config struct {
    PORT string `mapstructure:"PORT"`
    DBSource      string `mapstructure:"DB_SOURCE"`
    ReverseGeoCodingApiKey string `mapstructure:"REVERSE_GEOCODING_API_KEY"`
}