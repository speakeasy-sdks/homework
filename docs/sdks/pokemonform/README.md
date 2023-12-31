# PokemonForm

### Available Operations

* [PokemonFormList](#pokemonformlist)
* [PokemonFormRead](#pokemonformread)

## PokemonFormList

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
    res, err := s.PokemonForm.PokemonFormList(ctx, operations.PokemonFormListRequest{
        Limit: homework.Int64(244425),
        Offset: homework.Int64(623510),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonFormList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PokemonFormListRequest](../../models/operations/pokemonformlistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PokemonFormListResponse](../../models/operations/pokemonformlistresponse.md), error**


## PokemonFormRead

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
    res, err := s.PokemonForm.PokemonFormRead(ctx, operations.PokemonFormReadRequest{
        ID: 158969,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonForm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PokemonFormReadRequest](../../models/operations/pokemonformreadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PokemonFormReadResponse](../../models/operations/pokemonformreadresponse.md), error**

