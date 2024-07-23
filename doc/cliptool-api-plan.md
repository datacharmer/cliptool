# cliptool API planning

`cliptool` is an utility that can copy data from a file (or standard input) to the system clipboard, and inversely from the clipboard to the standard input or a given file.
Additionally, the tool allows to hold clipboard data into a holding area, and retrieve it at will.

## Commands:

Note: using commands is easy if we use a dedicated package such as [cobra](github.com/spf13/cobra), otherwise, it requires some infrastructure that would make the project unnecessarily complex.

### `load {filename|-}`

**(Alias: `copy`)**

Copies the content of a given filename to the clipboard. If the file is named "-", it reads from the standard input.

### `hold {filename|-} tag`

Copies the content of a filename (or "-" for standard input) into a holding area, identified by a tag.

### `get tag` 

**(Alias: `retrieve`, `fetch`)**

Retrieves the content of a wanted tag into the clipboard.

### `delete tag` 

**(Alias: `remove`)**

Removes the data associated with a given tag from the holding area.

### `clear all`

Empties the holding area.

### `unload [tag]`

**(Alias: `paste`)**

Empties the contents of the clipboard to standard output. If `tag` was requested, its contents will be retrieved instead of the clipboard.

### `help`

**(Alias: `usage`, `info`)**

Same as `--help`

### `version`

Same as `--version`

### `defaults`

**(Alias: `config`)**

Shows the current configurable options.

### `upgrade [version]`

Replaces the current executable with the most recent version released on GitHub (or a specific one if `version` was mentioned.)

### `api`

Shows all commands, subcommands and options currently available.

## Options and flags

### `-h`

Shows a succint usage screen.

### `--help`

**(Alias: `--usage`, `--info`)**

Shows a detailed usage screen

### `--quiet` 

**(Alias: `--silent`)**

Do not show any output.

### `--verbose`

**(Alias: `--debug`)**

Show more information during execution.

### `--output=fileName`

Changes the destination of the command to write to a given file instead of the regular output.

### `--log`

Logs all operations to the default log file

### `--logfile=name`

Changes the default log file (`cliptool.log`) to a custom one. (`stderr` and `stdout` accepted)

### `--version`

Show the version and exit.

### `--holding-area=name`

Changes the default holding area (`$HOME/.config/cliptool/holding`)

### `--config=name`

Changes the default configuration file (`$HOME/.config/cliptool/config.json`)
