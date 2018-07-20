package bc

import (
	"strings"
	"testing"
)

func TestErrorCases(t *testing.T) {
	notAUnit := Unit("unknown")
	errMsg := "unknown is not a valid unit. Try one of these instead: B, KB, MB, GB or TB"
	var err error

	_, err = convertToBytes(10, notAUnit)
	if !strings.EqualFold(err.Error(), errMsg) {
		t.Errorf("Expected %s, but got %s", errMsg, err.Error())
	}

	_, err = convertFromBytesToUnit(10, notAUnit)
	if !strings.EqualFold(err.Error(), errMsg) {
		t.Errorf("Expected %s, but got %s", errMsg, err.Error())
	}

	_, _, err = Convert(10, notAUnit, B)
	if !strings.EqualFold(err.Error(), errMsg) {
		t.Errorf("Expected %s, but got %s", errMsg, err.Error())
	}

	_, _, err = Convert(10, B, notAUnit)
	if !strings.EqualFold(err.Error(), errMsg) {
		t.Errorf("Expected %s, but got %s", errMsg, err.Error())
	}

}

func TestConvert(t *testing.T) {
	tests := []struct {
		Num      float64
		FromUnit Unit
		ToUnit   Unit
		Answer   float64
		String   string
	}{
		{1024, B, B, 1024, "1024.00 B"},
		{1024, B, KB, 1, "1.00 KB"},
		{262144, B, MB, 0.25, "0.25 MB"},
		{268435456, B, GB, 0.25, "0.25 GB"},
		{274877906944, B, TB, 0.25, "0.25 TB"},
		{1024, KB, MB, 1, "1.00 MB"},
		{262144, KB, GB, 0.25, "0.25 GB"},
		{268435456, KB, TB, 0.25, "0.25 TB"},
		{1024, MB, GB, 1, "1.00 GB"},
		{262144, MB, TB, 0.25, "0.25 TB"},
		{1024, GB, TB, 1, "1.00 TB"},
		{1, TB, GB, 1024, "1024.00 GB"},
		{0.25, TB, MB, 262144, "262144.00 MB"},
		{0.25, TB, KB, 268435456, "268435456.00 KB"},
		{0.25, TB, B, 274877906944, "274877906944.00 B"},
	}

	for _, test := range tests {
		result, resultStr, err := Convert(test.Num, test.FromUnit, test.ToUnit)
		if err != nil {
			t.Errorf("Failed: %s", err.Error())
		}

		if result != test.Answer {
			t.Errorf("Expected %f %s to be %f %s, but got %f", test.Num, test.FromUnit, test.Answer, test.ToUnit, result)
		}

		if !strings.EqualFold(resultStr, test.String) {
			t.Errorf("Expected %s, but got %s", test.String, resultStr)
		}
	}
}

func TestConvertFromBytesToUnit(t *testing.T) {
	tests := []struct {
		Num    int64
		Unit   Unit
		Answer float64
	}{
		{1, B, 1},
		{1024, KB, 1},
		{262144, MB, 0.25},
		{268435456, GB, 0.25},
		{274877906944, TB, 0.25},
	}

	for _, test := range tests {
		result, err := convertFromBytesToUnit(test.Num, test.Unit)
		if err != nil {
			t.Errorf("failed: %s", err.Error())
		}

		if result != test.Answer {
			t.Errorf("Expected %d bytes to be %f %s, but got %f", test.Num, test.Answer, test.Unit, result)
		}
	}
}

func TestConvertToBytes(t *testing.T) {
	tests := []struct {
		Num    float64
		Unit   Unit
		Answer int64
	}{
		{1, B, 1},
		{1, KB, 1024},
		{1, MB, 1048576},
		{1, GB, 1073741824},
		{1, TB, 1099511627776},
	}

	for _, test := range tests {
		result, err := convertToBytes(test.Num, test.Unit)
		if err != nil {
			t.Errorf("Failed: %s", err.Error())
		}

		if result != test.Answer {
			t.Errorf("Expected %.2f %s to equal %d bytes, but got %d", test.Num, test.Unit, test.Answer, result)
		}
	}
}
