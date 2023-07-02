// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveLearnMethodListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

// MoveLearnMethodList200ApplicationJSON - OK
type MoveLearnMethodList200ApplicationJSON struct {
	// The total number of move learn methods.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move learn methods.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move learn methods.
	Previous *string                  `json:"previous,omitempty"`
	Results  []shared.MoveLearnMethod `json:"results,omitempty"`
}

type MoveLearnMethodListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveLearnMethodList200ApplicationJSONObject *MoveLearnMethodList200ApplicationJSON
}
