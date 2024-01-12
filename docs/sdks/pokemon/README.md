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
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.Pokemon.PokemonList(ctx, operations.PokemonListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PokemonListRequest](../../pkg/models/operations/pokemonlistrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PokemonListResponse](../../pkg/models/operations/pokemonlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PokemonRead

### Example Usage

```go
package main

import(
	"homework"
	"context"
	"homework/pkg/models/operations"
	"log"
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PokemonReadRequest](../../pkg/models/operations/pokemonreadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PokemonReadResponse](../../pkg/models/operations/pokemonreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
