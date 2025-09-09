package main

import (
//	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CurrentWeather struct {
	Main map[string]float32 `json:"main"`
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "missing environment variable API_KEY")
		os.Exit(1)
	}
	baseURL := "https://api.openweathermap.org"
	URL := fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", baseURL, -37.654, 145.5172, apiKey)
	r, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stdout, "unexpected response status", r.Status)
		os.Exit(1)
	}
/*
	weather := CurrentWeather{}
	err = json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(weather.Main["temp"])
*/
	io.Copy(os.Stdout, r.Body)
}
