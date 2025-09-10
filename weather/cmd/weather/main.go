package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nw-code/tpg-tools/weather"
)

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
	conditions, err := weather.Get(apiKey, latitude, longitude)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(conditions)
}
