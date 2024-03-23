# ItemCategory
(*ItemCategory*)

### Available Operations

* [ItemCategoryList](#itemcategorylist)
* [ItemCategoryRead](#itemcategoryread)

## ItemCategoryList

### Example Usage

```go
package main

import(
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.ItemCategory.ItemCategoryList(ctx, operations.ItemCategoryListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ItemCategoryListRequest](../../pkg/models/operations/itemcategorylistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ItemCategoryListResponse](../../pkg/models/operations/itemcategorylistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ItemCategoryRead

### Example Usage

```go
package main

import(
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.ItemCategory.ItemCategoryRead(ctx, operations.ItemCategoryReadRequest{
        ID: 708802,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ItemCategoryReadRequest](../../pkg/models/operations/itemcategoryreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ItemCategoryReadResponse](../../pkg/models/operations/itemcategoryreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
