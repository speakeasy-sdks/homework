# PokemonColorListResponseBody

OK


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Count`                                                             | **int64*                                                            | :heavy_minus_sign:                                                  | The total number of pokemon colors.                                 | 3                                                                   |
| `Next`                                                              | **string*                                                           | :heavy_minus_sign:                                                  | URL to retrieve the next page of pokemon colors.                    | https://pokeapi.co/api/v2/pokemon-color/?offset=20&limit=20         |
| `Previous`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | URL to retrieve the previous page of pokemon colors.                |                                                                     |
| `Results`                                                           | [][shared.PokemonColor](../../../pkg/models/shared/pokemoncolor.md) | :heavy_minus_sign:                                                  | N/A                                                                 |                                                                     |