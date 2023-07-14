// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MoveLearnMethod - Methods by which Pokémon can learn moves.
type MoveLearnMethod struct {
	// The description of this move learn method listed in different languages.
	Descriptions []Description `json:"descriptions,omitempty"`
	// The identifier for this move learn method resource.
	ID int64 `json:"id"`
	// The name for this move learn method resource.
	Name string `json:"name"`
	// The name of this move learn method listed in different languages.
	Names []Name `json:"names,omitempty"`
}
