// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ItemPocket struct {
	Categories []ItemCategory `json:"categories,omitempty"`
	ID         *int           `json:"id,omitempty"`
	Name       *string        `json:"name,omitempty"`
}

func (o *ItemPocket) GetCategories() []ItemCategory {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *ItemPocket) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ItemPocket) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
