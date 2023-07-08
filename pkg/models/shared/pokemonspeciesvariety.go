// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PokemonSpeciesVariety struct {
	// Whether this is the default "natural" variety of this species. Note that "default" is
	// subjective and that it may not match the Pokémon games' official status.
	//
	IsDefault *bool `json:"is_default,omitempty"`
	// The name of this Pokémon species variety
	Name    *string  `json:"name,omitempty"`
	Pokemon *Pokemon `json:"pokemon,omitempty"`
}
