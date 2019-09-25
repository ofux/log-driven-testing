package assertions

import (
	"regexp"
	"testing"
)

func Test_regex_Assert(t *testing.T) {
	type fields struct {
		regexp *regexp.Regexp
	}
	type args struct {
		logLine string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantFailMessage string
		wantSuccessful  bool
	}{
		{
			name:            "Log actually contains the searched string",
			fields:          fields{regexp: regexp.MustCompile("Present\\!")},
			args:            args{logLine: "This is Present! as expected."},
			wantFailMessage: "",
			wantSuccessful:  true,
		}, {
			name:            "Log does not contain the searched string",
			fields:          fields{regexp: regexp.MustCompile("not present")},
			args:            args{logLine: "This isn't present as expected."},
			wantFailMessage: "Line 'This isn't present as expected.' doesn't match regex 'not present'.",
			wantSuccessful:  false,
		}, {
			name:            "Log matches the regular expression",
			fields:          fields{regexp: regexp.MustCompile("^[A-Z][a-z]+ matches.*")},
			args:            args{logLine: "This matches!!!"},
			wantFailMessage: "",
			wantSuccessful:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := regex{
				regexp: tt.fields.regexp,
			}
			gotFailMessage, gotSuccessful := r.Assert(tt.args.logLine)
			if gotFailMessage != tt.wantFailMessage {
				t.Errorf("Assert() gotFailMessage = %v, want %v", gotFailMessage, tt.wantFailMessage)
			}
			if gotSuccessful != tt.wantSuccessful {
				t.Errorf("Assert() gotSuccessful = %v, want %v", gotSuccessful, tt.wantSuccessful)
			}
		})
	}
}
