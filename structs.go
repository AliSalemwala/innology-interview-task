package main

// structs to help with parsing the json objects returned by the Google APIs

type Coordinates struct{
	Latitude	float64 `json:"lat"`
	Longitude	float64 `json:"lng"`
}

type ToFro	struct{
	From	Coordinates `json:"from"`
	To		Coordinates `json:"to"`
}

type AC struct {
	Name		string	`json:"long_name"`
	Types		[]string `json:"types"`
}

type Result struct {
	AddrComps	[]AC `json:"address_components"`
	Types		[]string `json:"types"`
}

type GeocodeResponse struct {
	Results		[]Result `json:"results"`
	Status		string	`json:"status"`
}

type Val struct {
	Value		int `json:"value"`
}

type Element struct {
	Distance	Val `json:"distance"`
	Duration	Val `json:"duration"`
}

type Row struct {
	Elements	[]Element `json:"elements"`
}

type DistMatResponse struct{
	Rows		[]Row `json:"rows"`
	Status		string	`json:"status"`
}