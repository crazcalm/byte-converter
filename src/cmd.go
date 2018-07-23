package bc

import (
	"strings"
)

const (
	empty = ""
	b     = "B"
	kb    = "KB"
	mb    = "MB"
	gb    = "GB"
	tb    = "TB"
)

var validFrom = []string{b, kb, mb, gb, tb}
var validTo = []string{empty, b, kb, mb, gb, tb}

//ValidInputFrom -- Checks valid input for the from unit of measure
func ValidInputFrom(unit string) bool {
	return validInput(unit, validFrom)
}

//ValidInputTo -- Checks valid input for the to unit of measure
func ValidInputTo(unit string) bool {
	return validInput(unit, validTo)
}

func validInput(unit string, bounds []string) (result bool) {
	for _, valid := range bounds {
		if strings.EqualFold(valid, unit) {
			result = true
			return
		}
	}
	return
}

//ReasonableOutput -- tries to convert the value to a unit of measure that is reasonable
func ReasonableOutput(value float64, from Unit) (numAnswer float64, strAnswer string, err error) {
	unitStringMap := map[Unit]string{B: b, KB: kb, MB: mb, GB: gb, TB: tb}
	topDown := []Unit{TB, GB, MB, KB, B}

	for _, unit := range topDown {
		if strings.EqualFold(unitStringMap[unit], unitStringMap[from]) {
			continue
		}

		numAnswer, strAnswer, err = Convert(value, from, unit)

		//Baseline for reasonable output
		if numAnswer > 0.1 {
			return
		}
	}
	//If not reasonable outout is found, return the smallest conversion
	return
}
