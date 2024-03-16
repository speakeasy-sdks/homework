// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type PokemonShapeReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PokemonShapeReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type PokemonShapeReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	PokemonShape *shared.PokemonShape
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PokemonShapeReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokemonShapeReadResponse) GetPokemonShape() *shared.PokemonShape {
	if o == nil {
		return nil
	}
	return o.PokemonShape
}

func (o *PokemonShapeReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokemonShapeReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
