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
	}

	return make([]workflows.Item, 0)
}
