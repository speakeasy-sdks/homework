# StatListResponseBody

OK


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Count`                                             | **int64*                                            | :heavy_minus_sign:                                  | The total number of stats.                          | 3                                                   |
| `Next`                                              | **string*                                           | :heavy_minus_sign:                                  | URL to retrieve the next page of stats.             | https://pokeapi.co/api/v2/stat/?offset=20&limit=20  |
| `Previous`                                          | **string*                                           | :heavy_minus_sign:                                  | URL to retrieve the previous page of stats.         |                                                     |
| `Results`                                           | [][shared.Stat](../../../pkg/models/shared/stat.md) | :heavy_minus_sign:                                  | N/A                                                 |                                                     |