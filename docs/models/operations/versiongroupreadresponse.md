# VersionGroupReadResponse


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ContentType`                                               | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `StatusCode`                                                | *int*                                                       | :heavy_check_mark:                                          | N/A                                                         |
| `RawResponse`                                               | [*http.Response](https://pkg.go.dev/net/http#Response)      | :heavy_minus_sign:                                          | N/A                                                         |
| `VersionGroup`                                              | [*shared.VersionGroup](../../models/shared/versiongroup.md) | :heavy_minus_sign:                                          | Successful response                                         |