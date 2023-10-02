# EvolutionTrigger
(*EvolutionTrigger*)

### Available Operations

* [EvolutionTriggerList](#evolutiontriggerlist)
* [EvolutionTriggerRead](#evolutiontriggerread)

## EvolutionTriggerList

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
    res, err := s.EvolutionTrigger.EvolutionTriggerList(ctx, operations.EvolutionTriggerListRequest{
        Limit: homework.Int64(841396),
        Offset: homework.Int64(113202),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EvolutionTriggerList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.EvolutionTriggerListRequest](../../models/operations/evolutiontriggerlistrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.EvolutionTriggerListResponse](../../models/operations/evolutiontriggerlistresponse.md), error**


## EvolutionTriggerRead

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
    res, err := s.EvolutionTrigger.EvolutionTriggerRead(ctx, operations.EvolutionTriggerReadRequest{
        ID: 848636,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EvolutionTrigger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.EvolutionTriggerReadRequest](../../models/operations/evolutiontriggerreadrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.EvolutionTriggerReadResponse](../../models/operations/evolutiontriggerreadresponse.md), error**

