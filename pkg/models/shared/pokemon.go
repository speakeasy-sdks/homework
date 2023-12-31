// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Pokemon struct {
	Abilities              []PokemonAbility   `json:"abilities"`
	BaseExperience         int                `json:"base_experience"`
	Forms                  []PokemonForm      `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	Height                 int                `json:"height"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	ID                     int                `json:"id"`
	IsDefault              bool               `json:"is_default"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                PokemonSprites     `json:"sprites"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
	Weight                 int                `json:"weight"`
}

func (o *Pokemon) GetAbilities() []PokemonAbility {
	if o == nil {
		return []PokemonAbility{}
	}
	return o.Abilities
}

func (o *Pokemon) GetBaseExperience() int {
	if o == nil {
		return 0
	}
	return o.BaseExperience
}

func (o *Pokemon) GetForms() []PokemonForm {
	if o == nil {
		return []PokemonForm{}
	}
	return o.Forms
}

func (o *Pokemon) GetGameIndices() []VersionGameIndex {
	if o == nil {
		return []VersionGameIndex{}
	}
	return o.GameIndices
}

func (o *Pokemon) GetHeight() int {
	if o == nil {
		return 0
	}
	return o.Height
}

func (o *Pokemon) GetHeldItems() []PokemonHeldItem {
	if o == nil {
		return []PokemonHeldItem{}
	}
	return o.HeldItems
}

func (o *Pokemon) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Pokemon) GetIsDefault() bool {
	if o == nil {
		return false
	}
	return o.IsDefault
}

func (o *Pokemon) GetLocationAreaEncounters() string {
	if o == nil {
		return ""
	}
	return o.LocationAreaEncounters
}

func (o *Pokemon) GetMoves() []PokemonMove {
	if o == nil {
		return []PokemonMove{}
	}
	return o.Moves
}

func (o *Pokemon) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Pokemon) GetOrder() int {
	if o == nil {
		return 0
	}
	return o.Order
}

func (o *Pokemon) GetSpecies() NamedAPIResource {
	if o == nil {
		return NamedAPIResource{}
	}
	return o.Species
}

func (o *Pokemon) GetSprites() PokemonSprites {
	if o == nil {
		return PokemonSprites{}
	}
	return o.Sprites
}

func (o *Pokemon) GetStats() []PokemonStat {
	if o == nil {
		return []PokemonStat{}
	}
	return o.Stats
}

func (o *Pokemon) GetTypes() []PokemonType {
	if o == nil {
		return []PokemonType{}
	}
	return o.Types
}

func (o *Pokemon) GetWeight() int {
	if o == nil {
		return 0
	}
	return o.Weight
}
