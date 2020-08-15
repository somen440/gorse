package gorse_test

import (
	"fmt"
	"testing"

	"github.com/somen440/gorse"
)

func TestToZenkaku(t *testing.T) {
	str := "ｼﾞｬｳﾞｧ"
	expected := "ジャヴァ"
	actual := gorse.ToZenkaku(str)
	if expected != actual {
		t.Fatalf("not match. expected=%s, actual=%s", expected, actual)
	}
}

func TestToMorse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str      string
		expected string
	}{
		{"A", "・－"},
		{"B", "－・・・"},
		{"C", "－・－・"},
		{"D", "－・・"},
		{"E", "・"},
		{"F", "・・－・"},
		{"G", "－－・"},
		{"H", "・・・・"},
		{"I", "・・"},
		{"J", "・－－－"},
		{"K", "－・－"},
		{"L", "・－・・"},
		{"M", "－－"},
		{"N", "－・"},
		{"O", "－－－"},
		{"P", "・－－・"},
		{"Q", "－－・－"},
		{"R", "・－・"},
		{"S", "・・・"},
		{"T", "－"},
		{"U", "・・－"},
		{"V", "・・・－"},
		{"W", "・－－"},
		{"X", "－・・－"},
		{"Y", "－・－－"},
		{"Z", "－－・・"},

		{"こんにちは Tom さん", "－・－－－－－・－・・・－・－・・・・・・・・・・・・－ －－－－－－ ・・・・－－・"},

		{"ｳｨﾝｸﾞ", "・・－・・－・－－・・・－"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint("case_", i), func(t *testing.T) {
			t.Parallel()

			actual, err := gorse.ToMorse(tt.str)
			if err != nil {
				t.Fatal(err)
			}
			if tt.expected != actual {
				t.Fatalf("not equal expected=%s, actual=%s", tt.expected, actual)
			}
		})
	}
}
