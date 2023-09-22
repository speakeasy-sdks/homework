# GrowthRate

### Available Operations

* [GrowthRateList](#growthratelist)
* [GrowthRateRead](#growthrateread)

## GrowthRateList

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
    res, err := s.GrowthRate.GrowthRateList(ctx, operations.GrowthRateListRequest{
        Limit: homework.Int64(720633),
        Offset: homework.Int64(639921),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GrowthRateList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GrowthRateListRequest](../../models/operations/growthratelistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GrowthRateListResponse](../../models/operations/growthratelistresponse.md), error**


## GrowthRateRead

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
    res, err := s.GrowthRate.GrowthRateRead(ctx, operations.GrowthRateReadRequest{
        ID: 582020,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GrowthRate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GrowthRateReadRequest](../../models/operations/growthratereadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GrowthRateReadResponse](../../models/operations/growthratereadresponse.md), error**

