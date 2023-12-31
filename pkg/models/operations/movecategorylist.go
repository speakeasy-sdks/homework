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

// MoveCategoryList200ApplicationJSON - OK
type MoveCategoryList200ApplicationJSON struct {
	// The total number of move categories.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of move categories.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of move categories.
	Previous *string               `json:"previous,omitempty"`
	Results  []shared.MoveCategory `json:"results,omitempty"`
}

func (o *MoveCategoryList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MoveCategoryList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MoveCategoryList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *MoveCategoryList200ApplicationJSON) GetResults() []shared.MoveCategory {
	if o == nil {
		return nil
	}
	return o.Results
}

type MoveCategoryListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MoveCategoryList200ApplicationJSONObject *MoveCategoryList200ApplicationJSON
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

func (o *MoveCategoryListResponse) GetMoveCategoryList200ApplicationJSONObject() *MoveCategoryList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MoveCategoryList200ApplicationJSONObject
}
