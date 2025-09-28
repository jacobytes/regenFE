# regenFE 

> Client for getting Formula E race data from https://fiaformulae.com/en/stats-centre.

## Features

- 🏆 Get championship data
- 🏎️ Get driver data
- 🛞 Get race data
- 🏁 Get team data

## Installation

Use `go get` to install the package:

```bash
go get github.com/jacobytes/regenFE

```

## Usage

```
package main

import (
    "fmt"
    "github.com/jacobytes/regenFE"
)

func main() {
    client := NewClient()
	ctx := context.Background()
	options := ListOptions{
		Page: 1,
	}

	results, _, err := client.Races.ListRaces(ctx, options) 
}
```

## Tests

Run all tests:

```
go test ./...

```

## License

This project is licensed under the MIT License — see the LICENSE file for details.
