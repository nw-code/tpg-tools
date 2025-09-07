package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type freq struct {
	ipAddress string
	count     int
}

func compileResults(filename string) ([]freq, error) {
	results := map[string]int{}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		freq := scanner.Text()
		fields := strings.Fields(freq)
		if len(fields) > 0 {
			results[fields[0]]++
		}
	}

	freqs := make([]freq, 0, len(results))
	for k, v := range results {
		freqs = append(freqs, freq{k, v})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	return freqs, nil
}

func main() {
	results, err := compileResults("log.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%-16s%s\n", "Address", "Requests")
	for i, freq := range results {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s%d\n", freq.ipAddress, freq.count)
	}
}
