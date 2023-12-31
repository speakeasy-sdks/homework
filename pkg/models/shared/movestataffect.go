// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MoveStatAffect struct {
	// The amount of change to the referenced stat.
	Change int64            `json:"change"`
	Move   NamedAPIResource `json:"move"`
}

func (o *MoveStatAffect) GetChange() int64 {
	if o == nil {
		return 0
	}
	return o.Change
}

func (o *MoveStatAffect) GetMove() NamedAPIResource {
	if o == nil {
		return NamedAPIResource{}
	}
	return o.Move
}
