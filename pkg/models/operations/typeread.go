// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type TypeReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *TypeReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type TypeReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	Type *shared.Type
}

func (o *TypeReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TypeReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TypeReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TypeReadResponse) GetType() *shared.Type {
	if o == nil {
		return nil
	}
	return o.Type
}
