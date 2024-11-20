package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//reading
	//read the file
	data := ReadFile("1.csv")

	//create reader
	reader := GetCSVReader(data)

	//parse csv
	ParseCSV(reader)

	//writing
	records := [][]string{
		{"John", "30", "New York"},
		{"Alice", "25", "Tokyo"},
		{"Bob", "35", "London"},
	}

	var Filename string = "newFile.csv"
	writer := CreateFile(Filename)
	WriteRecords(writer, records)

}

func WriteRecords(writer *os.File, records [][]string) {
	CSVWriter := csv.NewWriter(writer)
	CSVWriter.WriteAll(records)
}

func CreateFile(filename string) *os.File {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating writer intance for", filename)
	}
	return writer
}

func ParseCSV(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			log.Print("Reached End of File")
			break
		}
		if err != nil {
			log.Fatal("Error parsing csv")
		}
		fmt.Println(record)
	}

}

func ReadFile(filename string) []byte {
	File, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file", filename)
	}
	data, err := io.ReadAll(File)
	if err != nil {
		log.Fatal("Error reading file", filename)
	}
	return data
}

func GetCSVReader(data []byte) *csv.Reader {
	reader := bytes.NewReader(data)
	CSVReader := csv.NewReader(reader)
	return CSVReader
}
