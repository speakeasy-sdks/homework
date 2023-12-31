// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MoveContestEffect struct {
	URL string `json:"url"`
}

func (o *MoveContestEffect) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveContestType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *MoveContestType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveContestType) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveDamageClass struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *MoveDamageClass) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveDamageClass) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveGeneration struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *MoveGeneration) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveGeneration) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveSuperContestEffect struct {
	URL string `json:"url"`
}

func (o *MoveSuperContestEffect) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveTarget struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *MoveTarget) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveTarget) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type MoveType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *MoveType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveType) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

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

func (o *Move) GetAccuracy() *int {
	if o == nil {
		return nil
	}
	return o.Accuracy
}

func (o *Move) GetContestCombos() *ContestComboSets {
	if o == nil {
		return nil
	}
	return o.ContestCombos
}

func (o *Move) GetContestEffect() *MoveContestEffect {
	if o == nil {
		return nil
	}
	return o.ContestEffect
}

func (o *Move) GetContestType() *MoveContestType {
	if o == nil {
		return nil
	}
	return o.ContestType
}

func (o *Move) GetDamageClass() *MoveDamageClass {
	if o == nil {
		return nil
	}
	return o.DamageClass
}

func (o *Move) GetEffectChance() *int {
	if o == nil {
		return nil
	}
	return o.EffectChance
}

func (o *Move) GetEffectChanges() []AbilityEffectChange {
	if o == nil {
		return nil
	}
	return o.EffectChanges
}

func (o *Move) GetEffectEntries() []VerboseEffect {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *Move) GetGeneration() MoveGeneration {
	if o == nil {
		return MoveGeneration{}
	}
	return o.Generation
}

func (o *Move) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Move) GetMeta() *MoveMetaData {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Move) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Move) GetNames() []Name {
	if o == nil {
		return nil
	}
	return o.Names
}

func (o *Move) GetPastValues() []PastMoveStatValues {
	if o == nil {
		return nil
	}
	return o.PastValues
}

func (o *Move) GetPower() *int {
	if o == nil {
		return nil
	}
	return o.Power
}

func (o *Move) GetPp() int {
	if o == nil {
		return 0
	}
	return o.Pp
}

func (o *Move) GetPriority() int {
	if o == nil {
		return 0
	}
	return o.Priority
}

func (o *Move) GetStatChanges() []MoveStatChange {
	if o == nil {
		return nil
	}
	return o.StatChanges
}

func (o *Move) GetSuperContestEffect() *MoveSuperContestEffect {
	if o == nil {
		return nil
	}
	return o.SuperContestEffect
}

func (o *Move) GetTarget() MoveTarget {
	if o == nil {
		return MoveTarget{}
	}
	return o.Target
}

func (o *Move) GetType() MoveType {
	if o == nil {
		return MoveType{}
	}
	return o.Type
}
