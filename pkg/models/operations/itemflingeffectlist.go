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

func (o *ItemFlingEffectListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ItemFlingEffectListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// ItemFlingEffectListResponseBody - OK
type ItemFlingEffectListResponseBody struct {
	// The total number of item fling effects.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of item fling effects.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of item fling effects.
	Previous *string                       `json:"previous,omitempty"`
	Results  []shared.ItemFlingEffectInput `json:"results,omitempty"`
}

func (o *ItemFlingEffectListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ItemFlingEffectListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ItemFlingEffectListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ItemFlingEffectListResponseBody) GetResults() []shared.ItemFlingEffectInput {
	if o == nil {
		return nil
	}
	return o.Results
}

type ItemFlingEffectListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ItemFlingEffectListResponseBody
}

func (o *ItemFlingEffectListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemFlingEffectListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemFlingEffectListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ItemFlingEffectListResponse) GetObject() *ItemFlingEffectListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
