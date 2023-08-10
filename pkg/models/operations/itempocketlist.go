// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemPocketListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *ItemPocketListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ItemPocketListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// ItemPocketList200ApplicationJSON - OK
type ItemPocketList200ApplicationJSON struct {
	// The total number of item pockets.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of item pockets.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of item pockets.
	Previous *string             `json:"previous,omitempty"`
	Results  []shared.ItemPocket `json:"results,omitempty"`
}

func (o *ItemPocketList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ItemPocketList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ItemPocketList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ItemPocketList200ApplicationJSON) GetResults() []shared.ItemPocket {
	if o == nil {
		return nil
	}
	return o.Results
}

type ItemPocketListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ItemPocketList200ApplicationJSONObject *ItemPocketList200ApplicationJSON
}

func (o *ItemPocketListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemPocketListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemPocketListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ItemPocketListResponse) GetItemPocketList200ApplicationJSONObject() *ItemPocketList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ItemPocketList200ApplicationJSONObject
}
