# ContestEffectListResponseBody

A list of contest effects


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Count`                                                               | **int64*                                                              | :heavy_minus_sign:                                                    | The total number of contest effects                                   |
| `Next`                                                                | **string*                                                             | :heavy_minus_sign:                                                    | The URL to get the next page of contest effects (if it exists)        |
| `Previous`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | The URL to get the previous page of contest effects (if it exists)    |
| `Results`                                                             | [][shared.ContestEffect](../../../pkg/models/shared/contesteffect.md) | :heavy_minus_sign:                                                    | N/A                                                                   |