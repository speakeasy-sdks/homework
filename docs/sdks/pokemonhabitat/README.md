# PokemonHabitat

### Available Operations

* [PokemonHabitatList](#pokemonhabitatlist)
* [PokemonHabitatRead](#pokemonhabitatread)

## PokemonHabitatList

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
    res, err := s.PokemonHabitat.PokemonHabitatList(ctx, operations.PokemonHabitatListRequest{
        Limit: homework.Int64(338007),
        Offset: homework.Int64(110375),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonHabitatList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokemonHabitatListRequest](../../models/operations/pokemonhabitatlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonHabitatListResponse](../../models/operations/pokemonhabitatlistresponse.md), error**


## PokemonHabitatRead

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
    res, err := s.PokemonHabitat.PokemonHabitatRead(ctx, operations.PokemonHabitatReadRequest{
        ID: 674752,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokemonHabitat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PokemonHabitatReadRequest](../../models/operations/pokemonhabitatreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonHabitatReadResponse](../../models/operations/pokemonhabitatreadresponse.md), error**

