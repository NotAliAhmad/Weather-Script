package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
	owm "github.com/briandowns/openweathermap"
)

func main() {
	getData()
	getWeather()
}

/*

objective:
write a go program that takes your current IP, return location, weather, forcast, etc

*/

var apikey string = "x"

type myinfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func getData() {
	res, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var info myinfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("unable to marshal, the err is : ", err)
	}
	Query := gjson.GetBytes(body, "query")

	fmt.Println("Your public IP Address is",Query)
	fmt.Println("You live in",info.RegionName)
}

func getWeather() {
	w, err := owm.NewCurrent("F", "EN", apikey)
	if err != nil{
		fmt.Printf("Something went wrong when trying to get the current: %v",err)
	}
	err = w.CurrentByName("New York")
	if err != nil{
		fmt.Printf("Unable to get the weather due to: %v",err)

	}
	fmt.Printf("The current weather is %d",(int(w.Main.Temp)))

}
