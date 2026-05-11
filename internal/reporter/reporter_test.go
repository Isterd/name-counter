package reporter_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"name-counter/internal/reporter"
)

func TestPrint(t *testing.T) {
	counts := map[string]int{"Alyona": 2, "Misha": 1}

	var buf bytes.Buffer
	if err := reporter.Print(&buf, counts); err != nil {
		t.Fatalf("Print() returned unexpected error: %v", err)
	}

	got := buf.String()
	for name, count := range counts {
		expected := fmt.Sprintf("%s:%d\n", name, count)
		if !strings.Contains(got, expected) {
			t.Errorf("Print() output missing %q\ngot: %s", expected, got)
		}
	}
}

func TestPrint_Empty(t *testing.T) {
	var buf bytes.Buffer
	if err := reporter.Print(&buf, map[string]int{}); err != nil {
		t.Fatalf("Print() returned unexpected error: %v", err)
	}

	if buf.String() != "" {
		t.Errorf("Print() on empty map should produce empty output, got: %q", buf.String())
	}
}
