# Move

### Available Operations

* [MoveList](#movelist)
* [MoveRead](#moveread)

## MoveList

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
    res, err := s.Move.MoveList(ctx, operations.MoveListRequest{
        Limit: homework.Int64(449950),
        Offset: homework.Int64(359508),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.MoveListRequest](../../models/operations/movelistrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.MoveListResponse](../../models/operations/movelistresponse.md), error**


## MoveRead

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
    res, err := s.Move.MoveRead(ctx, operations.MoveReadRequest{
        ID: 613064,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Move != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.MoveReadRequest](../../models/operations/movereadrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.MoveReadResponse](../../models/operations/movereadresponse.md), error**

