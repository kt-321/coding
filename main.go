package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

type CityController struct {
	CityInteractor CityInteractor
}

type CityInteractor struct {
	CityRepository CityRepository
}

func (ci *CityInteractor) Show(url string) (*City, error) {
	return ci.CityRepository.getCity(url)
}

type CityRepository struct {}

type CityRepositoryInterface interface {
	getCity(string) (*City, error)
}

func NewCityController() *CityController {
	return &CityController{
		CityInteractor: CityInteractor{
			CityRepository: CityRepository{},
		},
	}
}

func (cc *CityController) GetCityHandler(w http.ResponseWriter, r *http.Request) {
	//city, err := cc.CityInteractor.getCity()
}


func getCity(url string) (*City, error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}

	//client := &http.Client{}
	//resp, err := client.Get(url)
	//if err != nil {
	//	return nil, err
	//}

	//req, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var city City
	if err:= json.Unmarshal(result, &city); err != nil {
		return nil, err
	}

	return &city, nil
}

func Dispatch() {
	cityController := NewCityController()

	r := mux.NewRouter()
	r.HandlerFunc("/api/cirty/{url}", cityController.GetCityHandler).Methods("GET")

	if err := http.ListenAndServe(":8084", r); err != nil {
		log.Println(err)
	}
}

func main() {
	urls := []string{
		"https://api.openweathermap.org/data/2.5/weather?q=Tokyo&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=London&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=Paris&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=Shanghai&APPID=bc0bad98583f474e3ae5166d871415f0",
		"https://api.openweathermap.org/data/2.5/weather?q=New+York&APPID=bc0bad98583f474e3ae5166d871415f0",
	}

	for _, v := range urls {
		city, err := getCity(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v : %v\n", city.Name, city.Weather[0].Description)
	}
}