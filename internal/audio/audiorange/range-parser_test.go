package audiorange

import (
	"reflect"
	"testing"
)

func TestRangeParser(t *testing.T) {
	tests := []struct {
		name     string
		rangeStr string
		want     []Range
		wantErr  bool
	}{
		{
			name:     "parsing header",
			rangeStr: "bytes=0-499",
			want: []Range{
				{
					Unit:  "bytes",
					Start: 0,
					End:   499,
				},
			},
			wantErr: false,
		},
		{
			name:     "parsing header without end",
			rangeStr: "bytes=499-",
			want: []Range{
				{
					Unit:  "bytes",
					Start: 499,
					End:   -1,
				},
			},
			wantErr: false,
		},
		{
			name:     "parsing header without start",
			rangeStr: "bytes=-1000",
			want: []Range{
				{
					Unit:  "bytes",
					Start: -1,
					End:   1000,
				},
			},
			wantErr: false,
		},
		{
			name:     "parsing multiple headers",
			rangeStr: "bytes=200-999, 2000-2499, 9500-",
			want: []Range{
				{
					Unit:  "bytes",
					Start: 200,
					End:   999,
				},
				{
					Unit:  "bytes",
					Start: 2000,
					End:   2499,
				},
				{
					Unit:  "bytes",
					Start: 9500,
					End:   -1,
				},
			},
			wantErr: false,
		},
		{
			name:     "empty range header",
			rangeStr: "",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid range format",
			rangeStr: "bytes=0-499-1000",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "non-numeric start",
			rangeStr: "bytes=abc-1000",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "non-numeric end",
			rangeStr: "bytes=1000-xyz",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "negative start and end - invalid",
			rangeStr: "bytes=-100--200",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "start greater than end",
			rangeStr: "bytes=500-400",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "multiple ranges with invalid format",
			rangeStr: "bytes=200-999, abc-xyz",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "range with only comma",
			rangeStr: "bytes=,",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "range with missing equal sign",
			rangeStr: "bytes0-499",
			want:     nil,
			wantErr:  true,
		},
	}

	parser := Parser{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranges, err := parser.Parse(tt.rangeStr)

			if tt.wantErr != (err != nil) {
				t.Fatalf("Test failed: err expected %v. received: %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(ranges, tt.want) {
				t.Errorf("Test failed: expected %v. received: %v", tt.want, ranges)
			}
		})
	}
}
