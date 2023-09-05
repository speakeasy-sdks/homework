// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ItemReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ItemReadResponse struct {
	ContentType string
	// Successful response
	Item        *shared.Item
	StatusCode  int
	RawResponse *http.Response
}

func (o *ItemReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemReadResponse) GetItem() *shared.Item {
	if o == nil {
		return nil
	}
	return o.Item
}

func (o *ItemReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
