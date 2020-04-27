# assembled-go


[![GoDoc](https://godoc.org/github.com/assembledhq/assembled-go?status.svg)](https://pkg.go.dev/github.com/assembledhq/assembled-go)
[![Riza](https://riza.io/a/riza-generated-badge.svg)](https://riza.io/)

The official [Assembled](https://www.assembled.com/) Go client library.

## Installation

```sh
go get -u github.com/assembledhq/assembled-go
```

Then, import it using:

``` go
import (
    "github.com/assembledhq/assembled-go"
)
```

## Usage

See full documentation [here](https://pkg.go.dev/github.com/assembledhq/assembled-go).

```go
package main

import (
    "context"
    "fmt"

    "github.com/assembledhq/assembled-go"
)


func main() {
    ctx := context.Background()
    client := assembled.NewClient("<api_key>")

    resp, err := client.ListAgents(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%v\n", resp)
}
```

## Request latency telemetry

By default, this package sends request latency telemetry back to Assembled.
These numbers help Assembled improve the API for everyone.

You can disable this behavior if you prefer:

```go
client := assembled.NewClient("<api_key>")
client.EnableTelemetry = false
```
