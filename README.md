# Serpentor

[Serpentor](https://en.wikipedia.org/wiki/Serpentor) is the [Cobra](https://github.com/spf13/cobra) Emporer.

It provides a concise way to create new commands as determined by rigorous subjectivity.

It is not meant to provide all the functionality `cobra.Command` does. 
You can get the underlying `cobra.Command` by calling `GetCobra()` to perform any actions you need to.

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
