package config

import (
	"reflect"
	"testing"
)

func TestNewSettings(t *testing.T) {
	s := NewSettings()

	want := "dbzo.Settings"
	got := reflect.TypeOf(s).String()

	if want != got {
		t.Errorf(`Returned type %s, wanted type %s`, got, want)
	}

}
