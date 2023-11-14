# BerryReadResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Berry`                                                 | [*shared.Berry](../../../pkg/models/shared/berry.md)    | :heavy_minus_sign:                                      | Successful response                                     |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |