# EcoAloha Go SDK

Official source: https://github.com/ArneFfm/ecoaloha-agents.
API documentation: https://ecoaloha.com/developers.
Go 1.22 or later is required. The SDK uses the standard library only.

```sh
go get github.com/ArneFfm/ecoaloha-agents/packages/sdk-go@v0.1.0
```

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	ecoaloha "github.com/ArneFfm/ecoaloha-agents/packages/sdk-go"
)

func main() {
	client, err := ecoaloha.NewClient("https://ecoaloha.com/api/sandbox/v1", &http.Client{Timeout: 15 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Experiences(context.Background(), url.Values{"destinationId": {"paris"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}
```

Use `Destinations(ctx)`, `Experiences(ctx, query)`, and `Compare(ctx, ids, currency)`.
An empty currency selects EUR. An empty base URL selects the production API.
An omitted HTTP client (`nil`) selects `http.DefaultClient`, which has no timeout.
Pass a context deadline or an HTTP client with a timeout.

Methods return `json.RawMessage` with the complete response envelope, including `nextCursor`.
Use `Request(ctx, path, RequestOptions{...})` for other documented operations.
Paths must be API-relative, such as `/destinations`. Query parameters belong in `RequestOptions.Query`.
Use `errors.As` to inspect `*ecoaloha.HTTPError` fields: `StatusCode`, `Body`, and `RetryAfter`.
The SDK does not retry failed requests.

The public discovery API is free and keyless. API rate limits still apply.
The SDK does not grant partner API access, content rights, or booking entitlements.
Affiliate links can earn EcoAloha a commission. Preserve required affiliate disclosures.
The SDK does not call partner APIs directly.

Run `go test ./...` from this directory.
