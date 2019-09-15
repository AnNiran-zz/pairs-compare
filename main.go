package main

import "fmt"

var divider = 2147483647

var initValueGenA = 65
var initValueGenB = 8921

func main() {

	number := calculateNumberOfMatchingPairs()
	fmt.Println(number)
}

// calculates next value for Generator A
func calculateGenANextValue() int {
	// take previous/first value and multiply it by factor -> return remainder
	return (initValueGenA * 16807) % divider
}

// calculate next value for Generator B
func calculateGenBNextValue() int {
	// take previous/first value and multiply it by factor -> return remainder
	return (initValueGenB * 48271) % divider
}

// obtain least significant 16 bits of a value
func getLeastSignificant16Bits(value int) uint16 {
	valueUint32 := uint32(value)
	least16Bits := ^uint32(0) >> (32 - 16)

	// return least significant bits for value
	return uint16(valueUint32 & least16Bits)
}

// gather matching pairs number by looping 40 million times
func calculateNumberOfMatchingPairs() int {
	var matchingPairs int

	// loop over 40 million calculations and gather result
	for i := 0; i < 40000001; i++ {

		// calculate next values for both generators
		resultGenA := calculateGenANextValue()
		resultGenB := calculateGenBNextValue()

		// reassign initial values for both generators
		initValueGenA = resultGenA
		initValueGenB = resultGenB

		// obtain least significant 16 bits from results
		resultGenALeastSignif16Bits := getLeastSignificant16Bits(resultGenA)
		resultGenBLeastSignif16Bits := getLeastSignificant16Bits(resultGenB)

		// increase number of matching pairs if least significant 16 bits values match
		if resultGenALeastSignif16Bits == resultGenBLeastSignif16Bits {
			matchingPairs++
		}
	}

	return matchingPairs
}
