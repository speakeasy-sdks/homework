// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContestCombos - A detail of combos this move can be used in
type ContestCombos struct {
}

type EffectChanges struct {
}

// Effect - An effect that occurs in a game, e.g. causing a Pokémon to fall asleep.
type Effect struct {
	// A detail of combos this move can be used in
	ContestCombos *ContestCombos `json:"contest_combos,omitempty"`
	ContestType   *ContestType   `json:"contest_type,omitempty"`
	// The chance of this move having an additional effect listed in percentage
	EffectChance *int64 `json:"effect_chance,omitempty"`
	// The list of effects that are changed by this ability
	EffectChanges []EffectChanges `json:"effect_changes,omitempty"`
	// The list of effect text entries
	EffectEntries []EffectEffect `json:"effect_entries,omitempty"`
	// The flavor text entries that describe this effect
	FlavorTextEntries []FlavorText `json:"flavor_text_entries,omitempty"`
	Generation        *Generation  `json:"generation,omitempty"`
	// The identifier for this effect resource
	ID *int64 `json:"id,omitempty"`
	// The machines that teach this move
	Machines []MachineVersionDetail `json:"machines,omitempty"`
	Meta     *MoveMetaData          `json:"meta,omitempty"`
	// The name for this effect resource
	Name *string `json:"name,omitempty"`
	// The flavor text entries that describe this effect
	PokemonFlavorTextEntries []EffectEffect `json:"pokemon_flavor_text_entries,omitempty"`
	// The short description of this effect listed in different languages
	ShortEffect *string `json:"short_effect,omitempty"`
	// The list of stat changes that are caused by this effect
	StatChanges        []MoveStatChange     `json:"stat_changes,omitempty"`
	SuperContestEffect *SuperContestEffect  `json:"super_contest_effect,omitempty"`
	TargetSpecies      *PokemonSpeciesInput `json:"target_species,omitempty"`
}

func (o *Effect) GetContestCombos() *ContestCombos {
	if o == nil {
		return nil
	}
	return o.ContestCombos
}

func (o *Effect) GetContestType() *ContestType {
	if o == nil {
		return nil
	}
	return o.ContestType
}

func (o *Effect) GetEffectChance() *int64 {
	if o == nil {
		return nil
	}
	return o.EffectChance
}

func (o *Effect) GetEffectChanges() []EffectChanges {
	if o == nil {
		return nil
	}
	return o.EffectChanges
}

func (o *Effect) GetEffectEntries() []EffectEffect {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *Effect) GetFlavorTextEntries() []FlavorText {
	if o == nil {
		return nil
	}
	return o.FlavorTextEntries
}

func (o *Effect) GetGeneration() *Generation {
	if o == nil {
		return nil
	}
	return o.Generation
}

func (o *Effect) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Effect) GetMachines() []MachineVersionDetail {
	if o == nil {
		return nil
	}
	return o.Machines
}

func (o *Effect) GetMeta() *MoveMetaData {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Effect) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Effect) GetPokemonFlavorTextEntries() []EffectEffect {
	if o == nil {
		return nil
	}
	return o.PokemonFlavorTextEntries
}

func (o *Effect) GetShortEffect() *string {
	if o == nil {
		return nil
	}
	return o.ShortEffect
}

func (o *Effect) GetStatChanges() []MoveStatChange {
	if o == nil {
		return nil
	}
	return o.StatChanges
}

func (o *Effect) GetSuperContestEffect() *SuperContestEffect {
	if o == nil {
		return nil
	}
	return o.SuperContestEffect
}

func (o *Effect) GetTargetSpecies() *PokemonSpeciesInput {
	if o == nil {
		return nil
	}
	return o.TargetSpecies
}
