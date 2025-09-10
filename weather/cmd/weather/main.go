package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/nw-code/tpg-tools/weather"
)

type CurrentWeather struct {
	Main map[string]float32 `json:"main"`
}

const Usage = `Usage: weather LATITUDE LONGITUDE

Example: weather -37.654 145.5172`

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	latitude, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error converting arg %s to float64\n", os.Args[1])
		os.Exit(1)
	}
	longitude, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error converting arg %s to float64\n", os.Args[2])
		os.Exit(1)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "missing environment variable API_KEY")
		os.Exit(1)
	}
	baseURL := "https://api.openweathermap.org"
	URL := weather.FormatURL(baseURL, latitude, longitude, apiKey)


	r, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "unexpected response status", r.Status)
		os.Exit(1)
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unexpected issue reading response body")
		os.Exit(1)
	}
	conditions, err := weather.ParseResponse(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(conditions)
}
