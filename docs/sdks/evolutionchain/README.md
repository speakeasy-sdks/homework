# EvolutionChain

### Available Operations

* [EvolutionChainList](#evolutionchainlist)
* [EvolutionChainRead](#evolutionchainread)

## EvolutionChainList

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
    res, err := s.EvolutionChain.EvolutionChainList(ctx, operations.EvolutionChainListRequest{
        Limit: homework.Int64(140350),
        Offset: homework.Int64(870013),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EvolutionChainList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EvolutionChainListRequest](../../models/operations/evolutionchainlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.EvolutionChainListResponse](../../models/operations/evolutionchainlistresponse.md), error**


## EvolutionChainRead

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
    res, err := s.EvolutionChain.EvolutionChainRead(ctx, operations.EvolutionChainReadRequest{
        ID: 870088,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EvolutionChain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EvolutionChainReadRequest](../../models/operations/evolutionchainreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.EvolutionChainReadResponse](../../models/operations/evolutionchainreadresponse.md), error**

