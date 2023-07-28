// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type PokeathlonStatListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *PokeathlonStatListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *PokeathlonStatListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// PokeathlonStatList200ApplicationJSON - OK
type PokeathlonStatList200ApplicationJSON struct {
	// The total number of pokeathlon stats.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of pokeathlon stats.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of pokeathlon stat.
	Previous *string                 `json:"previous,omitempty"`
	Results  []shared.PokeathlonStat `json:"results,omitempty"`
}

func (o *PokeathlonStatList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *PokeathlonStatList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PokeathlonStatList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PokeathlonStatList200ApplicationJSON) GetResults() []shared.PokeathlonStat {
	if o == nil {
		return nil
	}
	return o.Results
}

type PokeathlonStatListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PokeathlonStatList200ApplicationJSONObject *PokeathlonStatList200ApplicationJSON
}

func (o *PokeathlonStatListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokeathlonStatListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokeathlonStatListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PokeathlonStatListResponse) GetPokeathlonStatList200ApplicationJSONObject() *PokeathlonStatList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PokeathlonStatList200ApplicationJSONObject
}
