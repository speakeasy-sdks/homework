// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ItemCategory - Successful response
type ItemCategory struct {
	// The identifier for this item category resource
	ID *int64 `json:"id,omitempty"`
	// A list of items that are a part of this category
	Items []NamedAPIResource `json:"items,omitempty"`
	// The name for this item category resource
	Name *string `json:"name,omitempty"`
}

func (o *ItemCategory) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ItemCategory) GetItems() []NamedAPIResource {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ItemCategory) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
