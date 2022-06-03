//Actividad 5.2 Programación paralela y concurrente
//Archivo Secuencial

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
	n := 5000000
	fmt.Println(sec_prime(n))
	/*Reporte de Diego:
	En un Ryzen 5 3500X en un nucleo el programa (Metodo secuencial) termino en
	64.424 segundos
	Salida de consola:

	[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\main.go"
	838596693108
	[Done] exited with code=0 in 64.424 seconds

	En un Intel Core i5-11400H en un nucleo el programa (Metodo secuencial) termino en
	47.425 segundos
	Salida de consola:
	[Running] go run "c:\Users\Upijijis\OneDrive\Documentos\git\Go\Act 5.2\main.go"
	838596693108
	[Done] exited with code=0 in 47.425 seconds
	*/
}
