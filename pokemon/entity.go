package pokemon

type Pokemon struct {
	Id   string `json:"Id"`
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
			Id:   line[0],
			Name: line[1],
			Url:  line[2],
		}
		pokemonList = append(pokemonList, newPokemon)
	}
	return pokemonList
}
