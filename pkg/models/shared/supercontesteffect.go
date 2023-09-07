// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SuperContestEffect - Successful response
type SuperContestEffect struct {
	Appeal            int                `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	ID                int                `json:"id"`
	Moves             []NamedAPIResource `json:"moves"`
}

func (o *SuperContestEffect) GetAppeal() int {
	if o == nil {
		return 0
	}
	return o.Appeal
}

func (o *SuperContestEffect) GetFlavorTextEntries() []FlavorText {
	if o == nil {
		return []FlavorText{}
	}
	return o.FlavorTextEntries
}

func (o *SuperContestEffect) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *SuperContestEffect) GetMoves() []NamedAPIResource {
	if o == nil {
		return []NamedAPIResource{}
	}
	return o.Moves
}
