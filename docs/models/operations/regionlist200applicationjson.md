# RegionList200ApplicationJSON

OK


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Count`                                              | **int64*                                             | :heavy_minus_sign:                                   | The total number of regions.                         | 3                                                    |
| `Next`                                               | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the next page of regions.            | https://pokeapi.co/api/v2/region/?offset=20&limit=20 |
| `Previous`                                           | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the previous page of regions.        |                                                      |
| `Results`                                            | [][shared.Region](../../models/shared/region.md)     | :heavy_minus_sign:                                   | N/A                                                  |                                                      |