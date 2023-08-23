// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Nature - Successful response
type Nature struct {
	DecreasedStat              string                      `json:"decreased_stat"`
	HatesFlavor                string                      `json:"hates_flavor"`
	ID                         int64                       `json:"id"`
	IncreasedStat              string                      `json:"increased_stat"`
	LikesFlavor                string                      `json:"likes_flavor"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences,omitempty"`
	Name                       string                      `json:"name"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes,omitempty"`
}

func (o *Nature) GetDecreasedStat() string {
	if o == nil {
		return ""
	}
	return o.DecreasedStat
}

func (o *Nature) GetHatesFlavor() string {
	if o == nil {
		return ""
	}
	return o.HatesFlavor
}

func (o *Nature) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Nature) GetIncreasedStat() string {
	if o == nil {
		return ""
	}
	return o.IncreasedStat
}

func (o *Nature) GetLikesFlavor() string {
	if o == nil {
		return ""
	}
	return o.LikesFlavor
}

func (o *Nature) GetMoveBattleStylePreferences() []MoveBattleStylePreference {
	if o == nil {
		return nil
	}
	return o.MoveBattleStylePreferences
}

func (o *Nature) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Nature) GetPokeathlonStatChanges() []NatureStatChange {
	if o == nil {
		return nil
	}
	return o.PokeathlonStatChanges
}
