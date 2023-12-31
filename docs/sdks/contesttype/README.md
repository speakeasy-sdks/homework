# ContestType

### Available Operations

* [ContestTypeList](#contesttypelist)
* [ContestTypeRead](#contesttyperead)

## ContestTypeList

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
    res, err := s.ContestType.ContestTypeList(ctx, operations.ContestTypeListRequest{
        Limit: homework.Int64(528895),
        Offset: homework.Int64(479977),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContestTypeList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ContestTypeListRequest](../../models/operations/contesttypelistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ContestTypeListResponse](../../models/operations/contesttypelistresponse.md), error**


## ContestTypeRead

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
    res, err := s.ContestType.ContestTypeRead(ctx, operations.ContestTypeReadRequest{
        ID: 568045,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContestType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ContestTypeReadRequest](../../models/operations/contesttypereadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ContestTypeReadResponse](../../models/operations/contesttypereadresponse.md), error**

