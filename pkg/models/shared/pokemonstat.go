// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PokemonStat struct {
	// The base value of the stat.
	BaseStat *int64 `json:"base_stat,omitempty"`
	// The effort points (EV) the Pokémon has in the stat.
	Effort *int64            `json:"effort,omitempty"`
	Stat   *NamedAPIResource `json:"stat,omitempty"`
}
