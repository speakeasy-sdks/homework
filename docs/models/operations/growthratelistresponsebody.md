# GrowthRateListResponseBody

OK


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Count`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | The total number of growth rates.                         | 3                                                         |
| `Next`                                                    | **string*                                                 | :heavy_minus_sign:                                        | URL to retrieve the next page of growth rates.            | https://pokeapi.co/api/v2/growth-rate/?offset=20&limit=20 |
| `Previous`                                                | **string*                                                 | :heavy_minus_sign:                                        | URL to retrieve the previous page of growth rates.        |                                                           |
| `Results`                                                 | [][shared.GrowthRate](../../models/shared/growthrate.md)  | :heavy_minus_sign:                                        | N/A                                                       |                                                           |