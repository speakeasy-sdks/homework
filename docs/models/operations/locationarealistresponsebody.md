# LocationAreaListResponseBody

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Count`                                                      | **int64*                                                     | :heavy_minus_sign:                                           | The total number of location areas.                          | 3                                                            |
| `Next`                                                       | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the next page of location areas.             | https://pokeapi.co/api/v2/location-area/?offset=20&limit=20  |
| `Previous`                                                   | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the previous page of location areas.         |                                                              |
| `Results`                                                    | [][shared.LocationArea](../../models/shared/locationarea.md) | :heavy_minus_sign:                                           | N/A                                                          |                                                              |