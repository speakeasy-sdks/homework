// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EffectEntries struct {
	Effect       *VerboseEffect    `json:"effect,omitempty"`
	VersionGroup *NamedAPIResource `json:"version_group,omitempty"`
}

func (o *EffectEntries) GetEffect() *VerboseEffect {
	if o == nil {
		return nil
	}
	return o.Effect
}

func (o *EffectEntries) GetVersionGroup() *NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.VersionGroup
}

type PastMoveStatValues struct {
	// The percent value of how likely this move is to be successful.
	Accuracy int64 `json:"accuracy"`
	// The percent value of effect occurring.
	EffectChance *int64 `json:"effect_chance,omitempty"`
	// The list of previous effects this move has had across version groups.
	EffectEntries []EffectEntries `json:"effect_entries,omitempty"`
	// The base power of this move with a value of 0 if it does not have a base power.
	Power int64 `json:"power"`
	// The power points this move has left.
	Pp   int64 `json:"pp"`
	Type *Type `json:"type,omitempty"`
}

func (o *PastMoveStatValues) GetAccuracy() int64 {
	if o == nil {
		return 0
	}
	return o.Accuracy
}

func (o *PastMoveStatValues) GetEffectChance() *int64 {
	if o == nil {
		return nil
	}
	return o.EffectChance
}

func (o *PastMoveStatValues) GetEffectEntries() []EffectEntries {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *PastMoveStatValues) GetPower() int64 {
	if o == nil {
		return 0
	}
	return o.Power
}

func (o *PastMoveStatValues) GetPp() int64 {
	if o == nil {
		return 0
	}
	return o.Pp
}

func (o *PastMoveStatValues) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
