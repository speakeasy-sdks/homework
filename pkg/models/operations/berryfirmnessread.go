// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type BerryFirmnessReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *BerryFirmnessReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type BerryFirmnessReadResponse struct {
	// Successful response
	BerryFirmness *shared.BerryFirmness
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BerryFirmnessReadResponse) GetBerryFirmness() *shared.BerryFirmness {
	if o == nil {
		return nil
	}
	return o.BerryFirmness
}

func (o *BerryFirmnessReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BerryFirmnessReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BerryFirmnessReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
