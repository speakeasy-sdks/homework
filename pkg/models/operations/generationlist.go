// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type GenerationListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *GenerationListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GenerationListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// GenerationList200ApplicationJSON - OK
type GenerationList200ApplicationJSON struct {
	// The total number of generations.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of generations.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of generations.
	Previous *string             `json:"previous,omitempty"`
	Results  []shared.Generation `json:"results,omitempty"`
}

func (o *GenerationList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GenerationList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GenerationList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *GenerationList200ApplicationJSON) GetResults() []shared.Generation {
	if o == nil {
		return nil
	}
	return o.Results
}

type GenerationListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GenerationList200ApplicationJSONObject *GenerationList200ApplicationJSON
}

func (o *GenerationListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerationListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerationListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GenerationListResponse) GetGenerationList200ApplicationJSONObject() *GenerationList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GenerationList200ApplicationJSONObject
}
