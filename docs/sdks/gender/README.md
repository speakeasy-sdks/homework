# Gender
(*Gender*)

### Available Operations

* [GenderList](#genderlist)
* [GenderRead](#genderread)

## GenderList

### Example Usage

```go
package main

import(
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.Gender.GenderList(ctx, operations.GenderListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GenderListRequest](../../pkg/models/operations/genderlistrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GenderListResponse](../../pkg/models/operations/genderlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GenderRead

### Example Usage

```go
package main

import(
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.Gender.GenderRead(ctx, operations.GenderReadRequest{
        ID: 403245,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Gender != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GenderReadRequest](../../pkg/models/operations/genderreadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GenderReadResponse](../../pkg/models/operations/genderreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
