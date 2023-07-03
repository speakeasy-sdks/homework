// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type GrowthRateReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type GrowthRateReadResponse struct {
	ContentType string
	// Successful response
	GrowthRate  *shared.GrowthRate
	StatusCode  int
	RawResponse *http.Response
}
