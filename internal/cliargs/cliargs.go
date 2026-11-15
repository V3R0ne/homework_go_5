package cliargs

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ErrUsage marks invalid command-line arguments.
var ErrUsage = errors.New("usage: args <name> [repeat]")

// Options contains validated user arguments.
type Options struct {
	Name   string
	Repeat int
}

// UserArgs returns a new slice without os.Args[0].
func UserArgs(all []string) []string {
	// TODO: implement according to docs/task.md.
	if len(all) == 0 {
		return []string{}
	}
	args := make([]string, len(all)-1)
	copy(args, all[1:])
	return args
}

// Parse validates the arguments passed after os.Args[0].
func Parse(args []string) (Options, error) {
	// TODO: implement according to docs/task.md.
	if len(args) != 1 && len(args) != 2 {
		return Options{}, ErrUsage
	}
	name := strings.TrimSpace(args[0])
	if len(name) == 0 {
		return Options{}, ErrUsage
	}
	if len(args) == 1 {
		return Options{
			Name:   name,
			Repeat: 1,
		}, nil
	}
	if rep, err := strconv.Atoi(args[1]); err == nil && rep > 0 {
		return Options{
			Name:   name,
			Repeat: rep,
		}, nil
	}
	return Options{}, ErrUsage
}

// Run writes the command result to the supplied writer.
func Run(args []string, out io.Writer) error {
	// TODO: implement according to docs/task.md.
	opts, err := Parse(args)
	if err != nil {
		return err
	}
	for i := 0; i < opts.Repeat; i++ {
		_, err := fmt.Fprintf(out, "hello, %s\n", opts.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

// Example returns deterministic output used by cmd and the example test.
func Example() string {
	var out strings.Builder
	_ = Run([]string{"Maria", "2"}, &out)
	return out.String()
}
