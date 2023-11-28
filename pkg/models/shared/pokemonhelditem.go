// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type VersionDetails struct {
	Rarity  *int64   `json:"rarity,omitempty"`
	Version *Version `json:"version,omitempty"`
}

func (o *VersionDetails) GetRarity() *int64 {
	if o == nil {
		return nil
	}
	return o.Rarity
}

func (o *VersionDetails) GetVersion() *Version {
	if o == nil {
		return nil
	}
	return o.Version
}

type PokemonHeldItem struct {
	Item           *ItemInput       `json:"item,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

func (o *PokemonHeldItem) GetItem() *ItemInput {
	if o == nil {
		return nil
	}
	return o.Item
}

func (o *PokemonHeldItem) GetVersionDetails() []VersionDetails {
	if o == nil {
		return nil
	}
	return o.VersionDetails
}
