// Actividad 5.2 Programación paralela y concurrente
// Archivo Paralelo

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
	"runtime"
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
	hilos := runtime.NumCPU()
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

/*
Se utilizó el SO de Diego como referencia:
	En un Ryzen 5 3500X termino en los siguientes tiempos con las siguientes configuraciones:

	Salida de consola (Hilos = 1024):
		[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\Paralelo\main.go"
		838596693108
		[Done] exited with code=0 in 11.894 seconds

	Salida de consola (Hilos = 12):
		[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\Paralelo\main.go"
		838596693108
		[Done] exited with code=0 in 13.794 seconds

	Salida de consola (hilos := runtime.NumCPU(), que en este caso seria equivalente a hilos := 6)
		[Running] go run "e:\Seagate_4tb\Documentos\Github_clone\Equipo-Racket-1\Actividad 5.2\Paralelo\main.go"
		838596693108
		[Done] exited with code=0 in 15.573 seconds



Tomando en cuenta el último resultado en donde tenemos 6 procesos (Versión paralela) y el resultado secuencial en donde
	p = 6
	T1 = 62.389
	T(p) = 15.573
podemos decir que nuestro S(p) en esta ocación es de = 4.00622873
*/
