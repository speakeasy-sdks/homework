# ItemPocketListResponseBody

OK


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Count`                                                         | **int64*                                                        | :heavy_minus_sign:                                              | The total number of item pockets.                               | 3                                                               |
| `Next`                                                          | **string*                                                       | :heavy_minus_sign:                                              | URL to retrieve the next page of item pockets.                  | https://pokeapi.co/api/v2/item-pocket/?offset=20&limit=20       |
| `Previous`                                                      | **string*                                                       | :heavy_minus_sign:                                              | URL to retrieve the previous page of item pockets.              |                                                                 |
| `Results`                                                       | [][shared.ItemPocket](../../../pkg/models/shared/itempocket.md) | :heavy_minus_sign:                                              | N/A                                                             |                                                                 |