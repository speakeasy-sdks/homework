# MoveAilment

### Available Operations

* [MoveAilmentList](#moveailmentlist)
* [MoveAilmentRead](#moveailmentread)

## MoveAilmentList

### Example Usage

```go
package main

import(
	"context"
	"log"
	"homework"
	"homework/pkg/models/operations"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.MoveAilment.MoveAilmentList(ctx, operations.MoveAilmentListRequest{
        Limit: homework.Int64(437032),
        Offset: homework.Int64(902349),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveAilmentList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.MoveAilmentListRequest](../../models/operations/moveailmentlistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.MoveAilmentListResponse](../../models/operations/moveailmentlistresponse.md), error**


## MoveAilmentRead

### Example Usage

```go
package main

import(
	"context"
	"log"
	"homework"
	"homework/pkg/models/operations"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.MoveAilment.MoveAilmentRead(ctx, operations.MoveAilmentReadRequest{
        ID: 697631,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveAilment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.MoveAilmentReadRequest](../../models/operations/moveailmentreadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.MoveAilmentReadResponse](../../models/operations/moveailmentreadresponse.md), error**

