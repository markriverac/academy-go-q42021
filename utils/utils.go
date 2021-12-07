package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/academy/academy-go-q42021/pokemon"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func WritePokemonToCsv(pokemonList []pokemon.Pokemon) {
	csvFile, err := os.Create("pokemon.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	_ = csvwriter.Write([]string{"id", "Name", "url"})
	for index, empRow := range pokemonList {
		pokemonData := []string{strconv.Itoa(index + 1), empRow.Name, empRow.Url}
		_ = csvwriter.Write(pokemonData)
	}

	csvwriter.Flush()
	csvFile.Close()
}

func IsEven(id int) bool {
	return id%2 == 0
}
