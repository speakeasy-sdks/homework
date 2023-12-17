// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type PokemonHabitatReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PokemonHabitatReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type PokemonHabitatReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	PokemonHabitat *shared.PokemonHabitat
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PokemonHabitatReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokemonHabitatReadResponse) GetPokemonHabitat() *shared.PokemonHabitat {
	if o == nil {
		return nil
	}
	return o.PokemonHabitat
}

func (o *PokemonHabitatReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokemonHabitatReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
