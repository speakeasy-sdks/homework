# ItemList200ApplicationJSON

OK


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Count`                                            | **int64*                                           | :heavy_minus_sign:                                 | The total number of items.                         | 3                                                  |
| `Next`                                             | **string*                                          | :heavy_minus_sign:                                 | URL to retrieve the next page of items.            | https://pokeapi.co/api/v2/item/?offset=20&limit=20 |
| `Previous`                                         | **string*                                          | :heavy_minus_sign:                                 | URL to retrieve the previous page of items.        |                                                    |
| `Results`                                          | [][shared.Item](../../models/shared/item.md)       | :heavy_minus_sign:                                 | N/A                                                |                                                    |