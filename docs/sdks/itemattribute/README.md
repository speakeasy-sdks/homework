# ItemAttribute

### Available Operations

* [ItemAttributeList](#itemattributelist)
* [ItemAttributeRead](#itemattributeread)

## ItemAttributeList

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
    res, err := s.ItemAttribute.ItemAttributeList(ctx, operations.ItemAttributeListRequest{
        Limit: homework.Int64(537373),
        Offset: homework.Int64(944669),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemAttributeList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ItemAttributeListRequest](../../models/operations/itemattributelistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ItemAttributeListResponse](../../models/operations/itemattributelistresponse.md), error**


## ItemAttributeRead

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
    res, err := s.ItemAttribute.ItemAttributeRead(ctx, operations.ItemAttributeReadRequest{
        ID: 758616,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemAttribute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ItemAttributeReadRequest](../../models/operations/itemattributereadrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ItemAttributeReadResponse](../../models/operations/itemattributereadresponse.md), error**

