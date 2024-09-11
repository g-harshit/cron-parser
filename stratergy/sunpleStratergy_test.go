package stratergy

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		start, end int
		wantError  bool
		want       []int
	}{
		{
			name:  "all",
			input: "*",
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			start: 1,
			end:   31,
		},
		{
			name:  "all stepup 4",
			input: "*/4",
			want:  []int{1, 5, 9, 13, 17, 21, 25, 29},
			start: 1,
			end:   31,
		},
		{
			name:  "range",
			input: "0-4",
			want:  []int{1, 2, 3, 4},
			start: 1,
			end:   31,
		},
		{
			name:  "constant value",
			input: "0-4",
			want:  []int{1, 2, 3, 4},
			start: 1,
			end:   31,
		},
		{
			name:  "specific value",
			input: "1,2,5/2",
			want:  []int{2, 5, 7, 9},
			start: 2,
			end:   10,
		},
		{
			name:      "invalid",
			input:     "asd",
			want:      nil,
			wantError: true,
			start:     2,
			end:       10,
		},
		{
			name:  "single value",
			input: "9",
			want:  []int{9},
			start: 2,
			end:   10,
		},
	}
	p := NewSimpleServiceStatergy()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.input, tt.start, tt.end)
			if tt.wantError && err == nil {
				t.Errorf("want error")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayOfMonthParse : \ngot %v \nwant %v", got, tt.want)
			}
		})
	}
}
