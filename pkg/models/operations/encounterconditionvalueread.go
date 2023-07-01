// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EncounterConditionValueReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type EncounterConditionValueReadResponse struct {
	ContentType string
	// Successful response
	EncounterConditionValue *shared.EncounterConditionValue
	StatusCode              int
	RawResponse             *http.Response
}
