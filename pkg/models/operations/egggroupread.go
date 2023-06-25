// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EggGroupReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type EggGroupReadResponse struct {
	ContentType string
	// Successful response
	EggGroup    *shared.EggGroup
	StatusCode  int
	RawResponse *http.Response
}
