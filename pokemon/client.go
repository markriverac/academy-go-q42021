package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const POKEMON_API = "http://pokeapi.co/api/v2/"

func pokemonClient(endpoint string) string {
	response, err := http.Get(POKEMON_API + endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}

func GetManyPokemon(many int) []Pokemon {
	pokemonResponse := pokemonClient("pokemon?limit=" + strconv.Itoa(many))
	var pokemonListResponse PokemonResponse
	json.Unmarshal([]byte(pokemonResponse), &pokemonListResponse)
	return pokemonListResponse.Results
}
