<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"homework"
	"homework/pkg/models/operations"
	"log"
)

func main() {
	s := homework.New()

	ctx := context.Background()
	res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->