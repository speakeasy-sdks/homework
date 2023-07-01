# PokemonFormList200ApplicationJSON

OK


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Count`                                                    | **int64*                                                   | :heavy_minus_sign:                                         | The total number of pokemon forms.                         | 3                                                          |
| `Next`                                                     | **string*                                                  | :heavy_minus_sign:                                         | URL to retrieve the next page of pokemon forms.            | https://pokeapi.co/api/v2/pokemon-form/?offset=20&limit=20 |
| `Previous`                                                 | **string*                                                  | :heavy_minus_sign:                                         | URL to retrieve the previous page of pokemon forms.        |                                                            |
| `Results`                                                  | [][shared.PokemonForm](../../models/shared/pokemonform.md) | :heavy_minus_sign:                                         | N/A                                                        |                                                            |