# Pokemon
(*Pokemon*)

### Available Operations

* [PokemonList](#pokemonlist)
* [PokemonRead](#pokemonread)

## PokemonList

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
    res, err := s.Pokemon.PokemonList(ctx, operations.PokemonListRequest{
        Limit: homework.Int64(824181),
        Offset: homework.Int64(225561),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PokemonListRequest](../../models/operations/pokemonlistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PokemonListResponse](../../models/operations/pokemonlistresponse.md), error**


## PokemonRead

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
    res, err := s.Pokemon.PokemonRead(ctx, operations.PokemonReadRequest{
        ID: 386997,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pokemon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PokemonReadRequest](../../models/operations/pokemonreadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PokemonReadResponse](../../models/operations/pokemonreadresponse.md), error**

