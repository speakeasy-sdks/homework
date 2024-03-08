// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Characteristic struct {
	// The remainder of the highest stat/IV divided by 5
	GeneModulo int64 `json:"gene_modulo"`
	// The identifier for this characteristic resource
	ID int64 `json:"id"`
	// The possible values of the highest stat that would result in a Pokémon recieving this characteristic when divided by 5
	PossibleValues []int64 `json:"possible_values"`
}

func (o *Characteristic) GetGeneModulo() int64 {
	if o == nil {
		return 0
	}
	return o.GeneModulo
}

func (o *Characteristic) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Characteristic) GetPossibleValues() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.PossibleValues
}
