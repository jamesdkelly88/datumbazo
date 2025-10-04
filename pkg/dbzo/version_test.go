package dbzo

import (
	"regexp"
	"testing"
)

func TestVersionNumber(t *testing.T) {
	want := regexp.MustCompile(`^\d+?\.\d+?\.\d+?`)
	ver := GetVersion(true)
	if !want.MatchString(ver.Number) {
		t.Errorf(`VersionNumber = %q, want match for %#q, nil`, ver, want)
	}
}
