package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"math"
	"io"
)

// if longitude is to the left of the Bosphorus Bridge, returns Europe
// otherwise returns Asia
func GetContinent (loc Coordinates) string{
	if loc.Longitude <= 29.0302045 {
		return "Europe"
	}
	return "Asia"
}

// parses stream r into json-style struct dest
func GetJson (r io.ReadCloser, dest interface{}) {
	error := json.NewDecoder(r).Decode(dest)
	defer r.Close()
	if error != nil{
		panic (error)
	}
}

// parses response from url into json-style struct dest
func GetJsonFromURL (dest interface{}, url string, args ...interface{}){
	resp, error := http.Get (fmt.Sprintf(url, args...))
	if error != nil{
		panic (error)
	}
	GetJson (resp.Body, dest)
}


func GetRegion (w http.ResponseWriter, r *http.Request){
	defer Recoverer()
	var geo GeocodeResponse
	var loc Coordinates
	
	// gets the required json-style structs
	GetJson (r.Body, &loc)
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=AIzaSyByrWQikfLWsF1Q57WYLtTrJb_vSutb2uo"
	GetJsonFromURL (&geo, url, loc.Latitude, loc.Longitude)
	
	// very deep nest
	// if area exists
	// finds city name ("locality")
	// if city name is Istanbul sends response json
	// else error
	if geo.Status == "OK"{
		for _, result := range geo.Results{
			for _, resType := range result.Types{
				if resType == "locality"{
					for _, ac := range result.AddrComps{
						for _, actype := range ac.Types{
							if actype == "locality"{
								if city := ac.Name;  city == "Istanbul" || city == "İstanbul"{
									region := struct{
										Region	string `json:"region"`
									}{GetContinent (loc)}		
									json.NewEncoder(w).Encode(region)
								} else{
									http.Error(w, "", http.StatusBadRequest)
								}
								return
							}
						}
					}
				}
			}
		}
	}
	http.Error (w, "No result found", http.StatusBadRequest)
}

// determines fare for dist
func CalcFare (dist int) int {
	return int (100 * math.Max (10, 5 + float64 (dist) /100 * 0.2))
}


func GetFare (w http.ResponseWriter, r *http.Request){
	defer Recoverer()
	var locs ToFro
	var dmat DistMatResponse
	
	// gets the required json-style structs
	GetJson (r.Body, &locs)
	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=%f,%f&destinations=%f,%f&key=AIzaSyDLhWmmuwRAfjtq9mseCr6Fyd98TrSHIxE"
	GetJsonFromURL (&dmat, url, locs.From.Latitude, locs.From.Longitude, locs.To.Latitude, locs.To.Longitude)
	
	// if valid route available
	// calculates fare
	// sends response json object
	// else error
	if dmat.Status == "OK"{
		temp := dmat.Rows[0].Elements[0]
		fare := struct{
			Duration	int `json:"duration"`
			Distance	int `json:"distance"`
			Fare		int `json:"fare"`
		}{temp.Duration.Value, 
			temp.Distance.Value, 
			CalcFare (temp.Distance.Value)}
		json.NewEncoder(w).Encode(fare)
	} else{
		http.Error (w, "No result found", http.StatusBadRequest)
	}
}

// informs of error in case of panic
func Recoverer(){
	if r:= recover(); r != nil{
		fmt.Println ("There was a problem: ", r)
	}
}

func main(){
	http.HandleFunc("/fare", GetFare)
	http.HandleFunc("/region", GetRegion)
	http.ListenAndServe(":2210", nil)
}