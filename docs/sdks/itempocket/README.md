# ItemPocket
(*ItemPocket*)

### Available Operations

* [ItemPocketList](#itempocketlist)
* [ItemPocketRead](#itempocketread)

## ItemPocketList

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
    res, err := s.ItemPocket.ItemPocketList(ctx, operations.ItemPocketListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ItemPocketListRequest](../../pkg/models/operations/itempocketlistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ItemPocketListResponse](../../pkg/models/operations/itempocketlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ItemPocketRead

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
    res, err := s.ItemPocket.ItemPocketRead(ctx, operations.ItemPocketReadRequest{
        ID: 969735,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ItemPocket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ItemPocketReadRequest](../../pkg/models/operations/itempocketreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ItemPocketReadResponse](../../pkg/models/operations/itempocketreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
