package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/academy/academy-go-q42021/pokemon"
	"github.com/academy/academy-go-q42021/utils"

	"github.com/gorilla/mux"
)

var PokemonList []pokemon.Pokemon

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the World of Pokemon!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPokemon")
	json.NewEncoder(w).Encode(PokemonList)
}

func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	exists := false

	for _, pokemon := range PokemonList {
		if pokemon.Id == key {
			exists = true
			json.NewEncoder(w).Encode(pokemon)
		}
	}

	if !exists {
		fmt.Fprintf(w, "Pokemon doesn't exists")
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokemon", returnAllPokemon)
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	records := utils.ReadCsvFile("../pokemon/pokemon.csv")
	PokemonList = pokemon.ParseToPokemon(records)
	handleRequests()
}
