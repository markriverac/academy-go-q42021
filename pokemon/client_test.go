package pokemon

import "testing"

func TestGetManyPokemon(t *testing.T) {

	got := len(GetManyPokemon(150))
	want := 150

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPokemonClient(t *testing.T) {

	var got = PokemonClient("pokemon?limit=1")
	var want = "{\"count\":1118,\"next\":\"https://pokeapi.co/api/v2/pokemon?offset=1&limit=1\",\"previous\":null,\"results\":[{\"name\":\"bulbasaur\",\"url\":\"https://pokeapi.co/api/v2/pokemon/1/\"}]}"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = PokemonClient("invalidEndpoint")
	want = "Not Found"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
