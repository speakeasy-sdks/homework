# EncounterConditionList200ApplicationJSON

OK


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Count`                                                                  | **int64*                                                                 | :heavy_minus_sign:                                                       | The total number of encounter conditions.                                | 3                                                                        |
| `Next`                                                                   | **string*                                                                | :heavy_minus_sign:                                                       | URL to retrieve the next page of encounter conditions.                   | https://pokeapi.co/api/v2/encounter-condition/?offset=20&limit=20        |
| `Previous`                                                               | **string*                                                                | :heavy_minus_sign:                                                       | URL to retrieve the previous page of encounter conditions.               |                                                                          |
| `Results`                                                                | [][shared.EncounterCondition](../../models/shared/encountercondition.md) | :heavy_minus_sign:                                                       | N/A                                                                      |                                                                          |