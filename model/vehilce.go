package model

type Vehilce struct {
	Id     string  `json:"id"`
	Brand  string  `json:"brand"`
	Model  string  `json:"model"`
	Year   int     `json:"year"`
	Weight float64 `json:"weight"`
}
