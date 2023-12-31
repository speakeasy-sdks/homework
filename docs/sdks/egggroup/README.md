# EggGroup

### Available Operations

* [EggGroupList](#egggrouplist)
* [EggGroupRead](#egggroupread)

## EggGroupList

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
    res, err := s.EggGroup.EggGroupList(ctx, operations.EggGroupListRequest{
        Limit: homework.Int64(392785),
        Offset: homework.Int64(925597),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EggGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.EggGroupListRequest](../../models/operations/egggrouplistrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.EggGroupListResponse](../../models/operations/egggrouplistresponse.md), error**


## EggGroupRead

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
    res, err := s.EggGroup.EggGroupRead(ctx, operations.EggGroupReadRequest{
        ID: 836079,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EggGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.EggGroupReadRequest](../../models/operations/egggroupreadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.EggGroupReadResponse](../../models/operations/egggroupreadresponse.md), error**

