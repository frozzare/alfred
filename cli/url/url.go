package url

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/app"
	"github.com/frozzare/alfred/cli/root"
	"github.com/pkg/browser"
	"github.com/tj/go/clipboard"
)

func init() {
	cmd := root.Command("url", "Show, open or copy the application container url")

	open := cmd.Flag("open", "Open endpoint in the browser.").Short('o').Bool()
	copy := cmd.Flag("copy", "Copy endpoint to the clipboard.").Short('c').Bool()

	cmd.Action(func(_ *kingpin.ParseContext) error {
		c, err := root.Init()
		if err != nil {
			return err
		}

		app := app.NewApp(&app.Options{
			Config: c,
		})

		switch {
		case *open:
			browser.OpenURL(app.URL())
		case *copy:
			clipboard.Write(app.URL())
			fmt.Println("Copied to clipboard")
		default:
			fmt.Println(app.URL())
		}

		return nil
	})
}
