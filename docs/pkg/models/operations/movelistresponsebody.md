# MoveListResponseBody

OK


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Count`                                                       | **int64*                                                      | :heavy_minus_sign:                                            | The total number of moves.                                    | 3                                                             |
| `Next`                                                        | **string*                                                     | :heavy_minus_sign:                                            | URL to retrieve the next page of moves.                       | https://pokeapi.co/api/v2/move/?offset=20&limit=20            |
| `Previous`                                                    | **string*                                                     | :heavy_minus_sign:                                            | URL to retrieve the previous page of moves.                   |                                                               |
| `Results`                                                     | [][shared.MoveInput](../../../pkg/models/shared/moveinput.md) | :heavy_minus_sign:                                            | N/A                                                           |                                                               |