# Hobocode

Simply a very small wrapper for [go-pretty](https://github.com/jedib0t/go-pretty) for different opinionated colored print calls like Info() and Error()

This is meant to be a very simple utility for CLI applications relying on UX rather than structured log outputs.

# Example

```
import (
    "github.com/asciifaceman/hobocode"
)

func main() {
    userinput := hobocode.Input("nvim", "What is your preferred editor?")

    hobocode.Notef("You chose: %s", userinput)
}
```

![example](static/hobocode_example.png)