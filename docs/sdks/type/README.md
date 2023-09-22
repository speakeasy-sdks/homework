# Type

### Available Operations

* [TypeList](#typelist)
* [TypeRead](#typeread)

## TypeList

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
    res, err := s.Type.TypeList(ctx, operations.TypeListRequest{
        Limit: homework.Int64(971945),
        Offset: homework.Int64(976460),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TypeList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.TypeListRequest](../../models/operations/typelistrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.TypeListResponse](../../models/operations/typelistresponse.md), error**


## TypeRead

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
    res, err := s.Type.TypeRead(ctx, operations.TypeReadRequest{
        ID: 878194,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Type != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.TypeReadRequest](../../models/operations/typereadrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.TypeReadResponse](../../models/operations/typereadresponse.md), error**

