// Written by Eric Crosson
// 2017-05-20

package main

import (
	"os"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/codeskyblue/go-sh"
)

const DefaultMultiplexer = "screen"
const DefaultTerminalName = "mon"
const DefaultTerminalDirectory = "/home/eric"

type Multiplexer struct {
	Name string
	Create []string
	List []string
}

type MultiplexerCollection []Multiplexer

type Terminal struct {
	Name string
	Multiplexer Multiplexer
	Directory string
}

type TerminalList []Terminal

func getConfigFile() string {
	fmt.Println("TODO: implement getConfigFile")
	return os.Getenv("HOME")
}

func getConfig(key string) string {
	fmt.Println("TODO: implement getConfig")
	// file, err := os.Open(getConfigFile())
	return key
}

func programIsInstalled(program string) bool {
	_, err := sh.Command("which", program).Output()
	return err == nil
}

func createMultiplexer(name string) Multiplexer {
	var createCommand []string
	var listCommand []string
	switch name {
	case "screen":
		createCommand = append(createCommand, "screen", "-S")
		listCommand = append(listCommand, "screen", "-ls")
	case "tmux":
		panic("Not supported -- yet!")
	}
	return Multiplexer{name, createCommand, listCommand}
}

func getMultiplexers(defaultMultiplexer string) MultiplexerCollection {
	var multiplexers MultiplexerCollection
	if programIsInstalled(defaultMultiplexer) {
		multiplexers = append(multiplexers, createMultiplexer(defaultMultiplexer))
	}
	// fixme: add support for additional multiplexers
	return multiplexers
}

// RESUME: keep desining. terminal is useful? how about multiplexer?
func fzf(terminals TerminalList) Terminal {
	var terminal Terminal
	fmt.Println("TODO: implement fzf")
	return terminal
}

func (t TerminalList) FilterByDirectory(directory interface{}) TerminalList {
	var terminals TerminalList
	fmt.Println("TODO: implement FilterByDirectory")
	return terminals
}

func (t Terminal) Attach() {
	fmt.Println("TODO: implement attach")
}

func (t Terminal) Create() {
	fmt.Println("TODO: implement create")
}

func (t Terminal) AttachOrCreate() {
	fmt.Println("TODO: implement attachOrCreate(terminal)")
}

func (m Multiplexer) GetTerminals() TerminalList {
	var terminals TerminalList
	output, _ := sh.Command(m.Create).Output()
	return output
}

func (m MultiplexerCollection) Attach() {
	fmt.Println("TODO: implement attach(mc)")
}

func (m MultiplexerCollection) Create(terminalName string) {
	fmt.Println("TODO: implement create(mc)", terminalName)
	// fixme: convert to debug
	fmt.Println("Creating screen:", terminalName)
	fmt.Println()
}

func (m MultiplexerCollection) AttachOrCreate(terminalName interface{}, workingDirectory string) {
	fmt.Println("TODO: implement attachOrCreate(mc)")
}

func (m MultiplexerCollection) GetTerminals() TerminalList {
	var terminals TerminalList
	for _, element := range m {
		fmt.Println("M is", m)
		for _, terminal := range element.GetTerminals() {
			fmt.Println("Terminal is", terminal)
			terminals = append(terminals, terminal)
		}
	}
	fmt.Println("TODO: implement GetTerminals(mc)", terminals)
	return terminals
}

func main() {
	usage := `Multiplexed-Terminal Manager

Usage:
  musx [-d=<directory>] [name]
  musx -h | --help
  musx -v | --version

Options:
  -d=<directory>     Filter terminal-list by directory or specify
                     spawn directory for new terminals.
  -h --help          Show this screen.
  -v --version       Show version.`

	arguments, _ := docopt.Parse(usage, nil, true,
		"Multiplexed-Terminal Manager 1.0.0", false)

	if arguments["-m"] == false {
		// fixme: use conf files defaultMultiplexer field
		// arguments["-m"] = getConfig("defaultMultiplexer")
		arguments["-m"] = DefaultMultiplexer
	}

	multiplexers := getMultiplexers(DefaultMultiplexer)
	terminals := multiplexers.GetTerminals()

	fmt.Println(arguments)

	if arguments["name"] != false {
		// create terminal
		multiplexers.AttachOrCreate(arguments["name"], DefaultTerminalDirectory)
	} else {
		if arguments["-d"] != nil {
			// select terminal from filtered list
			fzf(multiplexers.GetTerminals().FilterByDirectory(arguments["-d"]))
		}  else {
			if len(terminals) > 1 {
				// select terminal from list
				fzf(multiplexers.GetTerminals())
			} else if len(terminals) == 1 {
				terminals[0].Attach()
			} else {
				// fixme: use conf files defaultTerminalName field
				multiplexers.Create(DefaultTerminalName)
			}
		}
	}
}
