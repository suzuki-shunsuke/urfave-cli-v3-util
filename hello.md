```console
$ hello --help
NAME:
   hello - A new cli application

USAGE:
   hello [global options] [command [command options]]

COMMANDS:
   foo      foo command
   bar      bar command
   version  Show version
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## hello foo

```console
$ hello foo --help
NAME:
   hello foo - foo command

USAGE:
   hello foo

DESCRIPTION:
   This is a foo command

OPTIONS:
   --help, -h  show help
```

## hello bar

```console
$ hello bar --help
NAME:
   hello bar - bar command

USAGE:
   hello bar

DESCRIPTION:
   This is a bar command

OPTIONS:
   --help, -h  show help
```

## hello version

```console
$ hello version --help
NAME:
   hello version - Show version

USAGE:
   hello version

OPTIONS:
   --json, -j  Output version in JSON format (default: false)
   --help, -h  show help
```
