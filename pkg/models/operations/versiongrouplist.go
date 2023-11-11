// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type VersionGroupListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *VersionGroupListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *VersionGroupListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// VersionGroupListResponseBody - OK
type VersionGroupListResponseBody struct {
	// The total number of version groups.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of version groups.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of version groups.
	Previous *string               `json:"previous,omitempty"`
	Results  []shared.VersionGroup `json:"results,omitempty"`
}

func (o *VersionGroupListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *VersionGroupListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *VersionGroupListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *VersionGroupListResponseBody) GetResults() []shared.VersionGroup {
	if o == nil {
		return nil
	}
	return o.Results
}

type VersionGroupListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *VersionGroupListResponseBody
}

func (o *VersionGroupListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VersionGroupListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VersionGroupListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VersionGroupListResponse) GetObject() *VersionGroupListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
