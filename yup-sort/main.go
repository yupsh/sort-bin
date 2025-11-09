package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

const (
	flagReverse             = "reverse"
	flagNumeric             = "numeric-sort"
	flagUnique              = "unique"
	flagIgnoreCase          = "ignore-case"
	flagField               = "key"
	flagDelimiter           = "field-separator"
	flagRandom              = "random-sort"
	flagIgnoreLeadingBlanks = "ignore-leading-blanks"
	flagVersionSort         = "version-sort"
	flagHumanNumeric        = "human-numeric-sort"
	flagMonthSort           = "month-sort"
	flagStableSort          = "stable"
)

func main() {
	app := &cli.App{
		Name:  "sort",
		Usage: "sort lines of text files",
		UsageText: `sort [OPTIONS] [FILE...]

   Write sorted concatenation of all FILE(s) to standard output.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagReverse,
				Aliases: []string{"r"},
				Usage:   "reverse the result of comparisons",
			},
			&cli.BoolFlag{
				Name:    flagNumeric,
				Aliases: []string{"n"},
				Usage:   "compare according to string numerical value",
			},
			&cli.BoolFlag{
				Name:    flagUnique,
				Aliases: []string{"u"},
				Usage:   "output only the first of an equal run",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreCase,
				Aliases: []string{"f"},
				Usage:   "fold lower case to upper case characters",
			},
			&cli.IntFlag{
				Name:    flagField,
				Aliases: []string{"k"},
				Usage:   "sort via a key; KEYDEF gives location and type",
			},
			&cli.StringFlag{
				Name:    flagDelimiter,
				Aliases: []string{"t"},
				Usage:   "use SEP instead of non-blank to blank transition",
			},
			&cli.BoolFlag{
				Name:    flagRandom,
				Aliases: []string{"R"},
				Usage:   "shuffle, but group identical keys",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreLeadingBlanks,
				Aliases: []string{"b"},
				Usage:   "ignore leading blanks",
			},
			&cli.BoolFlag{
				Name:    flagVersionSort,
				Aliases: []string{"V"},
				Usage:   "natural sort of (version) numbers within text",
			},
		&cli.BoolFlag{
			Name:  flagHumanNumeric,
			Usage: "compare human readable numbers (e.g., 2K 1G)",
		},
			&cli.BoolFlag{
				Name:    flagMonthSort,
				Aliases: []string{"M"},
				Usage:   "compare (unknown) < 'JAN' < ... < 'DEC'",
			},
			&cli.BoolFlag{
				Name:    flagStableSort,
				Aliases: []string{"s"},
				Usage:   "stabilize sort by disabling last-resort comparison",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "sort: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagReverse) {
		params = append(params, Reverse)
	}
	if c.Bool(flagNumeric) {
		params = append(params, Numeric)
	}
	if c.Bool(flagUnique) {
		params = append(params, Unique)
	}
	if c.Bool(flagIgnoreCase) {
		params = append(params, IgnoreCase)
	}
	if c.IsSet(flagField) {
		params = append(params, Field(c.Int(flagField)))
	}
	if c.IsSet(flagDelimiter) {
		params = append(params, Delimiter(c.String(flagDelimiter)))
	}
	if c.Bool(flagRandom) {
		params = append(params, Random)
	}
	if c.Bool(flagIgnoreLeadingBlanks) {
		params = append(params, IgnoreLeadingBlanks)
	}
	if c.Bool(flagVersionSort) {
		params = append(params, VersionSort)
	}
	if c.Bool(flagHumanNumeric) {
		params = append(params, HumanNumeric)
	}
	if c.Bool(flagMonthSort) {
		params = append(params, MonthSort)
	}
	if c.Bool(flagStableSort) {
		params = append(params, StableSort)
	}

	// Create and execute the sort command
	cmd := Sort(params...)
	return gloo.Run(cmd)
}
