package util

import (
	"math/rand"
	"time"
	"fmt"
    )
	    
    func RandTransactionAmounts() []float64 {
	rand.Seed(time.Now().UnixNano())
	var numberToGenerate int
	fmt.Scanln(&numberToGenerate)
	min := 0.0
	max := 99.9
	result := make([]float64, numberToGenerate)
	for i := range result {
	    result[i] = min + rand.Float64() * (max - min)
	}
	return result
    }