// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContestComboDetail struct {
	// A list of moves to use after this move.
	UseAfter []Move `json:"use_after,omitempty"`
	// A list of moves to use before this move.
	UseBefore []Move `json:"use_before,omitempty"`
}

func (o *ContestComboDetail) GetUseAfter() []Move {
	if o == nil {
		return nil
	}
	return o.UseAfter
}

func (o *ContestComboDetail) GetUseBefore() []Move {
	if o == nil {
		return nil
	}
	return o.UseBefore
}
