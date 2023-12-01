# ItemAttributeListResponseBody

OK


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Count`                                                               | **int64*                                                              | :heavy_minus_sign:                                                    | The total number of item attributes.                                  | 3                                                                     |
| `Next`                                                                | **string*                                                             | :heavy_minus_sign:                                                    | URL to retrieve the next page of item attributes.                     | https://pokeapi.co/api/v2/item-attribute/?offset=20&limit=20          |
| `Previous`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | URL to retrieve the previous page of item attributes.                 |                                                                       |
| `Results`                                                             | [][shared.ItemAttribute](../../../pkg/models/shared/itemattribute.md) | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |