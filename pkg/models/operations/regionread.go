// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type RegionReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RegionReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type RegionReadResponse struct {
	ContentType string
	// Successful response
	Region      *shared.Region
	StatusCode  int
	RawResponse *http.Response
}

func (o *RegionReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RegionReadResponse) GetRegion() *shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *RegionReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RegionReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
