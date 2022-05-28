//Actividad 5.2 Programación paralela y concurrente

/*Fecha: 27-05-2020
Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
    - Brenda Elena Saucedo González | A00829855*/

/*Utilizando el lenguaje de programaci ́on indicado por tu profesor (Scheme, Racket, Clojure), escribe dos versiones
de un programa que calcule la suma de todos los n ́umeros primos menores a 5,000,000 (cinco millones):
• La primera versi ́on debe ser una implementaci ́on convencional que realice el c ́omputo de manera secuencial.
• La segunda versi ́on debe realizar el cómputo de manera paralela a trav ́es de los mecanismos provistos por
el lenguaje siendo utilizado (por ejemplo places o la funci ́on pmap). Debes procurar paralelizar el c ́odigo
aprovechando todos los n ́ucleos disponibles en tu sistema.
Ambas versiones del programa deben dar 838,596,693,108 como resultado.*/

package main

import (
	"fmt"
	"math"
)

//Declaracion de variables
var n float64

//Checa si los numeros son primos
func check_prime(n float64) bool {
	for i := 2.0; i <= math.Sqrt(n); i += 1.0 {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Test")
}
