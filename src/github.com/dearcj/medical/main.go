package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Drug struct {
	Id        string
	Name      string
	WorldName string
	Type      string
	Condition string
}

func parseCSV(fileName string) error {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	reader.LazyQuotes = true

	var drugs []Drug
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return error

		}
		if line != nil {
			println(line[0], line[1], line[2], line[3], line[4], line[5], line[6])
		}
		drug

		drugs = append(drugs, drug)
	}

	return nil
}

func main() {
	//https://search.usa.gov/sayt?name=fda&q=albe&callback=jQuery331006678285769227399_1538243799425&_=1538243799426
	response, err := http.Get("https://search.usa.gov/sayt?name=fda&q=albe")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s\n", string(contents))
	}

	parseCSV("reestr2.csv")
}
