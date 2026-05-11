package reader_test

import (
	"strings"
	"testing"

	"name-counter/internal/reader"
)

func TestCountNames(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    map[string]int
		wantErr bool
	}{
		{
			name:  "normal file",
			input: "Alyona\nMisha\nDima\n",
			want:  map[string]int{"Alyona": 1, "Misha": 1, "Dima": 1},
		},
		{
			name:  "counts duplicates",
			input: "Alyona\nMisha\nAlyona\n",
			want:  map[string]int{"Alyona": 2, "Misha": 1},
		},
		{
			name:  "empty lines are skipped",
			input: "Alyona\n\nMisha\n\n",
			want:  map[string]int{"Alyona": 1, "Misha": 1},
		},
		{
			name:  "whitespace-only lines are skipped",
			input: "Alyona\n   \nMisha",
			want:  map[string]int{"Alyona": 1, "Misha": 1},
		},
		{
			name:  "surrounding whitespace is trimmed",
			input: "  Alyona  \nMisha\n",
			want:  map[string]int{"Alyona": 1, "Misha": 1},
		},
		{
			name:  "empty file",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "no trailing newline",
			input: "Alyona\nMisha",
			want:  map[string]int{"Alyona": 1, "Misha": 1},
		},
		{
			name:  "case sensitive",
			input: "alyona\nAlyona\n",
			want:  map[string]int{"alyona": 1, "Alyona": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.CountNames(strings.NewReader(tt.input))

			if (err != nil) != tt.wantErr {
				t.Fatalf("CountNames() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if len(got) != len(tt.want) {
				t.Fatalf("CountNames() returned %d entries, expected %d\ngot:  %v\nwant: %v",
					len(got), len(tt.want), got, tt.want)
			}

			for name, wantCount := range tt.want {
				if got[name] != wantCount {
					t.Errorf("CountNames()[%q] = %d, want %d", name, got[name], wantCount)
				}
			}
		})
	}
}
