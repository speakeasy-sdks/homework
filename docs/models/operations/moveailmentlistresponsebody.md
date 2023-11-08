# MoveAilmentListResponseBody

OK


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Count`                                                    | **int64*                                                   | :heavy_minus_sign:                                         | The total number of move ailments.                         | 3                                                          |
| `Next`                                                     | **string*                                                  | :heavy_minus_sign:                                         | URL to retrieve the next page of move ailments.            | https://pokeapi.co/api/v2/move-ailment/?offset=20&limit=20 |
| `Previous`                                                 | **string*                                                  | :heavy_minus_sign:                                         | URL to retrieve the previous page of move ailments.        |                                                            |
| `Results`                                                  | [][shared.MoveAilment](../../models/shared/moveailment.md) | :heavy_minus_sign:                                         | N/A                                                        |                                                            |