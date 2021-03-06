package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func readAll(r io.Reader) {
	csvr := csv.NewReader(r)
	for {
		_, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
}

func main() {
	// This is a 3.6GB file from a data set that can be downloaded here:
	// http://www2.census.gov/acs2010_5yr/pums/csv_pus.zip
	huge := "../examples/data/ss10pusa.csv"
	f, err := os.Open(huge)
	if err != nil {
		log.Fatal(err)
	}
	readAll(f)
}
