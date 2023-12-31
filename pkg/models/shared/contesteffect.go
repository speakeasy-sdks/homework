// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContestEffect struct {
	// The base number of hearts the user of this move gets
	Appeal int64 `json:"appeal"`
	// The flavor text of this contest effect listed in different languages
	EffectEntries []VerboseEffect `json:"effect_entries"`
	// The flavor text of this contest effect listed in different languages
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	// The identifier for this contest effect resource
	ID int64 `json:"id"`
	// The base number of hearts the user's opponent loses
	Jam int64 `json:"jam"`
}

func (o *ContestEffect) GetAppeal() int64 {
	if o == nil {
		return 0
	}
	return o.Appeal
}

func (o *ContestEffect) GetEffectEntries() []VerboseEffect {
	if o == nil {
		return []VerboseEffect{}
	}
	return o.EffectEntries
}

func (o *ContestEffect) GetFlavorTextEntries() []FlavorText {
	if o == nil {
		return []FlavorText{}
	}
	return o.FlavorTextEntries
}

func (o *ContestEffect) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *ContestEffect) GetJam() int64 {
	if o == nil {
		return 0
	}
	return o.Jam
}
