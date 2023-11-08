# ItemCategoryListResponseBody

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Count`                                                      | **int64*                                                     | :heavy_minus_sign:                                           | The total number of item categories.                         | 3                                                            |
| `Next`                                                       | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the next page of item categories.            | https://pokeapi.co/api/v2/item-category/?offset=20&limit=20  |
| `Previous`                                                   | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the previous page of item categories.        |                                                              |
| `Results`                                                    | [][shared.ItemCategory](../../models/shared/itemcategory.md) | :heavy_minus_sign:                                           | N/A                                                          |                                                              |