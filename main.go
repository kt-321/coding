package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type City struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}


//API叩いて都市情報取得
func getCity(url string) (*City, error) {
	var city City

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}


	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &city); err != nil {
		return nil, err
	}

	return &city, nil
}




func main() {
	urls := []string {
		"https://api.openweathermap.org/data/2.5/weather?q=Tokyo&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=London&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=Paris&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=Shanghai&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=New+York&APPID=bc0bad98583f474e3ae5166d871415f0",
	}

	for _, v := range urls {
		result, err := getCity(v)
		if err != nil {
			fmt.Println(err)
		}

		w := result.Weather[0]
		fmt.Printf("%v : %v\n", result.Name, w.Description)
	}
}

