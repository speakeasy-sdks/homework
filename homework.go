// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package homework

import (
	"fmt"
	"homework/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://pokeapi.co/",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type Homework struct {
	Ability                 *ability
	Berry                   *berry
	BerryFirmness           *berryFirmness
	BerryFlavor             *berryFlavor
	Characteristic          *characteristic
	ContestEffect           *contestEffect
	ContestType             *contestType
	EggGroup                *eggGroup
	EncounterCondition      *encounterCondition
	EncounterConditionValue *encounterConditionValue
	EncounterMethod         *encounterMethod
	EvolutionChain          *evolutionChain
	EvolutionTrigger        *evolutionTrigger
	Gender                  *gender
	Generation              *generation
	GrowthRate              *growthRate
	Item                    *item
	ItemAttribute           *itemAttribute
	ItemCategory            *itemCategory
	ItemFlingEffect         *itemFlingEffect
	ItemPocket              *itemPocket
	Language                *language
	Location                *location
	LocationArea            *locationArea
	Machine                 *machine
	Move                    *move
	MoveAilment             *moveAilment
	MoveBattleStyle         *moveBattleStyle
	MoveCategory            *moveCategory
	MoveDamageClass         *moveDamageClass
	MoveLearnMethod         *moveLearnMethod
	MoveTarget              *moveTarget
	Nature                  *nature
	PalParkArea             *palParkArea
	PokeathlonStat          *pokeathlonStat
	Pokedex                 *pokedex
	Pokemon                 *pokemon
	PokemonColor            *pokemonColor
	PokemonForm             *pokemonForm
	PokemonHabitat          *pokemonHabitat
	PokemonShape            *pokemonShape
	PokemonSpecies          *pokemonSpecies
	Region                  *region
	Stat                    *stat
	SuperContestEffect      *superContestEffect
	Type                    *typeT
	Version                 *version
	VersionGroup            *versionGroup

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Homework)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Homework) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Homework) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Homework) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Homework) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Homework) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Homework {
	sdk := &Homework{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "20220523",
			SDKVersion:        "0.2.0",
			GenVersion:        "2.150.1",
			UserAgent:         "speakeasy-sdk/go 0.2.0 2.150.1 20220523 homework",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.Ability = newAbility(sdk.sdkConfiguration)

	sdk.Berry = newBerry(sdk.sdkConfiguration)

	sdk.BerryFirmness = newBerryFirmness(sdk.sdkConfiguration)

	sdk.BerryFlavor = newBerryFlavor(sdk.sdkConfiguration)

	sdk.Characteristic = newCharacteristic(sdk.sdkConfiguration)

	sdk.ContestEffect = newContestEffect(sdk.sdkConfiguration)

	sdk.ContestType = newContestType(sdk.sdkConfiguration)

	sdk.EggGroup = newEggGroup(sdk.sdkConfiguration)

	sdk.EncounterCondition = newEncounterCondition(sdk.sdkConfiguration)

	sdk.EncounterConditionValue = newEncounterConditionValue(sdk.sdkConfiguration)

	sdk.EncounterMethod = newEncounterMethod(sdk.sdkConfiguration)

	sdk.EvolutionChain = newEvolutionChain(sdk.sdkConfiguration)

	sdk.EvolutionTrigger = newEvolutionTrigger(sdk.sdkConfiguration)

	sdk.Gender = newGender(sdk.sdkConfiguration)

	sdk.Generation = newGeneration(sdk.sdkConfiguration)

	sdk.GrowthRate = newGrowthRate(sdk.sdkConfiguration)

	sdk.Item = newItem(sdk.sdkConfiguration)

	sdk.ItemAttribute = newItemAttribute(sdk.sdkConfiguration)

	sdk.ItemCategory = newItemCategory(sdk.sdkConfiguration)

	sdk.ItemFlingEffect = newItemFlingEffect(sdk.sdkConfiguration)

	sdk.ItemPocket = newItemPocket(sdk.sdkConfiguration)

	sdk.Language = newLanguage(sdk.sdkConfiguration)

	sdk.Location = newLocation(sdk.sdkConfiguration)

	sdk.LocationArea = newLocationArea(sdk.sdkConfiguration)

	sdk.Machine = newMachine(sdk.sdkConfiguration)

	sdk.Move = newMove(sdk.sdkConfiguration)

	sdk.MoveAilment = newMoveAilment(sdk.sdkConfiguration)

	sdk.MoveBattleStyle = newMoveBattleStyle(sdk.sdkConfiguration)

	sdk.MoveCategory = newMoveCategory(sdk.sdkConfiguration)

	sdk.MoveDamageClass = newMoveDamageClass(sdk.sdkConfiguration)

	sdk.MoveLearnMethod = newMoveLearnMethod(sdk.sdkConfiguration)

	sdk.MoveTarget = newMoveTarget(sdk.sdkConfiguration)

	sdk.Nature = newNature(sdk.sdkConfiguration)

	sdk.PalParkArea = newPalParkArea(sdk.sdkConfiguration)

	sdk.PokeathlonStat = newPokeathlonStat(sdk.sdkConfiguration)

	sdk.Pokedex = newPokedex(sdk.sdkConfiguration)

	sdk.Pokemon = newPokemon(sdk.sdkConfiguration)

	sdk.PokemonColor = newPokemonColor(sdk.sdkConfiguration)

	sdk.PokemonForm = newPokemonForm(sdk.sdkConfiguration)

	sdk.PokemonHabitat = newPokemonHabitat(sdk.sdkConfiguration)

	sdk.PokemonShape = newPokemonShape(sdk.sdkConfiguration)

	sdk.PokemonSpecies = newPokemonSpecies(sdk.sdkConfiguration)

	sdk.Region = newRegion(sdk.sdkConfiguration)

	sdk.Stat = newStat(sdk.sdkConfiguration)

	sdk.SuperContestEffect = newSuperContestEffect(sdk.sdkConfiguration)

	sdk.Type = newType(sdk.sdkConfiguration)

	sdk.Version = newVersion(sdk.sdkConfiguration)

	sdk.VersionGroup = newVersionGroup(sdk.sdkConfiguration)

	return sdk
}
