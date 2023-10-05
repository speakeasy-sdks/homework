# MoveTarget
(*MoveTarget*)

### Available Operations

* [MoveTargetList](#movetargetlist)
* [MoveTargetRead](#movetargetread)

## MoveTargetList

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
    res, err := s.MoveTarget.MoveTargetList(ctx, operations.MoveTargetListRequest{
        Limit: homework.Int64(425857),
        Offset: homework.Int64(232779),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveTargetList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.MoveTargetListRequest](../../models/operations/movetargetlistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.MoveTargetListResponse](../../models/operations/movetargetlistresponse.md), error**


## MoveTargetRead

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
    res, err := s.MoveTarget.MoveTargetRead(ctx, operations.MoveTargetReadRequest{
        ID: 455703,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveTarget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.MoveTargetReadRequest](../../models/operations/movetargetreadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.MoveTargetReadResponse](../../models/operations/movetargetreadresponse.md), error**

