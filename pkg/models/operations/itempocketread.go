// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemPocketReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ItemPocketReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ItemPocketReadResponse struct {
	ContentType string
	// Successful response
	ItemPocket  *shared.ItemPocket
	StatusCode  int
	RawResponse *http.Response
}

func (o *ItemPocketReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemPocketReadResponse) GetItemPocket() *shared.ItemPocket {
	if o == nil {
		return nil
	}
	return o.ItemPocket
}

func (o *ItemPocketReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemPocketReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
