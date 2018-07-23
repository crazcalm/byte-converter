package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/crazcalm/byte-converter/src"
)

const (
	empty = ""
	b     = "B"
	kb    = "KB"
	mb    = "MB"
	gb    = "GB"
	tb    = "TB"
)

var num = flag.Uint64("num", 0, "num is used to pass in the value that you would like to convert")
var fromUnit = flag.String("from", b, fmt.Sprintf("The input's unit of measure: %s, %s, %s, %s or %s", b, kb, mb, gb, tb))
var toUnit = flag.String("to", empty, fmt.Sprintf("The outout's unit of measure: %s, %s, %s, %s or %s", b, kb, mb, gb, tb))
var err = fmt.Errorf(fmt.Sprintf("Try using a unit of measure in the form of %s, %s, %s, %s or %s", b, kb, mb, gb, tb))
var inputMap = map[string]bc.Unit{b: bc.B, kb: bc.KB, mb: bc.MB, gb: bc.GB, tb: bc.TB}
var validFrom = []string{b, kb, mb, gb, tb}
var validTo = []string{empty, b, kb, mb, gb, tb}

func validInput(unit string, bounds []string) (result bool) {
	for _, valid := range bounds {
		if strings.EqualFold(valid, unit) {
			result = true
			return
		}
	}
	return
}

func main() {
	fmt.Println(os.Args)
	flag.Parse()
	if !validInput(*fromUnit, validFrom) {
		fmt.Fprint(os.Stderr, err.Error())
	}
	if !validInput(*toUnit, validTo) {
		fmt.Fprint(os.Stderr, err.Error())
	}

	value := float64(*num)
	from := inputMap[*fromUnit]
	to, ok := inputMap[*toUnit]
	if !ok {
		topDown := []string{tb, gb, mb, kb, b}
		var strAnswer string
		var numAnswer float64
		var err error
		for _, unit := range topDown {
			if strings.EqualFold(unit, *fromUnit) {
				continue
			}
			numAnswer, strAnswer, err = bc.Convert(value, from, inputMap[unit])
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("Experienced an error: %s", err.Error()))
				os.Exit(1)
			}
			if numAnswer > 0.1 {
				fmt.Fprintf(os.Stdout, "%.2f %s is %s\n", value, *fromUnit, strAnswer)
				os.Exit(0)
			}
		}
		fmt.Fprintf(os.Stdout, "%.2f %s is %s\n", value, *fromUnit, strAnswer)
		os.Exit(0)
	}
	_, strAnswer, err := bc.Convert(value, from, to)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Experienced an error: %s", err.Error()))
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%.2f %s is %s\n", value, *fromUnit, strAnswer)
	os.Exit(0)

}
