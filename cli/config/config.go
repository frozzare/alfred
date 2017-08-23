package config

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/cli/root"
)

func init() {
	cmd := root.Command("config", "Show application config")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		c, err := root.Init()
		if err != nil {
			return err
		}

		j, err := json.Marshal(c)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err := json.Indent(&out, j, "", "  "); err != nil {
			return err
		}

		fmt.Println("This is the real config after Alfred prepared it and it maybe not the same as you're alfred.json file")
		fmt.Printf("\n%s\n", out.Bytes())

		return nil
	})
}
