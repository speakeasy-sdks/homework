# LanguageList200ApplicationJSON

OK


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Count`                                                | **int64*                                               | :heavy_minus_sign:                                     | The total number of languages.                         | 3                                                      |
| `Next`                                                 | **string*                                              | :heavy_minus_sign:                                     | URL to retrieve the next page of languages.            | https://pokeapi.co/api/v2/language/?offset=20&limit=20 |
| `Previous`                                             | **string*                                              | :heavy_minus_sign:                                     | URL to retrieve the previous page of languages.        |                                                        |
| `Results`                                              | [][shared.Language](../../models/shared/language.md)   | :heavy_minus_sign:                                     | N/A                                                    |                                                        |