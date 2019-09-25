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
			c:               "not present",
			args:            args{logLine: "This isn't present as expected."},
			wantFailMessage: "Expected 'This isn't present as expected.' to contain 'not present'.",
			wantSuccessful:  false,
		}, {
			name:            "Log doesn't contain the searched string as this is case sensitive",
			c:               "Present",
			args:            args{logLine: "This isn't present as expected."},
			wantFailMessage: "Expected 'This isn't present as expected.' to contain 'Present'.",
			wantSuccessful:  false,
		}, {
			name:            "Log contains the searched string at the very beginning",
			c:               "Present",
			args:            args{logLine: "Present! This is expected."},
			wantFailMessage: "",
			wantSuccessful:  true,
		}, {
			name:            "Log contains the searched string at the very end",
			c:               "present.",
			args:            args{logLine: "This is expected to be present."},
			wantFailMessage: "",
			wantSuccessful:  true,
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
