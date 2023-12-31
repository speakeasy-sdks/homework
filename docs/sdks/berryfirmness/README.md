# BerryFirmness

### Available Operations

* [BerryFirmnessList](#berryfirmnesslist)
* [BerryFirmnessRead](#berryfirmnessread)

## BerryFirmnessList

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
    res, err := s.BerryFirmness.BerryFirmnessList(ctx, operations.BerryFirmnessListRequest{
        Limit: homework.Int64(645894),
        Offset: homework.Int64(384382),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BerryFirmnessList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.BerryFirmnessListRequest](../../models/operations/berryfirmnesslistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.BerryFirmnessListResponse](../../models/operations/berryfirmnesslistresponse.md), error**


## BerryFirmnessRead

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
    res, err := s.BerryFirmness.BerryFirmnessRead(ctx, operations.BerryFirmnessReadRequest{
        ID: 437587,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BerryFirmness != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.BerryFirmnessReadRequest](../../models/operations/berryfirmnessreadrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.BerryFirmnessReadResponse](../../models/operations/berryfirmnessreadresponse.md), error**

