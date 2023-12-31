# PokeathlonStat

### Available Operations

* [PokeathlonStatList](#pokeathlonstatlist)
* [PokeathlonStatRead](#pokeathlonstatread)

## PokeathlonStatList

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
    res, err := s.PokeathlonStat.PokeathlonStatList(ctx, operations.PokeathlonStatListRequest{
        Limit: homework.Int64(102044),
        Offset: homework.Int64(652790),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokeathlonStatList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokeathlonStatListRequest](../../models/operations/pokeathlonstatlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokeathlonStatListResponse](../../models/operations/pokeathlonstatlistresponse.md), error**


## PokeathlonStatRead

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
    res, err := s.PokeathlonStat.PokeathlonStatRead(ctx, operations.PokeathlonStatReadRequest{
        ID: 208876,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokeathlonStat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokeathlonStatReadRequest](../../models/operations/pokeathlonstatreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokeathlonStatReadResponse](../../models/operations/pokeathlonstatreadresponse.md), error**

