// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type VerboseEffect struct {
	// The localized effect text for an API resource in a specific language.
	Effect *string `json:"effect,omitempty"`
}

func (o *VerboseEffect) GetEffect() *string {
	if o == nil {
		return nil
	}
	return o.Effect
}
