# ContestTypeListResponseBody

Successful response containing a list of contest types


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Count`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | The total number of contest types returned                        |
| `Next`                                                            | **string*                                                         | :heavy_minus_sign:                                                | URL to the next page of contest types, if any                     |
| `Previous`                                                        | **string*                                                         | :heavy_minus_sign:                                                | URL to the previous page of contest types, if any                 |
| `Results`                                                         | [][shared.ContestType](../../../pkg/models/shared/contesttype.md) | :heavy_minus_sign:                                                | N/A                                                               |