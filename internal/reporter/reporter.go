package reporter

import (
	"fmt"
	"io"
)

func Print(w io.Writer, counts map[string]int) error {
	for name, count := range counts {
		if _, err := fmt.Fprintf(w, "%s:%d\n", name, count); err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}

	return nil
}
