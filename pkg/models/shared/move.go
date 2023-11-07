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

type DamageClass struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *DamageClass) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DamageClass) GetURL() string {
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

type Target struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o *Target) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Target) GetURL() string {
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

type MoveInput struct {
	// The percent value of how likely this move is to be successful
	Accuracy      *int               `json:"accuracy,omitempty"`
	ContestCombos *ContestComboSets  `json:"contest_combos,omitempty"`
	ContestEffect *MoveContestEffect `json:"contest_effect,omitempty"`
	ContestType   *MoveContestType   `json:"contest_type,omitempty"`
	DamageClass   *DamageClass       `json:"damage_class,omitempty"`
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
	Target             Target                  `json:"target"`
	Type               MoveType                `json:"type"`
}

func (o *MoveInput) GetAccuracy() *int {
	if o == nil {
		return nil
	}
	return o.Accuracy
}

func (o *MoveInput) GetContestCombos() *ContestComboSets {
	if o == nil {
		return nil
	}
	return o.ContestCombos
}

func (o *MoveInput) GetContestEffect() *MoveContestEffect {
	if o == nil {
		return nil
	}
	return o.ContestEffect
}

func (o *MoveInput) GetContestType() *MoveContestType {
	if o == nil {
		return nil
	}
	return o.ContestType
}

func (o *MoveInput) GetDamageClass() *DamageClass {
	if o == nil {
		return nil
	}
	return o.DamageClass
}

func (o *MoveInput) GetEffectChance() *int {
	if o == nil {
		return nil
	}
	return o.EffectChance
}

func (o *MoveInput) GetEffectChanges() []AbilityEffectChange {
	if o == nil {
		return nil
	}
	return o.EffectChanges
}

func (o *MoveInput) GetEffectEntries() []VerboseEffect {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *MoveInput) GetGeneration() MoveGeneration {
	if o == nil {
		return MoveGeneration{}
	}
	return o.Generation
}

func (o *MoveInput) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *MoveInput) GetMeta() *MoveMetaData {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *MoveInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveInput) GetNames() []Name {
	if o == nil {
		return nil
	}
	return o.Names
}

func (o *MoveInput) GetPastValues() []PastMoveStatValues {
	if o == nil {
		return nil
	}
	return o.PastValues
}

func (o *MoveInput) GetPower() *int {
	if o == nil {
		return nil
	}
	return o.Power
}

func (o *MoveInput) GetPp() int {
	if o == nil {
		return 0
	}
	return o.Pp
}

func (o *MoveInput) GetPriority() int {
	if o == nil {
		return 0
	}
	return o.Priority
}

func (o *MoveInput) GetStatChanges() []MoveStatChange {
	if o == nil {
		return nil
	}
	return o.StatChanges
}

func (o *MoveInput) GetSuperContestEffect() *MoveSuperContestEffect {
	if o == nil {
		return nil
	}
	return o.SuperContestEffect
}

func (o *MoveInput) GetTarget() Target {
	if o == nil {
		return Target{}
	}
	return o.Target
}

func (o *MoveInput) GetType() MoveType {
	if o == nil {
		return MoveType{}
	}
	return o.Type
}
