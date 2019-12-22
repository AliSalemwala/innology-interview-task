package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"bytes"
)

func main() {
	jsonStr := []byte(`{"lat":41.0253611, "lng":29.0598525}`)
	req, _ := http.NewRequest("GET", "http://localhost:2210/region?lat=41.057808&lng=29.008149", bytes.NewBuffer(jsonStr))
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf(err.Error())
    }
	
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Region Response:", string(body))


	jsonStr = []byte(`{
						"from":{ "lat":41.057808, "lng":29.008149},
						"to":{ "lat":41.0253611, "lng":29.0598525}
					}`)
	req, _ = http.NewRequest("GET", "http://localhost:2210/fare", bytes.NewBuffer(jsonStr))
    
    resp, err = client.Do(req)
    if err != nil {
        fmt.Printf(err.Error())
    }
	
    body, _ = ioutil.ReadAll(resp.Body)
    fmt.Println("Fare Response:", string(body))
}