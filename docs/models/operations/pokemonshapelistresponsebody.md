# PokemonShapeListResponseBody

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Count`                                                      | **int64*                                                     | :heavy_minus_sign:                                           | The total number of pokemon shapes.                          | 3                                                            |
| `Next`                                                       | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the next page of pokemon shapes.             | https://pokeapi.co/api/v2/pokemon-shape/?offset=20&limit=20  |
| `Previous`                                                   | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the previous page of pokemon shapes.         |                                                              |
| `Results`                                                    | [][shared.PokemonShape](../../models/shared/pokemonshape.md) | :heavy_minus_sign:                                           | N/A                                                          |                                                              |