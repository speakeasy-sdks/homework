# PokemonSpeciesListResponseBody

OK


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Count`                                                                    | **int64*                                                                   | :heavy_minus_sign:                                                         | The total number of pokemon species list.                                  | 3                                                                          |
| `Next`                                                                     | **string*                                                                  | :heavy_minus_sign:                                                         | URL to retrieve the next page of pokemon species list.                     | https://pokeapi.co/api/v2/pokemon-species/?offset=20&limit=20              |
| `Previous`                                                                 | **string*                                                                  | :heavy_minus_sign:                                                         | URL to retrieve the previous page of pokemon species list.                 |                                                                            |
| `Results`                                                                  | [][shared.PokemonSpeciesInput](../../models/shared/pokemonspeciesinput.md) | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |