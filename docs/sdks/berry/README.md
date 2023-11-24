# Berry
(*Berry*)

### Available Operations

* [BerryList](#berrylist)
* [BerryRead](#berryread)

## BerryList

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
    res, err := s.Berry.BerryList(ctx, operations.BerryListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.BerryListRequest](../../pkg/models/operations/berrylistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.BerryListResponse](../../pkg/models/operations/berrylistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## BerryRead

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
    res, err := s.Berry.BerryRead(ctx, operations.BerryReadRequest{
        ID: 104763,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Berry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.BerryReadRequest](../../pkg/models/operations/berryreadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.BerryReadResponse](../../pkg/models/operations/berryreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
