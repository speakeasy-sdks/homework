// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PokemonType struct {
	// The order the Pokémon's types are listed in.
	Slot *int              `json:"slot,omitempty"`
	Type *NamedAPIResource `json:"type,omitempty"`
}

func (o *PokemonType) GetSlot() *int {
	if o == nil {
		return nil
	}
	return o.Slot
}

func (o *PokemonType) GetType() *NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.Type
}
