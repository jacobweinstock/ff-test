package one

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

type Data struct {
	Name     string
	Location string
	Debug    bool
}

func FlagSet(d *Data) *flag.FlagSet {
	fs := flag.NewFlagSet("one", flag.ExitOnError)
	fs.StringVar(&d.Name, "name", "", "name to be used")
	fs.StringVar(&d.Location, "loc", "here", "location to be used")
	fs.BoolVar(&d.Debug, "debug", false, "use debugging")

	return fs
}

func FFcli(d *Data) *ffcli.Command {
	fs := FlagSet(d)
	return &ffcli.Command{
		Name:    "one",
		FlagSet: fs,
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println("one")
			return nil
		},
	}
}
