package bc

import (
	"testing"
)

func TestReasonableOutput(t *testing.T) {
	tests := []struct {
		Num    float64
		Unit   Unit
		Answer float64
	}{
		{1024, B, 1.00},
		{262144, B, 0.25},
		{268435456, B, 0.25},
		{274877906944, B, 0.25},
	}

	for _, test := range tests {
		result, _, _ := ReasonableOutput(test.Num, test.Unit)
		if result != test.Answer {
			t.Errorf("Expected %.2f %s to be %.2f, but got %.2f", test.Num, test.Unit, test.Answer, result)
		}
	}
}

func TestValidInputTo(t *testing.T) {
	tests := []struct {
		Unit   string
		Answer bool
	}{
		{"", true},
		{"B", true},
		{"KB", true},
		{"MB", true},
		{"GB", true},
		{"TB", true},
		{"unkown", false},
	}

	for _, test := range tests {
		result := ValidInputTo(test.Unit)
		if result != test.Answer {
			t.Errorf("For %s, expected %t, but got %t", test.Unit, test.Answer, result)
		}
	}
}

func TestValidInputFrom(t *testing.T) {
	tests := []struct {
		Unit   string
		Answer bool
	}{
		{"", false},
		{"B", true},
		{"KB", true},
		{"MB", true},
		{"GB", true},
		{"TB", true},
		{"unkown", false},
	}

	for _, test := range tests {
		result := ValidInputFrom(test.Unit)
		if result != test.Answer {
			t.Errorf("For %s, expected %t, but got %t", test.Unit, test.Answer, result)
		}
	}
}
