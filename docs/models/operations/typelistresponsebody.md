# TypeListResponseBody

OK


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Count`                                            | **int64*                                           | :heavy_minus_sign:                                 | The total number of types.                         | 3                                                  |
| `Next`                                             | **string*                                          | :heavy_minus_sign:                                 | URL to retrieve the next page of types.            | https://pokeapi.co/api/v2/type/?offset=20&limit=20 |
| `Previous`                                         | **string*                                          | :heavy_minus_sign:                                 | URL to retrieve the previous page of types.        |                                                    |
| `Results`                                          | [][shared.Type](../../models/shared/type.md)       | :heavy_minus_sign:                                 | N/A                                                |                                                    |