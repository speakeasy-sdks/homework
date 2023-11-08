# MoveLearnMethodListResponseBody

OK


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Count`                                                            | **int64*                                                           | :heavy_minus_sign:                                                 | The total number of move learn methods.                            | 3                                                                  |
| `Next`                                                             | **string*                                                          | :heavy_minus_sign:                                                 | URL to retrieve the next page of move learn methods.               | https://pokeapi.co/api/v2/move-learn-method/?offset=20&limit=20    |
| `Previous`                                                         | **string*                                                          | :heavy_minus_sign:                                                 | URL to retrieve the previous page of move learn methods.           |                                                                    |
| `Results`                                                          | [][shared.MoveLearnMethod](../../models/shared/movelearnmethod.md) | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |