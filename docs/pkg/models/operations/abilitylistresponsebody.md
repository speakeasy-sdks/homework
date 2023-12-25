# AbilityListResponseBody

A list of abilities


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Count`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | The total number of abilities available                   |
| `Next`                                                    | **string*                                                 | :heavy_minus_sign:                                        | The URL for the next page of abilities (null if none)     |
| `Previous`                                                | **string*                                                 | :heavy_minus_sign:                                        | The URL for the previous page of abilities (null if none) |
| `Results`                                                 | [][shared.Ability](../../../pkg/models/shared/ability.md) | :heavy_minus_sign:                                        | N/A                                                       |