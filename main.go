package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kulikov/alfred-workflow/workflows"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "1.0"
	app     = kingpin.New("alfred-workflow", "Set of useful workflows for Alfred").Version(version)

	timestampConverter = app.Command("timestamp-converter", "Timestamp converter")
	timestampQuery     = timestampConverter.Arg("query", "Date or timestamp").Required().String()

	pwgen       = app.Command("pwgen", "Password generator")
	pwgenLength = pwgen.Arg("length", "Length").Default("24").Int()

	evmkeys = app.Command("evmkeys", "Generate Ethereum keys")
)

func main() {
	output := runCommand()

	result, _ := json.MarshalIndent(map[string]interface{}{"items": output}, "", "  ")

	fmt.Println(string(result))
}

func runCommand() []workflows.Item {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case timestampConverter.FullCommand():
		return workflows.ConvertTimestamp(*timestampQuery)

	case pwgen.FullCommand():
		return workflows.Pwgen(*pwgenLength)

	case evmkeys.FullCommand():
		return workflows.Evmkeys()
	}

	return make([]workflows.Item, 0)
}
