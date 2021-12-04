package pokemon

type Pokemon struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type PokemonResponse struct {
	Count    int       `json:"Count"`
	Next     string    `json:"Next"`
	Previous string    `json:"Previous`
	Results  []Pokemon `json:"Results"`
}

func ParseToPokemon(pokemonLines [][]string) []Pokemon {
	var pokemonList []Pokemon
	for _, line := range pokemonLines {
		newPokemon := Pokemon{
			Name: line[0],
			Url:  line[1],
		}
		pokemonList = append(pokemonList, newPokemon)
	}
	return pokemonList
}
