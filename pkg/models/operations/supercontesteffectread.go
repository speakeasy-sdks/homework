// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type SuperContestEffectReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *SuperContestEffectReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type SuperContestEffectReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	SuperContestEffect *shared.SuperContestEffect
}

func (o *SuperContestEffectReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SuperContestEffectReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SuperContestEffectReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SuperContestEffectReadResponse) GetSuperContestEffect() *shared.SuperContestEffect {
	if o == nil {
		return nil
	}
	return o.SuperContestEffect
}
