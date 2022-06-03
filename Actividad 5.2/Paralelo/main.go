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
Ambas versiones del programa deben dar 838,596,693,108 (838596693108) como resultado.*/

package main

import (
	"fmt"
	"math"
	"sync"
)

//Declaracion de variables
var n float64
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

//Declaracion de tipo de dato
type vector []int

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

//Funcion para sumar los numeros primos
func sec_prime(begining, ending, step int, ch chan int) {
	var sum int
	for i := begining; i < ending; i += step {
		if check_prime(float64(i)) == true {
			mutex.Lock()
			sum += i
			mutex.Unlock()
		}
	}
	ch <- sum
}

/*Encargado de dividir el numero entre el numero de procesos asignados, con
el fin de adecuar las goroutines al numero de procesos asignado*/
func rango_div(hilos, limite int) vector {
	var rango vector
	//Nuestro valor inicial para el rango es 0
	rango = append(rango, 0)
	//Base para poder definir el rango
	unidad := limite / hilos
	for i := 1; i < hilos+1; i++ {
		if i >= hilos && i*unidad != limite {
			rango = append(rango, (unidad*i)+limite-(unidad*i))
		} else {
			rango = append(rango, (unidad * i))
		}
	}
	return rango
}

func main() {
	//Declaramos las variables a usar
	//Hilos sera el numero de procesos a usar
	hilos := 1024
	//n sera el numero a analizar
	n := 5000000
	// sumCh se usa para poder pasar la suma de las goroutines
	sumCh := make(chan int)
	//Resultado recibira la información del canal sumCh
	var result int
	//Calculamos el rango
	div_range := rango_div(hilos, n)
	//Llamamos a los goroutines con el numero de procesos
	for i := 0; i < hilos; i++ {
		go sec_prime(div_range[i], div_range[(i+1)], 1, sumCh)
	}
	//Sumamos lo que tenga sumCh
	for i := 0; i < hilos; i++ {
		e := <-sumCh
		result += e
	}
	//Imprimimos el resultado
	fmt.Println(result)
}

/*Reporte de Diego:
En un Ryzen 5 3500X en un nucleo el programa (Metodo paralelo) termino en
64.424 segundos
Salida de consola:

[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\Paralelo\main.go"
838596693108
[Done] exited with code=0 in 11.894 seconds*/
