// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type PokemonSpeciesReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PokemonSpeciesReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type PokemonSpeciesReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	PokemonSpecies *shared.PokemonSpeciesInput
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PokemonSpeciesReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokemonSpeciesReadResponse) GetPokemonSpecies() *shared.PokemonSpeciesInput {
	if o == nil {
		return nil
	}
	return o.PokemonSpecies
}

func (o *PokemonSpeciesReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokemonSpeciesReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
