package two

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

type Data struct {
	Name   string
	Number int
}

func FlagSet(d *Data) *flag.FlagSet {
	fs := flag.NewFlagSet("two", flag.ExitOnError)
	fs.StringVar(&d.Name, "name", "", "name to be used")
	fs.IntVar(&d.Number, "number", 0, "number to be used")
	return fs
}

func FFcli(d *Data) *ffcli.Command {
	fs := FlagSet(d)

	return &ffcli.Command{
		Name:    "two",
		FlagSet: fs,
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println("two")
			return nil
		},
	}
}
