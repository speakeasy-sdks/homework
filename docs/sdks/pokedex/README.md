# Pokedex
(*Pokedex*)

### Available Operations

* [PokedexList](#pokedexlist)
* [PokedexRead](#pokedexread)

## PokedexList

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
    res, err := s.Pokedex.PokedexList(ctx, operations.PokedexListRequest{
        Limit: homework.Int64(731646),
        Offset: homework.Int64(727103),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PokedexList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PokedexListRequest](../../models/operations/pokedexlistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PokedexListResponse](../../models/operations/pokedexlistresponse.md), error**


## PokedexRead

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
    res, err := s.Pokedex.PokedexRead(ctx, operations.PokedexReadRequest{
        ID: 49683,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pokedex != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PokedexReadRequest](../../models/operations/pokedexreadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PokedexReadResponse](../../models/operations/pokedexreadresponse.md), error**

