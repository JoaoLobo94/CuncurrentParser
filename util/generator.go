package util

import (
	"math/rand"
	"time"
    )
	    
    func RandTransactionAmounts(min float64, numberToGenerate int) []float64 {
	rand.Seed(time.Now().UnixNano())
	max := 99.9
	result := make([]float64, numberToGenerate)
	for i := range result {
	    result[i] = min + rand.Float64() * (max - min)
	}
	return result
    }