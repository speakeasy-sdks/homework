// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type ItemCategoryReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ItemCategoryReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ItemCategoryReadResponse struct {
	ContentType string
	// Successful response
	ItemCategory *shared.ItemCategory
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ItemCategoryReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ItemCategoryReadResponse) GetItemCategory() *shared.ItemCategory {
	if o == nil {
		return nil
	}
	return o.ItemCategory
}

func (o *ItemCategoryReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ItemCategoryReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
