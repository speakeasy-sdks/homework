// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PokeathlonStatAffectingNatures - A detail of natures which affect this Pokéathlon stat positively or negatively
type PokeathlonStatAffectingNatures struct {
	// A list of natures that negatively affect this Pokéathlon stat
	Decrease []NaturePokeathlonStatAffectSets `json:"decrease,omitempty"`
	// A list of natures that positively affect this Pokéathlon stat
	Increase []NaturePokeathlonStatAffectSets `json:"increase,omitempty"`
}

func (o *PokeathlonStatAffectingNatures) GetDecrease() []NaturePokeathlonStatAffectSets {
	if o == nil {
		return nil
	}
	return o.Decrease
}

func (o *PokeathlonStatAffectingNatures) GetIncrease() []NaturePokeathlonStatAffectSets {
	if o == nil {
		return nil
	}
	return o.Increase
}

type PokeathlonStat struct {
	// A detail of natures which affect this Pokéathlon stat positively or negatively
	AffectingNatures PokeathlonStatAffectingNatures `json:"affecting_natures"`
	// The identifier for this Pokéathlon stat resource
	ID int64 `json:"id"`
	// The name for this Pokéathlon stat resource
	Name string `json:"name"`
	// The name of this Pokéathlon stat listed in different languages
	Names []Name `json:"names"`
}

func (o *PokeathlonStat) GetAffectingNatures() PokeathlonStatAffectingNatures {
	if o == nil {
		return PokeathlonStatAffectingNatures{}
	}
	return o.AffectingNatures
}

func (o *PokeathlonStat) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PokeathlonStat) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PokeathlonStat) GetNames() []Name {
	if o == nil {
		return []Name{}
	}
	return o.Names
}
