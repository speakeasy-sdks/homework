// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type VersionListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *VersionListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *VersionListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// VersionList200ApplicationJSON - OK
type VersionList200ApplicationJSON struct {
	// The total number of versions.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of versions.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of versions.
	Previous *string          `json:"previous,omitempty"`
	Results  []shared.Version `json:"results,omitempty"`
}

func (o *VersionList200ApplicationJSON) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *VersionList200ApplicationJSON) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *VersionList200ApplicationJSON) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *VersionList200ApplicationJSON) GetResults() []shared.Version {
	if o == nil {
		return nil
	}
	return o.Results
}

type VersionListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	VersionList200ApplicationJSONObject *VersionList200ApplicationJSON
}

func (o *VersionListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VersionListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VersionListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VersionListResponse) GetVersionList200ApplicationJSONObject() *VersionList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.VersionList200ApplicationJSONObject
}
