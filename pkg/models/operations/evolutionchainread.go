// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EvolutionChainReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type EvolutionChainReadResponse struct {
	ContentType string
	// Successful response
	EvolutionChain *shared.EvolutionChain
	StatusCode     int
	RawResponse    *http.Response
}
