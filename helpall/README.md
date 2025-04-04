# helpall

`helpall` is a Go Package to show the help of all commands of CLIs built with [urfave/cli/v3](https://pkg.go.dev/github.com/urfave/cli/v3).
This is useful if you want to put the usage of CLI built with urfave/cli/v3 into the document.

## How To Use

[Example](../cmd/hello/main.go)

Using this library, you can add a command `help-all` showing the help of all commands.

```go
import (
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/helpall"
	"github.com/urfave/cli/v3"
)

// helpall.With appends the "help-all" command to the given Command.
helpall.With(&cli.Command{
	Commands: []*cli.Command{
		// ...
	},
}, nil).Run(context.Background(). os.Args)
```

`help-all` command outputs the help message.
You can put it into the document.

```console
$ go run ./cmd/hello help-all > hello.md
```

Example: [hello.md](../hello.md)

### Customize the command

The function `helpall.New()` returns a `*cli.Command`. You can customize the returned value.

e.g. Change the command name

```go
rootCmd := &cli.Command{
	Commands: []*cli.Command{
		// ...
	},
}
cmd := helpall.New(rootCmd, nil)
cmd.Name = "help-markdown" // Change the command name
rootCmd.Commands = append(app.Commands, cmd)
```

By default, the command is hidden, so it isn't shown in the help message.
You can show the command by changing the `Hidden` field.

```go
cmd := helpall.New(rootCmd, nil)
cmd.Hidden = true // Show the help of help-all
rootCmd.Commands = append(app.Commands, cmd)
```
