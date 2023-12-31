// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemAttributeListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *ItemAttributeListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ItemAttributeListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// ItemAttributeList200ApplicationJSON - OK
type ItemAttributeList200ApplicationJSON struct {
	// The total number of item attributes.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of item attributes.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of item attributes.
	Previous *string                `json:"previous,omitempty"`
	Results  []shared.ItemAttribute `json:"results,omitempty"`
}

func (o *ItemAttributeList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ItemAttributeList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ItemAttributeList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ItemAttributeList200ApplicationJSON) GetResults() []shared.ItemAttribute {
	if o == nil {
		return nil
	}
	return o.Results
}

type ItemAttributeListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ItemAttributeList200ApplicationJSONObject *ItemAttributeList200ApplicationJSON
}

func (o *ItemAttributeListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemAttributeListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemAttributeListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ItemAttributeListResponse) GetItemAttributeList200ApplicationJSONObject() *ItemAttributeList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ItemAttributeList200ApplicationJSONObject
}
