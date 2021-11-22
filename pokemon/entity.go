package pokemon

type Pokemon struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func ParseToPokemon(pokemonLines [][]string) []Pokemon {
	var pokemonList []Pokemon
	for _, line := range pokemonLines {
		newPokemon := Pokemon{
			Id:   line[0],
			Name: line[1],
		}
		pokemonList = append(pokemonList, newPokemon)
	}
	return pokemonList
}
