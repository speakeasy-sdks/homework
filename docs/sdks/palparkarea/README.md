# PalParkArea
(*PalParkArea*)

### Available Operations

* [PalParkAreaList](#palparkarealist)
* [PalParkAreaRead](#palparkarearead)

## PalParkAreaList

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
    res, err := s.PalParkArea.PalParkAreaList(ctx, operations.PalParkAreaListRequest{
        Limit: homework.Int64(450714),
        Offset: homework.Int64(409785),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PalParkAreaList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PalParkAreaListRequest](../../models/operations/palparkarealistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PalParkAreaListResponse](../../models/operations/palparkarealistresponse.md), error**


## PalParkAreaRead

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
    res, err := s.PalParkArea.PalParkAreaRead(ctx, operations.PalParkAreaReadRequest{
        ID: 950036,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PalParkArea != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PalParkAreaReadRequest](../../models/operations/palparkareareadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PalParkAreaReadResponse](../../models/operations/palparkareareadresponse.md), error**

