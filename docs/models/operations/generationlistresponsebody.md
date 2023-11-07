# GenerationListResponseBody

OK


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Count`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | The total number of generations.                         | 3                                                        |
| `Next`                                                   | **string*                                                | :heavy_minus_sign:                                       | URL to retrieve the next page of generations.            | https://pokeapi.co/api/v2/generation/?offset=20&limit=20 |
| `Previous`                                               | **string*                                                | :heavy_minus_sign:                                       | URL to retrieve the previous page of generations.        |                                                          |
| `Results`                                                | [][shared.Generation](../../models/shared/generation.md) | :heavy_minus_sign:                                       | N/A                                                      |                                                          |