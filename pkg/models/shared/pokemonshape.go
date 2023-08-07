// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PokemonShape - Successful response
type PokemonShape struct {
	// The identifier for this Pokemon shape resource
	ID *int64 `json:"id,omitempty"`
	// The name for this Pokemon shape resource
	Name *string `json:"name,omitempty"`
}

func (o *PokemonShape) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PokemonShape) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
