// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// MoveList200ApplicationJSON - OK
type MoveList200ApplicationJSON struct {
	// The total number of moves.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of moves.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of moves.
	Previous *string       `json:"previous,omitempty"`
	Results  []shared.Move `json:"results,omitempty"`
}

func (o *MoveList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveList200ApplicationJSON) GetResults() []shared.Move {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveList200ApplicationJSONObject *MoveList200ApplicationJSON
}

func (o *MoveListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveListResponse) GetMoveList200ApplicationJSONObject() *MoveList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MoveList200ApplicationJSONObject
}
