# VersionListResponseBody

OK


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Count`                                               | **int64*                                              | :heavy_minus_sign:                                    | The total number of versions.                         | 3                                                     |
| `Next`                                                | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the next page of versions.            | https://pokeapi.co/api/v2/version/?offset=20&limit=20 |
| `Previous`                                            | **string*                                             | :heavy_minus_sign:                                    | URL to retrieve the previous page of versions.        |                                                       |
| `Results`                                             | [][shared.Version](../../models/shared/version.md)    | :heavy_minus_sign:                                    | N/A                                                   |                                                       |