// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TypeDamageRelationsDoubleDamageFrom struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsDoubleDamageFrom) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsDoubleDamageFrom) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelationsDoubleDamageTo struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsDoubleDamageTo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsDoubleDamageTo) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelationsHalfDamageFrom struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsHalfDamageFrom) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsHalfDamageFrom) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelationsHalfDamageTo struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsHalfDamageTo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsHalfDamageTo) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelationsNoDamageFrom struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsNoDamageFrom) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsNoDamageFrom) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelationsNoDamageTo struct {
	// The name of the related type.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the related type.
	URL *string `json:"url,omitempty"`
}

func (o *TypeDamageRelationsNoDamageTo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeDamageRelationsNoDamageTo) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeDamageRelations struct {
	DoubleDamageFrom []TypeDamageRelationsDoubleDamageFrom `json:"double_damage_from,omitempty"`
	DoubleDamageTo   []TypeDamageRelationsDoubleDamageTo   `json:"double_damage_to,omitempty"`
	HalfDamageFrom   []TypeDamageRelationsHalfDamageFrom   `json:"half_damage_from,omitempty"`
	HalfDamageTo     []TypeDamageRelationsHalfDamageTo     `json:"half_damage_to,omitempty"`
	NoDamageFrom     []TypeDamageRelationsNoDamageFrom     `json:"no_damage_from,omitempty"`
	NoDamageTo       []TypeDamageRelationsNoDamageTo       `json:"no_damage_to,omitempty"`
}

func (o *TypeDamageRelations) GetDoubleDamageFrom() []TypeDamageRelationsDoubleDamageFrom {
	if o == nil {
		return nil
	}
	return o.DoubleDamageFrom
}

func (o *TypeDamageRelations) GetDoubleDamageTo() []TypeDamageRelationsDoubleDamageTo {
	if o == nil {
		return nil
	}
	return o.DoubleDamageTo
}

func (o *TypeDamageRelations) GetHalfDamageFrom() []TypeDamageRelationsHalfDamageFrom {
	if o == nil {
		return nil
	}
	return o.HalfDamageFrom
}

func (o *TypeDamageRelations) GetHalfDamageTo() []TypeDamageRelationsHalfDamageTo {
	if o == nil {
		return nil
	}
	return o.HalfDamageTo
}

func (o *TypeDamageRelations) GetNoDamageFrom() []TypeDamageRelationsNoDamageFrom {
	if o == nil {
		return nil
	}
	return o.NoDamageFrom
}

func (o *TypeDamageRelations) GetNoDamageTo() []TypeDamageRelationsNoDamageTo {
	if o == nil {
		return nil
	}
	return o.NoDamageTo
}

type TypeGameIndicesGeneration struct {
	// The generation this game index is related to.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the generation.
	URL *string `json:"url,omitempty"`
}

func (o *TypeGameIndicesGeneration) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeGameIndicesGeneration) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeGameIndices struct {
	// The internal id of an api resource within game data.
	GameIndex  *int                       `json:"game_index,omitempty"`
	Generation *TypeGameIndicesGeneration `json:"generation,omitempty"`
}

func (o *TypeGameIndices) GetGameIndex() *int {
	if o == nil {
		return nil
	}
	return o.GameIndex
}

func (o *TypeGameIndices) GetGeneration() *TypeGameIndicesGeneration {
	if o == nil {
		return nil
	}
	return o.Generation
}

type TypeGeneration struct {
	// The generation this type originated in.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for this generation.
	URL *string `json:"url,omitempty"`
}

func (o *TypeGeneration) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeGeneration) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeMoveDamageClass struct {
	// The name of this move damage class.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for this move damage class.
	URL *string `json:"url,omitempty"`
}

func (o *TypeMoveDamageClass) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeMoveDamageClass) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeNamesLanguage struct {
	// The name of the language.
	Name *string `json:"name,omitempty"`
	// The URL of the API endpoint for the language.
	URL *string `json:"url,omitempty"`
}

func (o *TypeNamesLanguage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TypeNamesLanguage) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type TypeNames struct {
	Language *TypeNamesLanguage `json:"language,omitempty"`
	// The localized name for an API resource in a specific language.
	Name *string `json:"name,omitempty"`
}

func (o *TypeNames) GetLanguage() *TypeNamesLanguage {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *TypeNames) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TypePokemon struct {
	Slot *int64 `json:"slot,omitempty"`
}

func (o *TypePokemon) GetSlot() *int64 {
	if o == nil {
		return nil
	}
	return o.Slot
}

// Type - Successful response
type Type struct {
	DamageRelations *TypeDamageRelations `json:"damage_relations,omitempty"`
	GameIndices     []TypeGameIndices    `json:"game_indices,omitempty"`
	Generation      *TypeGeneration      `json:"generation,omitempty"`
	// The identifier for this type resource.
	ID              *int                 `json:"id,omitempty"`
	MoveDamageClass *TypeMoveDamageClass `json:"move_damage_class,omitempty"`
	// The name for this type resource.
	Name    *string       `json:"name,omitempty"`
	Names   []TypeNames   `json:"names,omitempty"`
	Pokemon []TypePokemon `json:"pokemon,omitempty"`
}

func (o *Type) GetDamageRelations() *TypeDamageRelations {
	if o == nil {
		return nil
	}
	return o.DamageRelations
}

func (o *Type) GetGameIndices() []TypeGameIndices {
	if o == nil {
		return nil
	}
	return o.GameIndices
}

func (o *Type) GetGeneration() *TypeGeneration {
	if o == nil {
		return nil
	}
	return o.Generation
}

func (o *Type) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Type) GetMoveDamageClass() *TypeMoveDamageClass {
	if o == nil {
		return nil
	}
	return o.MoveDamageClass
}

func (o *Type) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Type) GetNames() []TypeNames {
	if o == nil {
		return nil
	}
	return o.Names
}

func (o *Type) GetPokemon() []TypePokemon {
	if o == nil {
		return nil
	}
	return o.Pokemon
}
