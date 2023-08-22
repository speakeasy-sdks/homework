# PalParkEncounterArea

Areas where the Pokémon species can be encountered in pal park.


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `Area`                                                                                          | [*NamedAPIResource](../../models/shared/namedapiresource.md)                                    | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `BaseScore`                                                                                     | **int64*                                                                                        | :heavy_minus_sign:                                                                              | The base score given to the player when the referenced Pokemon is caught during a pal park run. |
| `Rate`                                                                                          | **int64*                                                                                        | :heavy_minus_sign:                                                                              | The base rate for encountering the referenced Pokemon in this pal park area.                    |