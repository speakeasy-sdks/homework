// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type MoveReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *MoveReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type MoveReadResponse struct {
	ContentType string
	// Successful response
	Move        *shared.Move
	StatusCode  int
	RawResponse *http.Response
}

func (o *MoveReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveReadResponse) GetMove() *shared.Move {
	if o == nil {
		return nil
	}
	return o.Move
}

func (o *MoveReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
