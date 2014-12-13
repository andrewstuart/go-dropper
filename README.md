# Go-Ocean (or ocean)

[Main Documentation](https://godoc.org/github.com/andrewstuart/go-dropper/ocean)

A simple CLI & go library for DO's api

Everything should be mostly stable as far as the API. I'll just be adding functionality with the same
general API contracts.

# Example

[Playground](https://play.golang.org/p/7QKLMBD_QB)

```go
package main

import (
  "log"

  "github.com/andrewstuart/go-dropper/ocean"
)

c := ocean.NewClient("0123456789abcdef")

imgs, err := c.GetImages()

if err != nil {
  //Handle err
}

for _, i := range imgs {
  d := ocean.Droplet{
    Name: "My-" + d.Slug,
    Image: d.Slug,
    Size: "512mb",
    Region: "sfo1",
  }

  c.CreateDroplet(d)

  log.Printf("Droplet ID for droplet %s: %d", d.Name, d.Id)
}
```
