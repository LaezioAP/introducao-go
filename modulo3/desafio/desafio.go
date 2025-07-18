// Criar um programa que faça a covnersão de escalas termométricas.
package main

import "fmt"

// Desafio: Programa para converter o valor do potno de ebulição da água de Kel para graus Celsius.

// DICA: C = K - 273

const tempKelvin = 373.15

func main()  {
	tempCelsius := tempKelvin - 273.15
	fmt.Printf("O valor do ponto de ebulição da água em C é = %g", tempCelsius)
} 