// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type LanguageListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *LanguageListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *LanguageListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// LanguageListResponseBody - OK
type LanguageListResponseBody struct {
	// The total number of languages.
	Count *int64 `json:"count,omitempty"`
	// URL to retrieve the next page of languages.
	Next *string `json:"next,omitempty"`
	// URL to retrieve the previous page of languages.
	Previous *string           `json:"previous,omitempty"`
	Results  []shared.Language `json:"results,omitempty"`
}

func (o *LanguageListResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *LanguageListResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *LanguageListResponseBody) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *LanguageListResponseBody) GetResults() []shared.Language {
	if o == nil {
		return nil
	}
	return o.Results
}

type LanguageListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *LanguageListResponseBody
}

func (o *LanguageListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LanguageListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LanguageListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LanguageListResponse) GetObject() *LanguageListResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
