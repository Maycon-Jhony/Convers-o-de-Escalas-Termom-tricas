package main

import "fmt"

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	// Ponto de ebulição da água em Kelvin
	pontoDeEbulicaoKelvin := 373.15

	// Convertendo para Celsius
	pontoDeEbulicaoCelsius := kelvinToCelsius(pontoDeEbulicaoKelvin)

	// Exibindo o resultado
	fmt.Printf("O ponto de ebulição da água em Kelvin é %.2f\n", pontoDeEbulicaoKelvin)
	fmt.Printf("Isso é equivalente a %.2f graus Celsius\n", pontoDeEbulicaoCelsius)
}
