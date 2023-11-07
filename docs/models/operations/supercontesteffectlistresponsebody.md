# SuperContestEffectListResponseBody

OK


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Count`                                                                  | **int64*                                                                 | :heavy_minus_sign:                                                       | The total number of super contest effects.                               | 3                                                                        |
| `Next`                                                                   | **string*                                                                | :heavy_minus_sign:                                                       | URL to retrieve the next page of super contest effects.                  | https://pokeapi.co/api/v2/super-contest-effect/?offset=20&limit=20       |
| `Previous`                                                               | **string*                                                                | :heavy_minus_sign:                                                       | URL to retrieve the previous page of super contest effects.              |                                                                          |
| `Results`                                                                | [][shared.SuperContestEffect](../../models/shared/supercontesteffect.md) | :heavy_minus_sign:                                                       | N/A                                                                      |                                                                          |