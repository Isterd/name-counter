package reader

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CountNames(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(r)
	lineNum := 0

	for scanner.Scan() {
		lineNum++

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		counts[line]++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read error near line %d: %w", lineNum, err)
	}

	return counts, nil
}
