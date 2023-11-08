// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ItemFlingEffect struct {
	EffectEntries []VerboseEffect `json:"effect_entries,omitempty"`
	ID            *int            `json:"id,omitempty"`
	Items         []Item          `json:"items,omitempty"`
	Name          *string         `json:"name,omitempty"`
}

func (o *ItemFlingEffect) GetEffectEntries() []VerboseEffect {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *ItemFlingEffect) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ItemFlingEffect) GetItems() []Item {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ItemFlingEffect) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ItemFlingEffectInput struct {
	EffectEntries []VerboseEffect `json:"effect_entries,omitempty"`
	ID            *int            `json:"id,omitempty"`
	Items         []Item          `json:"items,omitempty"`
	Name          *string         `json:"name,omitempty"`
}

func (o *ItemFlingEffectInput) GetEffectEntries() []VerboseEffect {
	if o == nil {
		return nil
	}
	return o.EffectEntries
}

func (o *ItemFlingEffectInput) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ItemFlingEffectInput) GetItems() []Item {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ItemFlingEffectInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
