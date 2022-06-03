//Actividad 5.2 Programación paralela y concurrente
//Archivo Paralelo

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
	"runtime"
	"sync"
)

//Declaracion de variables
var n float64

type vector []int

//Checa si los numeros son primos
func check_prime(n float64) bool {
	for i := 2.0; i <= math.Sqrt(n); i += 1.0 {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}

//Funcion para sumar los numeros primos
func sec_prime(begining, ending, step int, ch chan int, wg *sync.WaitGroup) {
	sum := <-ch
	for i := begining; i < ending; i += step {
		if check_prime(float64(i)) == true {
			sum += int(i)
		}
	}
	ch <- sum
	wg.Done()
}

/*Encargado de dividir el tiempo entre el numero de hilos asignados, con
el fin de aumentar la velocidad del chequeo*/
func rango_div(hilos, limite int) vector {
	var rango vector
	unidad := limite / hilos
	for i := 1; i < hilos; i++ {
		if i >= hilos && i*unidad != limite {
			rango[i] = (unidad * i) + limite - (unidad * i)
		} else {
			rango[i] = (unidad * i)
		}
	}
	return rango
}
func main() {
	hilos := runtime.NumCPU()
	n := 5000000
	sumCh := make(chan int)
	var wg sync.WaitGroup
	div_range := rango_div(hilos, n)
	for i := 0; i < hilos-1; i++ {
		wg.Add(1)
		sec_prime(div_range[i], div_range[(i+1)], 1, sumCh, &wg)
		wg.Wait()
	}
	result := <-sumCh
	fmt.Println(result)
}
