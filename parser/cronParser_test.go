package parser

import (
	"testing"
)

func TestNewCronParser(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{
			name:      "no command",
			input:     "2-5,7/3 1 * * 0-4",
			wantError: true,
		},
		{
			name:      "valid",
			input:     "2-5,7/3 1 * * 0-4 test-command",
			wantError: false,
		},
		{
			name:      "extra param",
			input:     "2-5,7/3 1 * * 0-4 test-command extra",
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCronParser(tt.input)
			if tt.wantError && err == nil {
				t.Errorf("want error")
			}
		})
	}
}
