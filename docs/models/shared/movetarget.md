# MoveTarget

Targets moves can be directed at during battle. Targets can be Pokémon, adjacent positions, all opponents, etc.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Descriptions`                                                    | [][shared.Description](../../models/shared/description.md)        | :heavy_check_mark:                                                | The description of this move target listed in different languages |                                                                   |
| `ID`                                                              | *int64*                                                           | :heavy_check_mark:                                                | The identifier for this move target resource                      | 1                                                                 |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | The name for this move target resource                            | specific-move                                                     |