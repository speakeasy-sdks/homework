# LocationListResponseBody

OK


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Count`                                                     | **int64*                                                    | :heavy_minus_sign:                                          | The total number of locations.                              | 3                                                           |
| `Next`                                                      | **string*                                                   | :heavy_minus_sign:                                          | URL to retrieve the next page of locations.                 | https://pokeapi.co/api/v2/location/?offset=20&limit=20      |
| `Previous`                                                  | **string*                                                   | :heavy_minus_sign:                                          | URL to retrieve the previous page of locations.             |                                                             |
| `Results`                                                   | [][shared.Location](../../../pkg/models/shared/location.md) | :heavy_minus_sign:                                          | N/A                                                         |                                                             |