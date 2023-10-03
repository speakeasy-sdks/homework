# Stat
(*Stat*)

### Available Operations

* [StatList](#statlist)
* [StatRead](#statread)

## StatList

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
    res, err := s.Stat.StatList(ctx, operations.StatListRequest{
        Limit: homework.Int64(771174),
        Offset: homework.Int64(892933),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.StatListRequest](../../models/operations/statlistrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.StatListResponse](../../models/operations/statlistresponse.md), error**


## StatRead

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
    res, err := s.Stat.StatRead(ctx, operations.StatReadRequest{
        ID: 160697,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Stat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.StatReadRequest](../../models/operations/statreadrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.StatReadResponse](../../models/operations/statreadresponse.md), error**

