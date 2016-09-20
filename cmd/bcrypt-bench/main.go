package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println(`
                      BCRYPT BENCHMARK

This  test program will encrypt and compare bcrypt hashes  at
increasing   costs   until  comparing  takes  over  a  second.
A good cost should take about ~120ms to compare a single hash.
`)

	testPassword := []byte("aRaNdoMPa55W2ord.")
	ideal := bcrypt.MinCost

	for cost := bcrypt.MinCost; ; cost++ {
		fmt.Printf("Cost %3d - ", cost)
		start := time.Now()
		hash, _ := bcrypt.GenerateFromPassword(testPassword, cost)
		timeTook := time.Now().Sub(start)
		fmt.Printf("hashing took %11s - ", timeTook.String())

		start = time.Now()
		bcrypt.CompareHashAndPassword(hash, testPassword)
		timeTook = time.Now().Sub(start)
		fmt.Printf("compare took %11s\r\n", timeTook.String())

		if timeTook > time.Millisecond*100 {
			if timeTook < time.Millisecond*160 || ideal == bcrypt.MinCost {
				ideal = cost
			}
		}
		if timeTook > time.Second {
			break
		}
	}

	fmt.Printf("\r\nYour ideal bcrypt cost is: %d\r\n\r\n", ideal)
}
