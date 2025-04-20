package models


//Defines type of output
type Weather struct {
	City        string  `json:"city"`
	Temperature int `json:"temperature"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
}
