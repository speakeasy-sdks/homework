// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContestName struct {
	// The localized name for an API resource in a specific language.
	Name *string `json:"name,omitempty"`
}

func (o *ContestName) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
