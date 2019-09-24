package assertions

import (
	"testing"
)

func Test_containedString_Assert(t *testing.T) {
	type args struct {
		logLine string
	}
	tests := []struct {
		name            string
		c               containedString
		args            args
		wantFailMessage string
		wantSuccessful  bool
	}{
		{
			name:            "Log actually contains the searched string",
			c:               "Present!",
			args:            args{logLine: "This is Present! as expected."},
			wantFailMessage: "",
			wantSuccessful:  true,
		}, {
			name:            "Log doesn't contain the searched string",
			c:               "not Present!",
			args:            args{logLine: "This isn't Present! as expected."},
			wantFailMessage: "Expected 'This isn't Present! as expected.' to contain 'not Present!'.",
			wantSuccessful:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFailMessage, gotSuccessful := tt.c.Assert(tt.args.logLine)
			if gotFailMessage != tt.wantFailMessage {
				t.Errorf("Assert() gotFailMessage = %v, want %v", gotFailMessage, tt.wantFailMessage)
			}
			if gotSuccessful != tt.wantSuccessful {
				t.Errorf("Assert() gotSuccessful = %v, want %v", gotSuccessful, tt.wantSuccessful)
			}
		})
	}
}
