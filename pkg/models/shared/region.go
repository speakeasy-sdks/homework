// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Region - Successful response
type Region struct {
	ID             int                `json:"id"`
	Locations      []NamedAPIResource `json:"locations"`
	MainGeneration NamedAPIResource   `json:"main_generation"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	Pokedexes      []NamedAPIResource `json:"pokedexes"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

func (o *Region) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Region) GetLocations() []NamedAPIResource {
	if o == nil {
		return []NamedAPIResource{}
	}
	return o.Locations
}

func (o *Region) GetMainGeneration() NamedAPIResource {
	if o == nil {
		return NamedAPIResource{}
	}
	return o.MainGeneration
}

func (o *Region) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Region) GetNames() []Name {
	if o == nil {
		return []Name{}
	}
	return o.Names
}

func (o *Region) GetPokedexes() []NamedAPIResource {
	if o == nil {
		return []NamedAPIResource{}
	}
	return o.Pokedexes
}

func (o *Region) GetVersionGroups() []NamedAPIResource {
	if o == nil {
		return []NamedAPIResource{}
	}
	return o.VersionGroups
}
