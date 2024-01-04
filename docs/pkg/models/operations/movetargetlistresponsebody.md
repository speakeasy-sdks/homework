# MoveTargetListResponseBody

OK


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Count`                                                         | **int64*                                                        | :heavy_minus_sign:                                              | The total number of move targets.                               | 3                                                               |
| `Next`                                                          | **string*                                                       | :heavy_minus_sign:                                              | URL to retrieve the next page of move targets.                  | https://pokeapi.co/api/v2/move-target/?offset=20&limit=20       |
| `Previous`                                                      | **string*                                                       | :heavy_minus_sign:                                              | URL to retrieve the previous page of move targets.              |                                                                 |
| `Results`                                                       | [][shared.MoveTarget](../../../pkg/models/shared/movetarget.md) | :heavy_minus_sign:                                              | N/A                                                             |                                                                 |