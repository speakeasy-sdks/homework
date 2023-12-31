// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GenerationGameIndex struct {
	// The internal id of an API resource within game data.
	GameIndex  *int        `json:"game_index,omitempty"`
	Generation *Generation `json:"generation,omitempty"`
}

func (o *GenerationGameIndex) GetGameIndex() *int {
	if o == nil {
		return nil
	}
	return o.GameIndex
}

func (o *GenerationGameIndex) GetGeneration() *Generation {
	if o == nil {
		return nil
	}
	return o.Generation
}
