# ItemCategory

### Available Operations

* [ItemCategoryList](#itemcategorylist)
* [ItemCategoryRead](#itemcategoryread)

## ItemCategoryList

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
    res, err := s.ItemCategory.ItemCategoryList(ctx, operations.ItemCategoryListRequest{
        Limit: homework.Int64(414662),
        Offset: homework.Int64(473600),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemCategoryList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ItemCategoryListRequest](../../models/operations/itemcategorylistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ItemCategoryListResponse](../../models/operations/itemcategorylistresponse.md), error**


## ItemCategoryRead

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
    res, err := s.ItemCategory.ItemCategoryRead(ctx, operations.ItemCategoryReadRequest{
        ID: 264555,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ItemCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ItemCategoryReadRequest](../../models/operations/itemcategoryreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ItemCategoryReadResponse](../../models/operations/itemcategoryreadresponse.md), error**

