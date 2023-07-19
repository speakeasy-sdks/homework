// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Version - Successful response
type Version struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	Names        []Name           `json:"names"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

func (o *Version) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Version) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Version) GetNames() []Name {
	if o == nil {
		return []Name{}
	}
	return o.Names
}

func (o *Version) GetVersionGroup() NamedAPIResource {
	if o == nil {
		return NamedAPIResource{}
	}
	return o.VersionGroup
}
