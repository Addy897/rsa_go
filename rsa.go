package main

import (
	"fmt"
	"math"
	"os"
)

func gcd(a int, b int) int {
	result := min(a, b)
	for result > 0 {
		if a%result == 0 && b%result == 0 {
			break
		}
		result--
	}
	return result
}

func Usage() {
	fmt.Printf("[-] Usage: %s -e or -d Message", os.Args[0])
}
func generateKeys() (int, int, int) {
	p := 3
	q := 5
	N := p * q
	phi := (p - 1) * (q - 1)
	fmt.Printf("phi: %d\n", phi)
	e := 2
	for e < phi {
		if gcd(e, phi) == 1 {
			break
		}
		e++
	}
	d := (1 + phi) / e
	return e, d, N
}
func encrypt(e int, N int, msg int) int {
	c := math.Pow(float64(msg), float64(e))
	c = math.Mod(c, float64(N))
	return int(c)
}
func decrypt(d int, N int, msg int) int {
	m := math.Pow(float64(msg), float64(d))
	m = math.Mod(m, float64(N))
	return int(m)
}
func run() {
	e, d, N := generateKeys()
	msg := 12
	fmt.Printf("E: %d D: %d N: %d Msg: %d\n", e, d, N, msg)
	en := encrypt(e, N, msg)

	fmt.Printf("Encoded: %d\n", en)

	fmt.Printf("Decoded: %d", decrypt(d, N, en))
}
func main() {
	run()
}
