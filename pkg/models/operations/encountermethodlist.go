// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EncounterMethodListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *EncounterMethodListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *EncounterMethodListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// EncounterMethodList200ApplicationJSON - OK
type EncounterMethodList200ApplicationJSON struct {
	// The total number of encounter methods.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of encounter methods.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of encounter methods.
	Previous *string                  `json:"previous,omitempty"`
	Results  []shared.EncounterMethod `json:"results,omitempty"`
}

func (o *EncounterMethodList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *EncounterMethodList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *EncounterMethodList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *EncounterMethodList200ApplicationJSON) GetResults() []shared.EncounterMethod {
	if o == nil {
		return nil
	}
	return o.Results
}

type EncounterMethodListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	EncounterMethodList200ApplicationJSONObject *EncounterMethodList200ApplicationJSON
}

func (o *EncounterMethodListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EncounterMethodListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EncounterMethodListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EncounterMethodListResponse) GetEncounterMethodList200ApplicationJSONObject() *EncounterMethodList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.EncounterMethodList200ApplicationJSONObject
}
