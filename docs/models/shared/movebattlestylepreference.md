# MoveBattleStylePreference


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `HighHpPreference`                                                           | *int64*                                                                      | :heavy_check_mark:                                                           | Chance of using the move, in percent, if HP is over one half of maximum HP.  |
| `LowHpPreference`                                                            | *int64*                                                                      | :heavy_check_mark:                                                           | Chance of using the move, in percent, if HP is under one half of maximum HP. |
| `MoveBattleStyle`                                                            | [shared.NamedAPIResource](../../models/shared/namedapiresource.md)           | :heavy_check_mark:                                                           | N/A                                                                          |