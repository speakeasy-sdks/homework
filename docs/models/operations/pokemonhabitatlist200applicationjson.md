# PokemonHabitatList200ApplicationJSON

OK


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Count`                                                          | **int64*                                                         | :heavy_minus_sign:                                               | The total number of pokemon habitats.                            | 3                                                                |
| `Next`                                                           | **string*                                                        | :heavy_minus_sign:                                               | URL to retrieve the next page of pokemon habitats.               | https://pokeapi.co/api/v2/language/?offset=20&limit=20           |
| `Previous`                                                       | **string*                                                        | :heavy_minus_sign:                                               | URL to retrieve the previous page of pokemon habitats.           |                                                                  |
| `Results`                                                        | [][shared.PokemonHabitat](../../models/shared/pokemonhabitat.md) | :heavy_minus_sign:                                               | N/A                                                              |                                                                  |