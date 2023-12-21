package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfigDate struct {
	OpenWeatherMapApiKey string `json:"openWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigDate, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigDate{}, err
	}
	var c apiConfigDate
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigDate{}, err

	}
	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))

}
func query(city string) (weatherData, error) {
	apiCinfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiCinfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err

	}
	d.Main.Kelvin = d.Main.Kelvin - 273.15

	return d, nil

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})
	http.ListenAndServe("localhost:8080", nil)

}
