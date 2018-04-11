package main

import (
	"fmt"
	"os"
	"strconv"

	"exercises-go/donovan/2-2/tempconv"
	"exercises-go/donovan/2-2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		k := weightconv.KG(t)
		p := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
}
