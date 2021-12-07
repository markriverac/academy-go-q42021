package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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
	_ = csvwriter.Write([]string{"Name", "url"})
	for _, empRow := range pokemonList {
		pokemonData := []string{empRow.Name, empRow.Url}
		_ = csvwriter.Write(pokemonData)
	}

	csvwriter.Flush()
	csvFile.Close()
}
