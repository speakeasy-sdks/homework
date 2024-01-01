// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EncounterConditionReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *EncounterConditionReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type EncounterConditionReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	EncounterCondition *shared.EncounterCondition
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *EncounterConditionReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EncounterConditionReadResponse) GetEncounterCondition() *shared.EncounterCondition {
	if o == nil {
		return nil
	}
	return o.EncounterCondition
}

func (o *EncounterConditionReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EncounterConditionReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
