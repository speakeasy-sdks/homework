# MachineListResponseBody

OK


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Count`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | The total number of machines.                             | 3                                                         |
| `Next`                                                    | **string*                                                 | :heavy_minus_sign:                                        | URL to retrieve the next page of machines.                | https://pokeapi.co/api/v2/machine/?offset=20&limit=20     |
| `Previous`                                                | **string*                                                 | :heavy_minus_sign:                                        | URL to retrieve the previous page of machines.            |                                                           |
| `Results`                                                 | [][shared.Machine](../../../pkg/models/shared/machine.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |