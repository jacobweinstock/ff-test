package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jacobweinstock/flagset/one"
	"github.com/jacobweinstock/flagset/two"
)

type both struct {
	one one.Data
	two two.Data
}

func main() {
	d := &both{
		one: one.Data{},
		two: two.Data{},
	}

	fs1 := one.FlagSet(&d.one)
	fs2 := two.FlagSet(&d.two)

	flags := os.Args[1:]

	oneFlags := parseArgs(fs1, flags)
	twoFlags := parseArgs(fs2, flags)
	fmt.Println(oneFlags)
	fmt.Println(twoFlags)
	fs1.Parse(oneFlags)
	fs2.Parse(twoFlags)
	fmt.Printf("%+v\n", *d)


	
}

func parseArgs(fs *flag.FlagSet, allArgs []string) []string {
	var result []string
	for index, elem := range allArgs {
		if index%2 == 0 {
			if f := fs.Lookup(strings.TrimLeft(elem, "-")); f != nil && strings.HasPrefix(elem, "-") {
				if fv, ok := f.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
					result = append(result, allArgs[index], "true")
				} else {
					result = append(result, allArgs[index], allArgs[index+1])
				}
			}
		}
	}

	return result
}

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}
