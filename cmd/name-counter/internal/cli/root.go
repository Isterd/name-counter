package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"name-counter/internal/reader"
	"name-counter/internal/reporter"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "name-counter <file>",
		Short: "Count occurrences of each name in a file",
		Long: `name-counter reads a file with one name per line and prints the count of each name.

Example:
  name-counter names.txt`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args[0])
		},
	}

	return cmd
}

func run(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file %q: %w", filename, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("could not close file: %w", cerr)
		}
	}()

	counts, err := reader.CountNames(f)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return reporter.Print(os.Stdout, counts)
}
