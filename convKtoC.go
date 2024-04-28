/**
	Conversão de medidas - Ponto de ebulição da Agua de Kelvin em Celsius.
	Projeto com objetivo de estudo e aprendizado da linguagem GO
	O projeto faz parte do laboratorio de estudos do curso da DIO.me
	Professora Tenille Martins - Curso de GO Leng.
	Autor: Paulo Roberto - PRSant0s
	27/04/2024
	Versão 00.1

**/

package main

import "fmt"

const pontoEbulicaoKelvin float64 = 373 // Valor fixo do ponto de Ebulição em Kelvin

func main() {

	var tempKelvin = pontoEbulicaoKelvin
	var tempCelsius = (pontoEbulicaoKelvin - 273) // Formula aplicada na conversão: C = K -273 A diferença entre as escals Celsius e Kelvin
	// é o ponto 0

	fmt.Printf("O valor convertido em Culsius é: %g e o valor em Kelvin é: %g", tempCelsius, tempKelvin)
}
