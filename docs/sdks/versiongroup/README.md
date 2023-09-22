# VersionGroup

### Available Operations

* [VersionGroupList](#versiongrouplist)
* [VersionGroupRead](#versiongroupread)

## VersionGroupList

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
    res, err := s.VersionGroup.VersionGroupList(ctx, operations.VersionGroupListRequest{
        Limit: homework.Int64(55714),
        Offset: homework.Int64(604846),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionGroupList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.VersionGroupListRequest](../../models/operations/versiongrouplistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.VersionGroupListResponse](../../models/operations/versiongrouplistresponse.md), error**


## VersionGroupRead

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
    res, err := s.VersionGroup.VersionGroupRead(ctx, operations.VersionGroupReadRequest{
        ID: 451159,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.VersionGroupReadRequest](../../models/operations/versiongroupreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.VersionGroupReadResponse](../../models/operations/versiongroupreadresponse.md), error**

