# PokemonShape

### Available Operations

* [PokemonShapeList](#pokemonshapelist)
* [PokemonShapeRead](#pokemonshaperead)

## PokemonShapeList

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
    res, err := s.PokemonShape.PokemonShapeList(ctx, operations.PokemonShapeListRequest{
        Limit: homework.Int64(656330),
        Offset: homework.Int64(317202),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonShapeList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PokemonShapeListRequest](../../models/operations/pokemonshapelistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PokemonShapeListResponse](../../models/operations/pokemonshapelistresponse.md), error**


## PokemonShapeRead

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
    res, err := s.PokemonShape.PokemonShapeRead(ctx, operations.PokemonShapeReadRequest{
        ID: 138183,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonShape != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PokemonShapeReadRequest](../../models/operations/pokemonshapereadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PokemonShapeReadResponse](../../models/operations/pokemonshapereadresponse.md), error**

