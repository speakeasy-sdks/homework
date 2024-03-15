// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type BerryFirmnessListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *BerryFirmnessListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *BerryFirmnessListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// BerryFirmnessListResponseBody - OK
type BerryFirmnessListResponseBody struct {
	// The total number of berry firmnesses available from this API.
	Count *int64 `json:"count,omitempty"`
	// The URL for the next page of results, or null if there are no more results.
	Next *string `json:"next,omitempty"`
	// The URL for the previous page of results, or null if this is the first page.
	Previous *string                `json:"previous,omitempty"`
	Results  []shared.BerryFirmness `json:"results,omitempty"`
}

func (o *BerryFirmnessListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *BerryFirmnessListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *BerryFirmnessListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *BerryFirmnessListResponseBody) GetResults() []shared.BerryFirmness {
	if o == nil {
		return nil
	}
	return o.Results
}

type BerryFirmnessListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *BerryFirmnessListResponseBody
}

func (o *BerryFirmnessListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BerryFirmnessListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BerryFirmnessListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BerryFirmnessListResponse) GetObject() *BerryFirmnessListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
