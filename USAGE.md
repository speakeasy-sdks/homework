<!-- Start SDK Example Usage -->
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
    res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{
        Limit: homework.Int64(548814),
        Offset: homework.Int64(592845),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AbilityList200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->