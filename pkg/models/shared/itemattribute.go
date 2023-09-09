// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ItemAttribute struct {
	// The identifier for this item attribute resource
	ID *int64 `json:"id,omitempty"`
	// A list of items that have this attribute
	Items []NamedAPIResource `json:"items,omitempty"`
	// The name for this item attribute resource
	Name *string `json:"name,omitempty"`
}

func (o *ItemAttribute) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ItemAttribute) GetItems() []NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ItemAttribute) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
