package config

import (
	"regexp"
	"testing"
)

func TestVersionNumber(t *testing.T) {
	var tests = []struct {
		name    string
		server  bool
		pattern string
	}{
		{"server", true, `^Datumbazo Server \d+?\.\d+?\.\d+?`},
		{"client", false, `^Datumbazo Client \d+?\.\d+?\.\d+?`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := regexp.MustCompile(tt.pattern)
			ver := GetVersion(tt.server)
			if !want.MatchString(ver.Full) {
				t.Errorf(`VersionString = %q, want match for %#q, nil`, ver.Full, want)
			}
		})
	}
}
