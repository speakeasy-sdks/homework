# EncounterConditionValue

### Available Operations

* [EncounterConditionValueList](#encounterconditionvaluelist)
* [EncounterConditionValueRead](#encounterconditionvalueread)

## EncounterConditionValueList

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
    res, err := s.EncounterConditionValue.EncounterConditionValueList(ctx, operations.EncounterConditionValueListRequest{
        Limit: homework.Int64(648172),
        Offset: homework.Int64(20218),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterConditionValues != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.EncounterConditionValueListRequest](../../models/operations/encounterconditionvaluelistrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.EncounterConditionValueListResponse](../../models/operations/encounterconditionvaluelistresponse.md), error**


## EncounterConditionValueRead

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
    res, err := s.EncounterConditionValue.EncounterConditionValueRead(ctx, operations.EncounterConditionValueReadRequest{
        ID: 368241,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EncounterConditionValue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.EncounterConditionValueReadRequest](../../models/operations/encounterconditionvaluereadrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.EncounterConditionValueReadResponse](../../models/operations/encounterconditionvaluereadresponse.md), error**

