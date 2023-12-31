# EncounterMethod

### Available Operations

* [EncounterMethodList](#encountermethodlist)
* [EncounterMethodRead](#encountermethodread)

## EncounterMethodList

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
    res, err := s.EncounterMethod.EncounterMethodList(ctx, operations.EncounterMethodListRequest{
        Limit: homework.Int64(832620),
        Offset: homework.Int64(957156),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterMethodList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.EncounterMethodListRequest](../../models/operations/encountermethodlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.EncounterMethodListResponse](../../models/operations/encountermethodlistresponse.md), error**


## EncounterMethodRead

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
    res, err := s.EncounterMethod.EncounterMethodRead(ctx, operations.EncounterMethodReadRequest{
        ID: 778157,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.EncounterMethodReadRequest](../../models/operations/encountermethodreadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.EncounterMethodReadResponse](../../models/operations/encountermethodreadresponse.md), error**

