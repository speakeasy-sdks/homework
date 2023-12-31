# Berry

### Available Operations

* [BerryList](#berrylist)
* [BerryRead](#berryread)

## BerryList

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
    res, err := s.Berry.BerryList(ctx, operations.BerryListRequest{
        Limit: homework.Int64(847252),
        Offset: homework.Int64(423655),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Berries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.BerryListRequest](../../models/operations/berrylistrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.BerryListResponse](../../models/operations/berrylistresponse.md), error**


## BerryRead

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
    res, err := s.Berry.BerryRead(ctx, operations.BerryReadRequest{
        ID: 623564,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Berry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.BerryReadRequest](../../models/operations/berryreadrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.BerryReadResponse](../../models/operations/berryreadresponse.md), error**

