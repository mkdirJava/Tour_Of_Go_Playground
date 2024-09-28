package controlflow

import (
	"fmt"
)

func SquareRoot(x float64) float64 {
	return squareRoot(x, 1.0, 1.0)
}

func squareRoot(x float64, runningValue float64, tolerance float64) float64 {
	newCalculatedValue := runningValue - (runningValue*runningValue-x)/(2*runningValue)
	fmt.Printf("Iterating again the value is %f", newCalculatedValue)

	if runningValue-newCalculatedValue <= tolerance && runningValue != 1 {
		return runningValue
	} else {
		return squareRoot(x, newCalculatedValue, tolerance)
	}
}
