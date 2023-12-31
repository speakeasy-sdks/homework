# EncounterCondition

### Available Operations

* [EncounterConditionList](#encounterconditionlist)
* [EncounterConditionRead](#encounterconditionread)

## EncounterConditionList

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
    res, err := s.EncounterCondition.EncounterConditionList(ctx, operations.EncounterConditionListRequest{
        Limit: homework.Int64(71036),
        Offset: homework.Int64(337396),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterConditionList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EncounterConditionListRequest](../../models/operations/encounterconditionlistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EncounterConditionListResponse](../../models/operations/encounterconditionlistresponse.md), error**


## EncounterConditionRead

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
    res, err := s.EncounterCondition.EncounterConditionRead(ctx, operations.EncounterConditionReadRequest{
        ID: 87129,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterCondition != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EncounterConditionReadRequest](../../models/operations/encounterconditionreadrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EncounterConditionReadResponse](../../models/operations/encounterconditionreadresponse.md), error**

