// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveTargetListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveTargetListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveTargetListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// MoveTargetListResponseBody - OK
type MoveTargetListResponseBody struct {
	// The total number of move targets.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move targets.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move targets.
	Previous *string             `json:"previous,omitempty"`
	Results  []shared.MoveTarget `json:"results,omitempty"`
}

func (o *MoveTargetListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveTargetListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveTargetListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveTargetListResponseBody) GetResults() []shared.MoveTarget {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveTargetListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *MoveTargetListResponseBody
}

func (o *MoveTargetListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveTargetListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveTargetListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveTargetListResponse) GetObject() *MoveTargetListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
