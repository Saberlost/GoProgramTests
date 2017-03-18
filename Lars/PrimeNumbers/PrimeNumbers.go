package main

import "math"
import "fmt"

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= n/2; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true

}

func bruteforcePrimesAndPrint(n int) {
	for i := 1; i < n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func getNextValue(p int, list []bool) int {
	i := p + 1
	for ; p < len(list) && list[p]; p++ {

	}
	return i
}

func sieveofEratosthenes(n int) []bool {
	list := make([]bool, n+1)
	for index := range list {
		list[index] = true

	}
	for p := 2; p <= n; p = getNextValue(p, list) {
		for i := 2; p*i <= n; i++ {
			list[p*i] = false
		}
	}
	return list

}

func solveLargerTable() {
	n := math.MaxInt32 - 100
	n = 50000000
	fmt.Println("Solving all primes up to", n, " will only count number of solutions found. ")
	list := sieveofEratosthenes(n)

	foundprimes := 0
	for _, v := range list {
		if v {
			//fmt.Println(i)
			foundprimes++
		}
	}
	fmt.Println(foundprimes)
}

func usesieveofEratosthenesAndPrint(n int) {
	list := sieveofEratosthenes(n)
	for i, v := range list {
		if v {
			fmt.Println(i)

		}
	}

}

func main() {

	bruteforcePrimesAndPrint(100)

	usesieveofEratosthenesAndPrint(100)
	usesieveofEratosthenesAndPrint(10000)

	solveLargerTable()
}
