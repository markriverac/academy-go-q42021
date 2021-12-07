package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/academy/academy-go-q42021/pokemon"
	"github.com/academy/academy-go-q42021/utils"

	"github.com/gorilla/mux"
)

var PokemonList []pokemon.Pokemon

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the World of Pokemon!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPokemon")
	json.NewEncoder(w).Encode(PokemonList)
}

func ReturnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]
	exists := false

	for _, pokemon := range PokemonList {
		if pokemon.Name == key {
			exists = true
			json.NewEncoder(w).Encode(pokemon)
		}
	}

	if !exists {
		fmt.Fprintf(w, "Pokemon doesn't exists")
	}
}

func ConcurrencyPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qType := vars["type"]
	qItems, itemsErr := strconv.Atoi(vars["items"])
	qWorkerItems, iWorkerErr := strconv.Atoi(vars["items_per_worker"])
	var concPokemonList []pokemon.Pokemon

	if itemsErr != nil || iWorkerErr != nil || !utils.IsValidKey(qType, []string{"even", "odd"}) {
		fmt.Fprintf(w, "Please check your queries. Type must be 'odd' or 'even', items and items_per_worker must be numbers")
	} else {

		workers := qItems / qWorkerItems

		items := make(chan pokemon.Pokemon, qItems)
		cPokemon := make(chan pokemon.Pokemon, qItems)

		var wg sync.WaitGroup

		for w := 0; w < workers; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				worker(qType, items, cPokemon)
			}()
		}

		for j := 1; j <= qItems; j++ {
			if j > (len(PokemonList) - 1) {
				fmt.Fprintf(w, "You are asking for too many pokemon. Wait for the next generation, maybe we can get that many then. In the meanwhile, here you have as many as we can give :D \n")
				break
			}
			items <- PokemonList[j]
		}

		close(items)
		wg.Wait()
		close(cPokemon)

		for pokemon := range cPokemon {
			concPokemonList = append(concPokemonList, pokemon)
		}
		json.NewEncoder(w).Encode(concPokemonList)
	}
}

func worker(qType string, pokemonList <-chan pokemon.Pokemon, results chan<- pokemon.Pokemon) {

	for pokemon := range pokemonList {
		pokemonId, _ := strconv.Atoi(pokemon.Id)
		switch qType {
		case "even":
			if utils.IsEven(pokemonId) {
				results <- pokemon
			}
		case "odd":
			if !utils.IsEven(pokemonId) {
				results <- pokemon
			}
		}
	}
}

func HandleRequests() {

	pokemonList := pokemon.GetManyPokemon(-1)
	utils.WritePokemonToCsv(pokemonList)
	records := utils.ReadCsvFile("pokemon.csv")
	PokemonList = pokemon.ParseToPokemon(records)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/pokemon", ReturnAllPokemon)
	myRouter.HandleFunc("/pokemon/{name}", ReturnSinglePokemon)
	myRouter.HandleFunc("/concurrency", ConcurrencyPokemon).Queries("type", "{type}", "items", "{items}", "items_per_worker", "{items_per_worker}")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
