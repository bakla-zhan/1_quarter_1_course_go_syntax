package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	in, err := ioutil.ReadFile("file.csv")
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	r := csv.NewReader(bytes.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
