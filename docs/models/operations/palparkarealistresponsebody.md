# PalParkAreaListResponseBody

OK


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Count`                                                     | **int64*                                                    | :heavy_minus_sign:                                          | The total number of pal park areas.                         | 3                                                           |
| `Next`                                                      | **string*                                                   | :heavy_minus_sign:                                          | URL to retrieve the next page of pal park areas.            | https://pokeapi.co/api/v2/pal-park-area/?offset=20&limit=20 |
| `Previous`                                                  | **string*                                                   | :heavy_minus_sign:                                          | URL to retrieve the previous page of pal park areas.        |                                                             |
| `Results`                                                   | [][shared.PalParkArea](../../models/shared/palparkarea.md)  | :heavy_minus_sign:                                          | N/A                                                         |                                                             |