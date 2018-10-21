# ginger
Core utilities to create a gin app with config, routing, db(qb) and more

# Example
```go
package main

import(
    "github.com/slicebit/ginger"
)

type Config struct {
    Host string
    Port int
    Env string
}

func main() {

    config := &Config{}
    err := ginger.NewConfig("config.json", config)
    if err != nil {
        panic(err)
    }

    app, err := ginger.NewApp(config.Host, config.Port, config.Env)
    if err != nil {
        panic(err)
    }

    // Start the app
    app.Start()
}

```