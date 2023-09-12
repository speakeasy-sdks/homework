# PokeathlonStatList200ApplicationJSON

OK


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Count`                                                          | **int64*                                                         | :heavy_minus_sign:                                               | The total number of pokeathlon stats.                            | 3                                                                |
| `Next`                                                           | **string*                                                        | :heavy_minus_sign:                                               | URL to retrieve the next page of pokeathlon stats.               | https://pokeapi.co/api/v2/pokeathlon-stat/?offset=20&limit=20    |
| `Previous`                                                       | **string*                                                        | :heavy_minus_sign:                                               | URL to retrieve the previous page of pokeathlon stat.            |                                                                  |
| `Results`                                                        | [][shared.PokeathlonStat](../../models/shared/pokeathlonstat.md) | :heavy_minus_sign:                                               | N/A                                                              |                                                                  |