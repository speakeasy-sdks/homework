// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MoveLearnMethod - Methods by which Pokémon can learn moves.
type MoveLearnMethod struct {
	// The description of this move learn method listed in different languages.
	Descriptions []Description `json:"descriptions,omitempty"`
	// The identifier for this move learn method resource.
	ID int64 `json:"id"`
	// The name for this move learn method resource.
	Name string `json:"name"`
	// The name of this move learn method listed in different languages.
	Names []Name `json:"names,omitempty"`
}

func (o *MoveLearnMethod) GetDescriptions() []Description {
	if o == nil {
		return nil
	}
	return o.Descriptions
}

func (o *MoveLearnMethod) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *MoveLearnMethod) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoveLearnMethod) GetNames() []Name {
	if o == nil {
		return nil
	}
	return o.Names
}
