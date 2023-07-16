# PokemonList200ApplicationJSON

OK


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Count`                                               | **int64*                                              | :heavy_minus_sign:                                    | The total number of pokemons.                         | 3                                                     |
| `Next`                                                | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the next page of pokemons.            | https://pokeapi.co/api/v2/pokemon/?offset=20&limit=20 |
| `Previous`                                            | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the previous page of pokemons.        |                                                       |
| `Results`                                             | [][shared.Pokemon](../../models/shared/pokemon.md)    | :heavy_minus_sign:                                    | N/A                                                   |                                                       |