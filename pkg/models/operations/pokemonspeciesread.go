// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type PokemonSpeciesReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type PokemonSpeciesReadResponse struct {
	ContentType string
	// Successful response
	PokemonSpecies *shared.PokemonSpecies
	StatusCode     int
	RawResponse    *http.Response
}
