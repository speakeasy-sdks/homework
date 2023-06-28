// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EncounterVersionDetails - Version details of an encounter.
type EncounterVersionDetails struct {
	// A list of encounters and their specifics.
	EncounterDetails []Encounter `json:"encounter_details,omitempty"`
	// The total percentage of all encounter potential.
	MaxChance *int64            `json:"max_chance,omitempty"`
	Version   *NamedAPIResource `json:"version,omitempty"`
}
