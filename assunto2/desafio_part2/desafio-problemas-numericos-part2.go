package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Número: ", i, " - PANPIN")
		} else if i%3 == 0 {
			fmt.Println("Número: ", i, " - PAN")
		} else if i%5 == 0 {
			fmt.Println("Número: ", i, " - PIN")
		}
	}
}
