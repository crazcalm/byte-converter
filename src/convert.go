package bc

import (
	"fmt"
	"math"
)

//Unit of measure
type Unit string

const (
	//B -- Byte
	B = Unit("B")
	//KB -- Kilobyte
	KB = Unit("KB")
	// MB -- Megabyte
	MB = Unit("MB")
	//GB -- Gigabyte
	GB = Unit("GB")
	//TB -- Terabyte
	TB       = Unit("TB")
	kilobyte = 1024
)

var (
	unitMap     = map[Unit]float64{B: 1, KB: kilobyte, MB: megabyte, GB: gigabyte, TB: terabyte}
	megabyte    = math.Pow(1024, 2)
	gigabyte    = math.Pow(1024, 3)
	terabyte    = math.Pow(1024, 4)
	unitsString = fmt.Sprintf("%s, %s, %s, %s or %s", B, KB, MB, GB, TB)
)

func convertFromBytesToUnit(totalBytes int64, unit Unit) (result float64, err error) {
	unitStandard, ok := unitMap[unit]
	if !ok {
		err = fmt.Errorf("%s is not a valid unit. Try one of these instead: %s", unit, unitsString)
	}
	result = float64(totalBytes) / unitStandard
	return
}

func convertToBytes(num float64, unit Unit) (result int64, err error) {
	unitStandard, ok := unitMap[unit]
	if !ok {
		err = fmt.Errorf("%s is not a valid unit. Try one of these instead: %s", unit, unitsString)
		return
	}
	result = int64(num * unitStandard)
	return
}

//Convert -- converts one unit of measure (B, KB, MB, GB, TB) to another.
func Convert(num float64, fromUnit, toUnit Unit) (intAnswer float64, strAnswer string, err error) {
	var totalBytes int64

	totalBytes, err = convertToBytes(num, fromUnit)
	if err != nil {
		return
	}

	intAnswer, err = convertFromBytesToUnit(totalBytes, toUnit)
	if err != nil {
		return
	}

	strAnswer = fmt.Sprintf("%.2f %s", intAnswer, toUnit)
	return
}
