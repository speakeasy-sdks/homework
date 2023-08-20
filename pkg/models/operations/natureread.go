// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type NatureReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *NatureReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type NatureReadResponse struct {
	ContentType string
	// Successful response
	Nature      *shared.Nature
	StatusCode  int
	RawResponse *http.Response
}

func (o *NatureReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *NatureReadResponse) GetNature() *shared.Nature {
	if o == nil {
		return nil
	}
	return o.Nature
}

func (o *NatureReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *NatureReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
