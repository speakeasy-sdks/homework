// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EggGroupListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *EggGroupListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *EggGroupListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type EggGroupListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	EggGroups []shared.EggGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *EggGroupListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EggGroupListResponse) GetEggGroups() []shared.EggGroup {
	if o == nil {
		return nil
	}
	return o.EggGroups
}

func (o *EggGroupListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EggGroupListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
