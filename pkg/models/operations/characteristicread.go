// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type CharacteristicReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type CharacteristicReadResponse struct {
	// Successful response
	Characteristic *shared.Characteristic
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
