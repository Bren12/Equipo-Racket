// Actividad 5.2 Programación paralela y concurrente
// Archivo Secuencial

/*
Fecha: 03-06-2020
Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
    - Brenda Elena Saucedo González | A00829855

Utilizando el lenguaje de programación indicado por tu profesor (Scheme, Racket, Clojure), escribe dos versiones
de un programa que calcule la suma de todos los números primos menores a 5,000,000 (cinco millones):
• La primera versión debe ser una implementación convencional que realice el cómputo de manera secuencial.
• La segunda versión debe realizar el cómputo de manera paralela a través de los mecanismos provistos por
el lenguaje siendo utilizado (por ejemplo places o la función pmap). Debes procurar paralelizar el código
aprovechando todos los núcleos disponibles en tu sistema.
Ambas versiones del programa deben dar 838,596,693,108 (838596693108) como resultado.
*/

package main

import (
	"fmt"
	"math"
)

//Declaracion de variables
var n float64

//Checa si los numeros son primos
func check_prime(n float64) bool {
	if n <= 1 {
		return false
	}
	for i := 2.0; i <= math.Sqrt(n); i += 1.0 {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}

//Funcion para realizar la parte secuencial de la actividad
func sec_prime(limit int) int {
	var sum int
	for i := 2; i < limit; i++ {
		if check_prime(float64(i)) == true {
			sum += int(i)
		}
	}
	return sum
}

func main() {
	//n sera el numero a analizar
	n := 5000000
	//imprimimos el resultado
	fmt.Println(sec_prime(n))
}

/*
Se utilizó el SO de Diego como referencia:
	En un Ryzen 5 3500X termino en 62.389 segundos
	
	Salida de consola:
		[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\Secuencial\main.go"
		838596693108
		[Done] exited with code=0 in 62.389 seconds
*/