package main

import "fmt"

/*
Conversion from signed 8-bit to float
char a = -128;
char b = 127;
(a + 0.5f) / 32.0f  =  -3.984375
(b + 0.5f) / 32.0f  =   3.984375
*/

func ToSynapticWeight(encoded int8) float64 {
	return (float64(encoded) + 0.5) / 16.0
}

func main() {
	fmt.Println("Hello world");
}
