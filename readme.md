# musx

> Manage multiplexed-terminals with style

*musx* helps you quickly select the terminal-of-interest:

```bash
musx                 # create the default terminal or attach if only multiplexed-term
musx                 # select terminal-of-interest from list of terms using fuzzy-selection
musx bugfix          # create or open the terminal named 'bugfix'
musx -d ~/workspace  # fuzzy-select terminal-of-interest from list of terms based in ~/workspace
```

> Pro tip: type even less by using the alias `m` for `musx`.

# Install

```bash
go get github.com/ericcrosson/musx
```

# Usage

```
Usage:
musx [-d=<directory>] [name]
musx -h | --help
musx -v | --version

Options:
-d=<directory>  Filter terminals by directory [default: $HOME].
-h --help       Show this screen.
-v --version    Show version.
```
