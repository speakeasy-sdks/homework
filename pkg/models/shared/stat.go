// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Stat struct {
	AffectingMoves   *MoveStatAffectSets   `json:"affecting_moves,omitempty"`
	AffectingNatures *NatureStatAffectSets `json:"affecting_natures,omitempty"`
	GameIndex        int                   `json:"game_index"`
	ID               int                   `json:"id"`
	IsBattleOnly     *bool                 `json:"is_battle_only,omitempty"`
	Name             string                `json:"name"`
}

func (o *Stat) GetAffectingMoves() *MoveStatAffectSets {
	if o == nil {
		return nil
	}
	return o.AffectingMoves
}

func (o *Stat) GetAffectingNatures() *NatureStatAffectSets {
	if o == nil {
		return nil
	}
	return o.AffectingNatures
}

func (o *Stat) GetGameIndex() int {
	if o == nil {
		return 0
	}
	return o.GameIndex
}

func (o *Stat) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Stat) GetIsBattleOnly() *bool {
	if o == nil {
		return nil
	}
	return o.IsBattleOnly
}

func (o *Stat) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
