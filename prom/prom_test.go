package prom_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/prom"
)

func TestConfigFromYAML_CorrectlyParsesYAMLData(t *testing.T) {
	got, err := prom.ConfigFromYAML("testdata/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	want := prom.Config{
		Global: prom.GlobalConfig{
			ScrapeInterval:     15 * time.Second,
			EvaluationInterval: 30 * time.Second,
			ScrapeTimeout:      10 * time.Second,
			ExternalLabels: map[string]string{
				"monitor": "codelab",
				"foo":     "bar",
			},
		},
	}
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}
