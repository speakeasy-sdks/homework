# MoveCategory
(*MoveCategory*)

### Available Operations

* [MoveCategoryList](#movecategorylist)
* [MoveCategoryRead](#movecategoryread)

## MoveCategoryList

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
    res, err := s.MoveCategory.MoveCategoryList(ctx, operations.MoveCategoryListRequest{
        Limit: homework.Int64(883675),
        Offset: homework.Int64(142026),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveCategoryList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.MoveCategoryListRequest](../../models/operations/movecategorylistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.MoveCategoryListResponse](../../models/operations/movecategorylistresponse.md), error**


## MoveCategoryRead

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
    res, err := s.MoveCategory.MoveCategoryRead(ctx, operations.MoveCategoryReadRequest{
        ID: 69379,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.MoveCategoryReadRequest](../../models/operations/movecategoryreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.MoveCategoryReadResponse](../../models/operations/movecategoryreadresponse.md), error**

