# NatureList200ApplicationJSON

OK


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Count`                                              | **int64*                                             | :heavy_minus_sign:                                   | The total number of natures.                         | 3                                                    |
| `Next`                                               | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the next page of natures.            | https://pokeapi.co/api/v2/nature/?offset=20&limit=20 |
| `Previous`                                           | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the previous page of natures.        |                                                      |
| `Results`                                            | [][shared.Nature](../../models/shared/nature.md)     | :heavy_minus_sign:                                   | N/A                                                  |                                                      |