// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveBattleStyleListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveBattleStyleListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveBattleStyleListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// MoveBattleStyleList200ApplicationJSON - OK
type MoveBattleStyleList200ApplicationJSON struct {
	// The total number of move battle styles.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move battle styles.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move battle styles.
	Previous *string                  `json:"previous,omitempty"`
	Results  []shared.MoveBattleStyle `json:"results,omitempty"`
}

func (o *MoveBattleStyleList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveBattleStyleList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveBattleStyleList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveBattleStyleList200ApplicationJSON) GetResults() []shared.MoveBattleStyle {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveBattleStyleListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveBattleStyleList200ApplicationJSONObject *MoveBattleStyleList200ApplicationJSON
}

func (o *MoveBattleStyleListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveBattleStyleListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveBattleStyleListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveBattleStyleListResponse) GetMoveBattleStyleList200ApplicationJSONObject() *MoveBattleStyleList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MoveBattleStyleList200ApplicationJSONObject
}
