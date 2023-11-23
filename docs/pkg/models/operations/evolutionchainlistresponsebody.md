# EvolutionChainListResponseBody

OK


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Count`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | The total number of evolution chains.                                   | 3                                                                       |
| `Next`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | URL to retrieve the next page of evolution chains.                      | https://pokeapi.co/api/v2/evolution-chain/?offset=20&limit=20           |
| `Previous`                                                              | **string*                                                               | :heavy_minus_sign:                                                      | URL to retrieve the previous page of evolution chains.                  |                                                                         |
| `Results`                                                               | [][shared.EvolutionChain](../../../pkg/models/shared/evolutionchain.md) | :heavy_minus_sign:                                                      | N/A                                                                     |                                                                         |