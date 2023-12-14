# ContestEffect


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Appeal`                                                              | *int64*                                                               | :heavy_check_mark:                                                    | The base number of hearts the user of this move gets                  |
| `EffectEntries`                                                       | [][shared.VerboseEffect](../../../pkg/models/shared/verboseeffect.md) | :heavy_check_mark:                                                    | The flavor text of this contest effect listed in different languages  |
| `FlavorTextEntries`                                                   | [][shared.FlavorText](../../../pkg/models/shared/flavortext.md)       | :heavy_check_mark:                                                    | The flavor text of this contest effect listed in different languages  |
| `ID`                                                                  | *int64*                                                               | :heavy_check_mark:                                                    | The identifier for this contest effect resource                       |
| `Jam`                                                                 | *int64*                                                               | :heavy_check_mark:                                                    | The base number of hearts the user's opponent loses                   |