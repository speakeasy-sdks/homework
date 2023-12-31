// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIResource struct {
	// The ID of the referenced resource.
	ID *int `json:"id,omitempty"`
	// The name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// The URL of the referenced resource.
	URL *string `json:"url,omitempty"`
}

func (o *APIResource) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *APIResource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *APIResource) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
