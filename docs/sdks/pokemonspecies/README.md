# PokemonSpecies

### Available Operations

* [PokemonSpeciesList](#pokemonspecieslist)
* [PokemonSpeciesRead](#pokemonspeciesread)

## PokemonSpeciesList

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
    res, err := s.PokemonSpecies.PokemonSpeciesList(ctx, operations.PokemonSpeciesListRequest{
        Limit: homework.Int64(778346),
        Offset: homework.Int64(196582),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonSpeciesList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokemonSpeciesListRequest](../../models/operations/pokemonspecieslistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonSpeciesListResponse](../../models/operations/pokemonspecieslistresponse.md), error**


## PokemonSpeciesRead

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
    res, err := s.PokemonSpecies.PokemonSpeciesRead(ctx, operations.PokemonSpeciesReadRequest{
        ID: 949572,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonSpecies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokemonSpeciesReadRequest](../../models/operations/pokemonspeciesreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonSpeciesReadResponse](../../models/operations/pokemonspeciesreadresponse.md), error**

