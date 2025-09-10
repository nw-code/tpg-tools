package weather_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/weather"
)

func TestParseResponse_CorrectlyParsesJSONData(t *testing.T) {
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	want := weather.Conditions{
		Summary: "Clouds",
	}
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestGetWeather_ReturnsExpectedConditions(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/weather.json")
		}))
	defer ts.Close()
	apiKey := "662e2f7b337efe11afc601a8ad6d36ff"
	c := weather.NewClient(apiKey)
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	latitude := -37.654
	longitude := 145.5172
	got, err := c.GetWeather(latitude, longitude)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary: "Clouds",
	}
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestParseResponse_ErrorsOnEmptyData(t *testing.T) {
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing emtpy response, got nil")
	}
}

func TestParseResponse_ReturnsEmptyConditionsForMissingWeather(t *testing.T) {
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error parsing invalid response, got nil")
	}
}

func TestFormatURL_ReturnsCorrectlyFormattedURL(t *testing.T) {
	apiKey := "662e2f7b337efe11afc601a8ad6d36ff"
	c := weather.NewClient(apiKey)
	latitude := -37.654
	longitude := 145.5172
	got := c.FormatURL(latitude, longitude)
	want := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", latitude, longitude, apiKey)
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestHTTPS(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/weather.json")
	}))
	defer ts.Close()
	client := ts.Client()
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	got := resp.StatusCode
	want := http.StatusOK
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}
