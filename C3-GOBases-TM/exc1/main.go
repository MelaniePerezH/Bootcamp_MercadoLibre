package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"ID", "PRECIO", "CANTIDAD"},
		{"4561", "250000", "5"},
		{"7895", "100000", "2"},
		{"9621", "891000", "1"},
		{"7894", "890000", "10"},
	}
	myFile, err := os.Create("productos.csv")
	if err != nil {
		log.Fatalf("Falló la creacion del archivo")
	}
	cswriter := csv.NewWriter(myFile)
	for _, row := range rows {
		_ = cswriter.Write(row)
	}
	cswriter.Flush()
	myFile.Close()

	//OPEN THE FILE
	opening, err := os.Open("productos.csv")
	if err != nil {
		log.Fatalln("Can't open the file")
	}

	r := csv.NewReader(opening)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t %s\t %s\n", record[0], record[1], record[2])
	}

}
