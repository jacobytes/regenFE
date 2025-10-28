# regenfe

> Client for getting Formula E race data from https://fiaformulae.com/en/stats-centre.

## Features

- 🏆 Get championship data
- 🏎️ Get driver data
- 🛞 Get race data
- 🏁 Get team data

## Installation

Use `go get` to install the package:

```bash
go get github.com/jacobytes/regenfe
```

## Usage

```
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jacobytes/regenfe"
)

func main() {
	client := regenfe.NewClient()

	ctx := context.Background()
	options := regenfe.ListOptions{
		Page: 1,
	}

	results, _, _ := client.Races.ListRaces(ctx, options)
	json, _ := json.Marshal(results)
	fmt.Println(string(json))
}
```

## Tests

Run all tests:

```
go test ./...

```

## License

This project is licensed under the MIT License — see the LICENSE file for details.
