// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type GenderListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *GenderListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GenderListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// GenderList200ApplicationJSON - OK
type GenderList200ApplicationJSON struct {
	// The total number of genders.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of genders.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of genders.
	Previous *string         `json:"previous,omitempty"`
	Results  []shared.Gender `json:"results,omitempty"`
}

func (o *GenderList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GenderList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GenderList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *GenderList200ApplicationJSON) GetResults() []shared.Gender {
	if o == nil {
		return nil
	}
	return o.Results
}

type GenderListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GenderList200ApplicationJSONObject *GenderList200ApplicationJSON
}

func (o *GenderListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenderListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenderListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GenderListResponse) GetGenderList200ApplicationJSONObject() *GenderList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GenderList200ApplicationJSONObject
}
