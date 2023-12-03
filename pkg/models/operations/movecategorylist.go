// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveCategoryListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveCategoryListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveCategoryListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// MoveCategoryListResponseBody - OK
type MoveCategoryListResponseBody struct {
	// The total number of move categories.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move categories.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move categories.
	Previous *string               `json:"previous,omitempty"`
	Results  []shared.MoveCategory `json:"results,omitempty"`
}

func (o *MoveCategoryListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveCategoryListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveCategoryListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveCategoryListResponseBody) GetResults() []shared.MoveCategory {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveCategoryListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *MoveCategoryListResponseBody
}

func (o *MoveCategoryListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveCategoryListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveCategoryListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveCategoryListResponse) GetObject() *MoveCategoryListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
