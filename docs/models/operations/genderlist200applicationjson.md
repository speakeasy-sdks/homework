# GenderList200ApplicationJSON

OK


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Count`                                              | **int64*                                             | :heavy_minus_sign:                                   | The total number of genders.                         | 3                                                    |
| `Next`                                               | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the next page of genders.            | https://pokeapi.co/api/v2/gender/?offset=20&limit=20 |
| `Previous`                                           | **string*                                            | :heavy_minus_sign:                                   | URL to retrieve the previous page of genders.        |                                                      |
| `Results`                                            | [][shared.Gender](../../models/shared/gender.md)     | :heavy_minus_sign:                                   | N/A                                                  |                                                      |