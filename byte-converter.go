package main

import (
	"fmt"

	"github.com/crazcalm/byte-converter/src"
)

func main() {
	fmt.Println(bc.Convert(1024, bc.B, bc.B))
}
