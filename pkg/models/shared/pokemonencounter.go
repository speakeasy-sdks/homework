// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PokemonEncounter - Encounters Pokemon that can be encountered in the game and the version groups in which they can be encountered.
type PokemonEncounter struct {
	Pokemon        *NamedAPIResource         `json:"pokemon,omitempty"`
	VersionDetails []EncounterVersionDetails `json:"version_details,omitempty"`
}
