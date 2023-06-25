// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type StatReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type StatReadResponse struct {
	ContentType string
	// Successful response
	Stat        *shared.Stat
	StatusCode  int
	RawResponse *http.Response
}
