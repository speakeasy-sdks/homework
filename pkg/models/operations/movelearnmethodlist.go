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

func (o *MoveLearnMethodListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveLearnMethodListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
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

func (o *MoveLearnMethodList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveLearnMethodList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveLearnMethodList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveLearnMethodList200ApplicationJSON) GetResults() []shared.MoveLearnMethod {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveLearnMethodListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveLearnMethodList200ApplicationJSONObject *MoveLearnMethodList200ApplicationJSON
}

func (o *MoveLearnMethodListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveLearnMethodListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveLearnMethodListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveLearnMethodListResponse) GetMoveLearnMethodList200ApplicationJSONObject() *MoveLearnMethodList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MoveLearnMethodList200ApplicationJSONObject
}
