# EncounterMethodListResponseBody

OK


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Count`                                                            | **int64*                                                           | :heavy_minus_sign:                                                 | The total number of encounter methods.                             | 3                                                                  |
| `Next`                                                             | **string*                                                          | :heavy_minus_sign:                                                 | URL to retrieve the next page of encounter methods.                | https://pokeapi.co/api/v2/encounter-method/?offset=20&limit=20     |
| `Previous`                                                         | **string*                                                          | :heavy_minus_sign:                                                 | URL to retrieve the previous page of encounter methods.            |                                                                    |
| `Results`                                                          | [][shared.EncounterMethod](../../models/shared/encountermethod.md) | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |