# Generation

### Available Operations

* [GenerationList](#generationlist)
* [GenerationRead](#generationread)

## GenerationList

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
    res, err := s.Generation.GenerationList(ctx, operations.GenerationListRequest{
        Limit: homework.Int64(780529),
        Offset: homework.Int64(678880),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerationList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GenerationListRequest](../../models/operations/generationlistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GenerationListResponse](../../models/operations/generationlistresponse.md), error**


## GenerationRead

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
    res, err := s.Generation.GenerationRead(ctx, operations.GenerationReadRequest{
        ID: 118274,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Generation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GenerationReadRequest](../../models/operations/generationreadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GenerationReadResponse](../../models/operations/generationreadresponse.md), error**

