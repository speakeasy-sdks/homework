# MoveLearnMethod

### Available Operations

* [MoveLearnMethodList](#movelearnmethodlist)
* [MoveLearnMethodRead](#movelearnmethodread)

## MoveLearnMethodList

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
    res, err := s.MoveLearnMethod.MoveLearnMethodList(ctx, operations.MoveLearnMethodListRequest{
        Limit: homework.Int64(128926),
        Offset: homework.Int64(750686),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveLearnMethodList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveLearnMethodListRequest](../../models/operations/movelearnmethodlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveLearnMethodListResponse](../../models/operations/movelearnmethodlistresponse.md), error**


## MoveLearnMethodRead

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
    res, err := s.MoveLearnMethod.MoveLearnMethodRead(ctx, operations.MoveLearnMethodReadRequest{
        ID: 315428,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MoveLearnMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveLearnMethodReadRequest](../../models/operations/movelearnmethodreadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveLearnMethodReadResponse](../../models/operations/movelearnmethodreadresponse.md), error**

