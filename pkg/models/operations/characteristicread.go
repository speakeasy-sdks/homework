// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type CharacteristicReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *CharacteristicReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type CharacteristicReadResponse struct {
	// Successful response
	Characteristic *shared.Characteristic
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *CharacteristicReadResponse) GetCharacteristic() *shared.Characteristic {
	if o == nil {
		return nil
	}
	return o.Characteristic
}

func (o *CharacteristicReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CharacteristicReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CharacteristicReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
