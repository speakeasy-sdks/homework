# ItemPocket

### Available Operations

* [ItemPocketList](#itempocketlist)
* [ItemPocketRead](#itempocketread)

## ItemPocketList

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
    res, err := s.ItemPocket.ItemPocketList(ctx, operations.ItemPocketListRequest{
        Limit: homework.Int64(456150),
        Offset: homework.Int64(216550),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemPocketList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ItemPocketListRequest](../../models/operations/itempocketlistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ItemPocketListResponse](../../models/operations/itempocketlistresponse.md), error**


## ItemPocketRead

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
    res, err := s.ItemPocket.ItemPocketRead(ctx, operations.ItemPocketReadRequest{
        ID: 568434,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemPocket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ItemPocketReadRequest](../../models/operations/itempocketreadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ItemPocketReadResponse](../../models/operations/itempocketreadresponse.md), error**

