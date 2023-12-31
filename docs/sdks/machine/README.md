# Machine

### Available Operations

* [MachineList](#machinelist)
* [MachineRead](#machineread)

## MachineList

### Example Usage

```go
package main

import(
	"context"
	"log"
	"homework"
	"homework/pkg/models/operations"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.Machine.MachineList(ctx, operations.MachineListRequest{
        Limit: homework.Int64(943749),
        Offset: homework.Int64(902599),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MachineList200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.MachineListRequest](../../models/operations/machinelistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.MachineListResponse](../../models/operations/machinelistresponse.md), error**


## MachineRead

### Example Usage

```go
package main

import(
	"context"
	"log"
	"homework"
	"homework/pkg/models/operations"
)

func main() {
    s := homework.New()

    ctx := context.Background()
    res, err := s.Machine.MachineRead(ctx, operations.MachineReadRequest{
        ID: 681820,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Machine != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.MachineReadRequest](../../models/operations/machinereadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.MachineReadResponse](../../models/operations/machinereadresponse.md), error**

