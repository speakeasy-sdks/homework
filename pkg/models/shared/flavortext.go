// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type FlavorText struct {
	FlavorText *string           `json:"flavor_text,omitempty"`
	Language   *NamedAPIResource `json:"language,omitempty"`
	Version    *NamedAPIResource `json:"version,omitempty"`
}

func (o *FlavorText) GetFlavorText() *string {
	if o == nil {
		return nil
	}
	return o.FlavorText
}

func (o *FlavorText) GetLanguage() *NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *FlavorText) GetVersion() *NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.Version
}
