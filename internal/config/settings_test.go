package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSettings(t *testing.T) {
	for _, b := range [2]bool{true, false} {
		t.Run(fmt.Sprintf("%v", b), func(t *testing.T) {
			s := NewSettings(b)

			want := "config.Settings"
			got := reflect.TypeOf(s).String()

			if want != got {
				t.Errorf(`Returned type %s, wanted type %s`, got, want)
			}
		})
	}
}
