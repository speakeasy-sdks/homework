// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"homework/pkg/models/shared"
	"net/http"
)

type EggGroupListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type EggGroupListResponse struct {
	ContentType string
	// OK
	EggGroups   []shared.EggGroup
	StatusCode  int
	RawResponse *http.Response
}
