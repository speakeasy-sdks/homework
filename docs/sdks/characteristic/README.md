# Characteristic

### Available Operations

* [CharacteristicList](#characteristiclist)
* [CharacteristicRead](#characteristicread)

## CharacteristicList

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
    res, err := s.Characteristic.CharacteristicList(ctx, operations.CharacteristicListRequest{
        Limit: homework.Int64(963663),
        Offset: homework.Int64(272656),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Characteristics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CharacteristicListRequest](../../models/operations/characteristiclistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CharacteristicListResponse](../../models/operations/characteristiclistresponse.md), error**


## CharacteristicRead

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
    res, err := s.Characteristic.CharacteristicRead(ctx, operations.CharacteristicReadRequest{
        ID: 383441,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Characteristic != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CharacteristicReadRequest](../../models/operations/characteristicreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CharacteristicReadResponse](../../models/operations/characteristicreadresponse.md), error**

