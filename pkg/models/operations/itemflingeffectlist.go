// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemFlingEffectListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

// ItemFlingEffectList200ApplicationJSON - OK
type ItemFlingEffectList200ApplicationJSON struct {
	// The total number of item fling effects.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of item fling effects.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of item fling effects.
	Previous *string                  `json:"previous,omitempty"`
	Results  []shared.ItemFlingEffect `json:"results,omitempty"`
}

type ItemFlingEffectListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ItemFlingEffectList200ApplicationJSONObject *ItemFlingEffectList200ApplicationJSON
}
