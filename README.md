# Serpentor
![picture of serpentor at podium](https://en.wikipedia.org/wiki/Serpentor#/media/File:Gijoe23.jpg)
[Serpentor](https://en.wikipedia.org/wiki/Serpentor), Emperer of [Cobra](https://github.com/spf13/cobra).

This module provides a concise way to create new commands as determined by rigorous subjectivity.

Serpentor is not meant to provide all the functionality of `cobra.Command`, for anything missing use `GetCobra()` to get the underlying `cobra.Command`.

```go
package main

import "github.com/ncsnw/serpentor"

func main() {
    serpentor.NewCmd("example").
        Short("An example program using serpentor with cobra and viper").
        AddCommand(
            serpentor.NewCmd("config").
                Short("Commands for interacting with configuration.").
                AddCommand(
                    serpentor.NewCmd("show").
                        Short("Print the current config to stdout").
                        Run(showConfig),
                    serpentor.NewCmd("get [key]").
                        Short("Get the key's value from the config").
                        Run(getConfig),
                    serpentor.NewCmd("set key value").
                        Short("Set the config key to the specified value").
                        Run(setConfig),
                ),
        )
}
```
