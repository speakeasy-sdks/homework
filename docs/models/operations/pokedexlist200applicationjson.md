# PokedexList200ApplicationJSON

OK


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Count`                                               | **int64*                                              | :heavy_minus_sign:                                    | The total number of pokedexes.                        | 3                                                     |
| `Next`                                                | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the next page of pokedexes.           | https://pokeapi.co/api/v2/pokedex/?offset=20&limit=20 |
| `Previous`                                            | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the previous page of pokedexes.       |                                                       |
| `Results`                                             | [][shared.Pokedex](../../models/shared/pokedex.md)    | :heavy_minus_sign:                                    | N/A                                                   |                                                       |