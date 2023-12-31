// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveCategoryReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *MoveCategoryReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type MoveCategoryReadResponse struct {
	ContentType string
	// Successful response
	MoveCategory *shared.MoveCategory
	StatusCode   int
	RawResponse  *http.Response
}

func (o *MoveCategoryReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveCategoryReadResponse) GetMoveCategory() *shared.MoveCategory {
	if o == nil {
		return nil
	}
	return o.MoveCategory
}

func (o *MoveCategoryReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveCategoryReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
