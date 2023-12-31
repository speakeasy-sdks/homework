// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveAilmentListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveAilmentListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveAilmentListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// MoveAilmentList200ApplicationJSON - OK
type MoveAilmentList200ApplicationJSON struct {
	// The total number of move ailments.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move ailments.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move ailments.
	Previous *string              `json:"previous,omitempty"`
	Results  []shared.MoveAilment `json:"results,omitempty"`
}

func (o *MoveAilmentList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveAilmentList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveAilmentList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveAilmentList200ApplicationJSON) GetResults() []shared.MoveAilment {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveAilmentListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveAilmentList200ApplicationJSONObject *MoveAilmentList200ApplicationJSON
}

func (o *MoveAilmentListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveAilmentListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveAilmentListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveAilmentListResponse) GetMoveAilmentList200ApplicationJSONObject() *MoveAilmentList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MoveAilmentList200ApplicationJSONObject
}
