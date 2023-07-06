// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MoveContestEffect struct {
	URL string `json:"url"`
}

type MoveContestType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveDamageClass struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveGeneration struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveSuperContestEffect struct {
	URL string `json:"url"`
}

type MoveTarget struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Move - Successful response
type Move struct {
	// The percent value of how likely this move is to be successful
	Accuracy      *int               `json:"accuracy,omitempty"`
	ContestCombos *ContestComboSets  `json:"contest_combos,omitempty"`
	ContestEffect *MoveContestEffect `json:"contest_effect,omitempty"`
	ContestType   *MoveContestType   `json:"contest_type,omitempty"`
	DamageClass   *MoveDamageClass   `json:"damage_class,omitempty"`
	// The percent value of the additional effects this move has occuring
	EffectChance  *int                  `json:"effect_chance,omitempty"`
	EffectChanges []AbilityEffectChange `json:"effect_changes,omitempty"`
	EffectEntries []VerboseEffect       `json:"effect_entries,omitempty"`
	Generation    MoveGeneration        `json:"generation"`
	// The identifier for this move resource
	ID   int           `json:"id"`
	Meta *MoveMetaData `json:"meta,omitempty"`
	// The name for this move resource
	Name       string               `json:"name"`
	Names      []Name               `json:"names,omitempty"`
	PastValues []PastMoveStatValues `json:"past_values,omitempty"`
	// The base power of this move with a value of 0 if it does not have a base power
	Power *int `json:"power,omitempty"`
	// Power points. The number of times this move can be used
	Pp int `json:"pp"`
	// A value of 0 means this move goes last in the turn, and 1 means it goes first
	Priority           int                     `json:"priority"`
	StatChanges        []MoveStatChange        `json:"stat_changes,omitempty"`
	SuperContestEffect *MoveSuperContestEffect `json:"super_contest_effect,omitempty"`
	Target             MoveTarget              `json:"target"`
	Type               MoveType                `json:"type"`
}
