package status

import (
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/frozzare/alfred/cli/root"
	"github.com/frozzare/alfred/internal/docker"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

func init() {
	cmd := root.Command("status", "Show application statuses")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Host",
			"Status",
			"Start Time",
		})
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		_, err := root.Init()
		if err != nil {
			return err
		}

		d, err := docker.NewDocker()
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		cos, err := d.FindContainers(map[string][]string{
			"label": []string{"alfred=true"},
		})
		if err != nil {
			return errors.Wrap(err, "Docker")
		}

		for _, co := range cos {
			table.Append([]string{
				strings.Replace(co.Name, "/", "", -1),
				strings.Title(co.State.StateString()),
				co.State.StartedAt.Format("2006-01-02 15:04:05"),
			})
		}

		table.Render()

		return nil
	})
}
