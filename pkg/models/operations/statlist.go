// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type StatListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *StatListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *StatListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// StatListResponseBody - OK
type StatListResponseBody struct {
	// The total number of stats.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of stats.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of stats.
	Previous *string       `json:"previous,omitempty"`
	Results  []shared.Stat `json:"results,omitempty"`
}

func (o *StatListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *StatListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *StatListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *StatListResponseBody) GetResults() []shared.Stat {
	if o == nil {
		return nil
	}
	return o.Results
}

type StatListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *StatListResponseBody
}

func (o *StatListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StatListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StatListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *StatListResponse) GetObject() *StatListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
