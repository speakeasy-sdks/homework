// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GenderPokemonSpeciesDetails struct {
	// The chance of this Pokémon being female, in eighths; or -1 for genderless
	Rate *int64 `json:"rate,omitempty"`
}

func (o *GenderPokemonSpeciesDetails) GetRate() *int64 {
	if o == nil {
		return nil
	}
	return o.Rate
}

// Gender - Successful response
type Gender struct {
	// The identifier for this gender resource
	ID *int64 `json:"id,omitempty"`
	// The name for this gender resource
	Name                  *string                       `json:"name,omitempty"`
	PokemonSpeciesDetails []GenderPokemonSpeciesDetails `json:"pokemon_species_details,omitempty"`
}

func (o *Gender) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Gender) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Gender) GetPokemonSpeciesDetails() []GenderPokemonSpeciesDetails {
	if o == nil {
		return nil
	}
	return o.PokemonSpeciesDetails
}
