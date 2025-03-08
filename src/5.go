  package main

import "math/rand"

func GenerateCode() {
	return strconv.Itoa(rand.Intn(10)) + "x^2" + strconv.Itoa(rand.Intn(10)) + "+x" + strconv.Itoa(rand.Intn(10))
}
