# ginger
Core utilities to create a gin app with config, routing, db(qb) and more

# Example
```go
package main

import(
    "github.com/slicebit/ginger"
)

func main() {

    config, err := ginger.NewConfig("config.toml")
    if err != nil {
        panic(err)
    }

    app, err := ginger.NewApp(config)
    if err != nil {
        panic(err)
    }

    // Start the app
    app.Start()
}

```