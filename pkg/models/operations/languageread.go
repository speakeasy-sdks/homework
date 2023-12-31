// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type LanguageReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *LanguageReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type LanguageReadResponse struct {
	ContentType string
	// Successful response
	Language    *shared.Language
	StatusCode  int
	RawResponse *http.Response
}

func (o *LanguageReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LanguageReadResponse) GetLanguage() *shared.Language {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *LanguageReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LanguageReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
