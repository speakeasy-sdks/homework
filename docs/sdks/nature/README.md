# Nature
(*Nature*)

### Available Operations

* [NatureList](#naturelist)
* [NatureRead](#natureread)

## NatureList

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
    res, err := s.Nature.NatureList(ctx, operations.NatureListRequest{
        Limit: homework.Int64(654864),
        Offset: homework.Int64(615974),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NatureList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.NatureListRequest](../../models/operations/naturelistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.NatureListResponse](../../models/operations/naturelistresponse.md), error**


## NatureRead

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
    res, err := s.Nature.NatureRead(ctx, operations.NatureReadRequest{
        ID: 756843,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Nature != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.NatureReadRequest](../../models/operations/naturereadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.NatureReadResponse](../../models/operations/naturereadresponse.md), error**

