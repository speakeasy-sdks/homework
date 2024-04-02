// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MoveTarget1 - Targets moves can be directed at during battle. Targets can be Pokémon, adjacent positions, all opponents, etc.
type MoveTarget1 struct {
	// The description of this move target listed in different languages
	Descriptions []Description `json:"descriptions"`
	// The identifier for this move target resource
	ID int64 `json:"id"`
	// The name for this move target resource
	Name string `json:"name"`
}

func (o *MoveTarget1) GetDescriptions() []Description {
	if o == nil {
		return []Description{}
	}
	return o.Descriptions
}

func (o *MoveTarget1) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *MoveTarget1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}