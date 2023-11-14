# Pokedex


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Descriptions`                                                            | [][shared.Description](../../../pkg/models/shared/description.md)         | :heavy_check_mark:                                                        | N/A                                                                       |                                                                           |
| `ID`                                                                      | *int64*                                                                   | :heavy_check_mark:                                                        | N/A                                                                       | 2                                                                         |
| `IsMainSeries`                                                            | *bool*                                                                    | :heavy_check_mark:                                                        | N/A                                                                       | true                                                                      |
| `Name`                                                                    | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       | National                                                                  |
| `PokemonEntries`                                                          | [][shared.PokemonEntry](../../../pkg/models/shared/pokemonentry.md)       | :heavy_check_mark:                                                        | N/A                                                                       |                                                                           |
| `Region`                                                                  | [shared.NamedAPIResource](../../../pkg/models/shared/namedapiresource.md) | :heavy_check_mark:                                                        | N/A                                                                       |                                                                           |