# ContestEffect

### Available Operations

* [ContestEffectList](#contesteffectlist)
* [ContestEffectRead](#contesteffectread)

## ContestEffectList

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
    res, err := s.ContestEffect.ContestEffectList(ctx, operations.ContestEffectListRequest{
        Limit: homework.Int64(272656),
        Offset: homework.Int64(383441),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContestEffectList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ContestEffectListRequest](../../models/operations/contesteffectlistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ContestEffectListResponse](../../models/operations/contesteffectlistresponse.md), error**


## ContestEffectRead

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
    res, err := s.ContestEffect.ContestEffectRead(ctx, operations.ContestEffectReadRequest{
        ID: 477665,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContestEffect != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ContestEffectReadRequest](../../models/operations/contesteffectreadrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ContestEffectReadResponse](../../models/operations/contesteffectreadresponse.md), error**

