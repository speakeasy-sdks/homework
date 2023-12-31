# MoveBattleStyle

### Available Operations

* [MoveBattleStyleList](#movebattlestylelist)
* [MoveBattleStyleRead](#movebattlestyleread)

## MoveBattleStyleList

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
    res, err := s.MoveBattleStyle.MoveBattleStyleList(ctx, operations.MoveBattleStyleListRequest{
        Limit: homework.Int64(99280),
        Offset: homework.Int64(60225),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveBattleStyleList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveBattleStyleListRequest](../../models/operations/movebattlestylelistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveBattleStyleListResponse](../../models/operations/movebattlestylelistresponse.md), error**


## MoveBattleStyleRead

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
    res, err := s.MoveBattleStyle.MoveBattleStyleRead(ctx, operations.MoveBattleStyleReadRequest{
        ID: 969810,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveBattleStyle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveBattleStyleReadRequest](../../models/operations/movebattlestylereadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveBattleStyleReadResponse](../../models/operations/movebattlestylereadresponse.md), error**

