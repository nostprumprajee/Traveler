package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://cometari-airportsfinder-v1.p.rapidapi.com/api/airports/by-text?text=Bangkok"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "cometari-airportsfinder-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "93be61accemsh1c6e7979e10e3c7p1ad963jsn20338522fb9d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// StationBeanList []Station `json:"stationBeanList"`
	// City string `json:"city"`
    // Code string `json:"code"`
    // Name string `json:"name"`
    // Location string `json:"location"`
	fmt.Println(res)
	fmt.Println(string(body),`json:"city"`)
	// fmt.Println(City)
	// fmt.Println(Code)
	// fmt.Println(Name)
	// fmt.Println(Location)

}
