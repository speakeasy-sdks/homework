# SuperContestEffect

### Available Operations

* [SuperContestEffectList](#supercontesteffectlist)
* [SuperContestEffectRead](#supercontesteffectread)

## SuperContestEffectList

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
    res, err := s.SuperContestEffect.SuperContestEffectList(ctx, operations.SuperContestEffectListRequest{
        Limit: homework.Int64(837945),
        Offset: homework.Int64(673660),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuperContestEffectList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.SuperContestEffectListRequest](../../models/operations/supercontesteffectlistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.SuperContestEffectListResponse](../../models/operations/supercontesteffectlistresponse.md), error**


## SuperContestEffectRead

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
    res, err := s.SuperContestEffect.SuperContestEffectRead(ctx, operations.SuperContestEffectReadRequest{
        ID: 96098,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuperContestEffect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.SuperContestEffectReadRequest](../../models/operations/supercontesteffectreadrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.SuperContestEffectReadResponse](../../models/operations/supercontesteffectreadresponse.md), error**

