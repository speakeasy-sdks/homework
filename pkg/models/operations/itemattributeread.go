// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemAttributeReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ItemAttributeReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ItemAttributeReadResponse struct {
	ContentType string
	// Successful response
	ItemAttribute *shared.ItemAttribute
	StatusCode    int
	RawResponse   *http.Response
}

func (o *ItemAttributeReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemAttributeReadResponse) GetItemAttribute() *shared.ItemAttribute {
	if o == nil {
		return nil
	}
	return o.ItemAttribute
}

func (o *ItemAttributeReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemAttributeReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
