# VersionGroupList200ApplicationJSON

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Count`                                                      | **int64*                                                     | :heavy_minus_sign:                                           | The total number of version groups.                          | 3                                                            |
| `Next`                                                       | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the next page of version groups.             | https://pokeapi.co/api/v2/version-group/?offset=20&limit=20  |
| `Previous`                                                   | **string*                                                    | :heavy_minus_sign:                                           | URL to retrieve the previous page of version groups.         |                                                              |
| `Results`                                                    | [][shared.VersionGroup](../../models/shared/versiongroup.md) | :heavy_minus_sign:                                           | N/A                                                          |                                                              |