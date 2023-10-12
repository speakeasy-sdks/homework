// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type SuperContestEffectListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *SuperContestEffectListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SuperContestEffectListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// SuperContestEffectList200ApplicationJSON - OK
type SuperContestEffectList200ApplicationJSON struct {
	// The total number of super contest effects.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of super contest effects.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of super contest effects.
	Previous *string                     `json:"previous,omitempty"`
	Results  []shared.SuperContestEffect `json:"results,omitempty"`
}

func (o *SuperContestEffectList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *SuperContestEffectList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *SuperContestEffectList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *SuperContestEffectList200ApplicationJSON) GetResults() []shared.SuperContestEffect {
	if o == nil {
		return nil
	}
	return o.Results
}

type SuperContestEffectListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SuperContestEffectList200ApplicationJSONObject *SuperContestEffectList200ApplicationJSON
}

func (o *SuperContestEffectListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SuperContestEffectListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SuperContestEffectListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SuperContestEffectListResponse) GetSuperContestEffectList200ApplicationJSONObject() *SuperContestEffectList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SuperContestEffectList200ApplicationJSONObject
}
