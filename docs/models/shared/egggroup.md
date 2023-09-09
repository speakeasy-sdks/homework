# EggGroup


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | *int64*                                                          | :heavy_check_mark:                                               | The identifier for this egg group resource                       |
| `Name`                                                           | *string*                                                         | :heavy_check_mark:                                               | The name for this egg group resource                             |
| `Names`                                                          | [][Name](../../models/shared/name.md)                            | :heavy_check_mark:                                               | The name of this egg group listed in different languages         |
| `PokemonSpecies`                                                 | [][NamedAPIResource](../../models/shared/namedapiresource.md)    | :heavy_check_mark:                                               | A list of all Pokemon species that are members of this egg group |