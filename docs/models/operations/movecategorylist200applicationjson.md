# MoveCategoryList200ApplicationJSON

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Count`                                                      | **int64*                                                     | :heavy_minus_sign:                                           | The total number of move categories.                         | 3                                                            |
| `Next`                                                       | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the next page of move categories.            | https://pokeapi.co/api/v2/move-category/?offset=20&limit=20  |
| `Previous`                                                   | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the previous page of move categories.        |                                                              |
| `Results`                                                    | [][shared.MoveCategory](../../models/shared/movecategory.md) | :heavy_minus_sign:                                           | N/A                                                          |                                                              |