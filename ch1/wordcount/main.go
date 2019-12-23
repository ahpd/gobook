package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fn := range os.Args[1:] {
		countWords(fn, counts)
	}
	for line, count := range counts {
		fmt.Printf("%s => %03d\n", line, count)
	}

}

func countWords(fn string, counts map[string]int) {
	fmt.Printf("Loading data from file %s\n", fn)
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wc: %v\n", err)
	} else {
		for _, line := range strings.Split(string(bytes), "\n") {
			for _, word := range strings.Split(line, " ") {
				word = strings.TrimSpace(word)
				if len(word) > 0 {
					counts[word]++
				}
			}
		}
	}
}
