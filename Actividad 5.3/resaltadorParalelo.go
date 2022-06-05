/*
Actividad Integradora 5.3 - Resaltador de sintaxis paralelo
Fecha: 10-06-2022
Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
    - Brenda Elena Saucedo González | A00829855
En el presente programa se definen categorías léxicas pertenecientes al lenguaje de programación de C++
*/

package main

// Librerias a usar
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func indexOf(element string, strSlice []string) int {
	for k, v := range strSlice {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func containsArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isLiteral(expresion []string, original []string, pos int, operador []bool) bool {
	// Se declaran variables para manejar los casos de excepción que nos indican que no son literales
	numeros := []string{}
	punto := false
	guion := false
	eReal := false
	fReal := false
	uReal := false
	lReal := 0
	wait := false
	letter := false

	// Si se encontro una o doble comilla, retorna verdadero, ya que se esta por leer un string o char
	if expresion[0] == `'` || expresion[0] == `"` {
		return true
	}

	// Añade los numeros del 0 al 9
	for i := 0; i < 10; i++ {
		numeros = append(numeros, (strconv.Itoa(i)))
	}

	// Guarda la posicion en la que se encuentra el inicio de la expresion
	pos2 := indexOf(expresion[0], original)

	// Ciclo que itera cada caracter de la expresion
	for i := 0; i < len(expresion); i++ {
		//Verifica que sea un número
		if (containsArray(numeros, expresion[i])) && (!letter) {
			wait = true
		} else {
			if (expresion[i] == ".") && (!punto) && (containsArray(numeros, original[pos2+i+1])) {
				// Verifica si es real
				punto = true
			} else if (expresion[i] == "L" || expresion[i] == "l") && (!wait) && ((lReal < 2 && !uReal) || (uReal && lReal == 0) || (indexOf("ul", expresion[:i]) != -1)) && (!fReal) {
				// Verifica si es un dato de tipo long o long long
				lReal = lReal + 1
				letter = true
			} else if (expresion[i] == "U" || expresion[i] == "u") && (!uReal) && (!punto) && (!eReal) && (original[pos2+i+1] != ".") {
				// Verifica si es un dato de tipo unsigned
				uReal = true // Antes era uReal + 1. Aquí parece que la intención es hacer que cambie de false a true. (???)
				letter = true
			} else if (expresion[i] == "E" || expresion[i] == "e") && (!eReal) && ((containsArray(numeros, original[pos2+i+1])) || original[pos2+i+1] == "-") {
				// Verifica si se esta leyendo la "E" de la notación científica y que reciba un número o guión a continuación
				operador[0] = true
				eReal = true
				wait = true
			} else if expresion[i] == "-" && (!guion) && eReal && (containsArray(numeros, original[pos2+i+1])) {
				// Verifica si es un guión y que reciba un número a continuación
				guion = true
				wait = true
			} else if (expresion[i] == "F" || expresion[i] == "f") && (!wait) && (!uReal) && (lReal == 0) && eReal {
				// Verifica si es un fast int
				fReal = true // Antes era fReal + 1. Aquí parece que la intención es hacer que cambie de false a true. (???)
				letter = true
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	// Definimos las variables que usaremos para abrir el archivo con la ayuda de la libreria OS.
	// Esto nos ayudará a evitar conflictos a la hora de abrirlo en equipos distintos.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}

	// Lista que guardará el contenido del archivo TXT
	lista_sintaxis := []string{}

	// Abre el archivo de texto (sintaxis.txt)
	fileTxt, ferr := os.Open(dir + "\\Actividad 5.3\\sintaxis.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(fileTxt)
	for scanner.Scan() {
		lista_sintaxis = append(lista_sintaxis, scanner.Text())
	}

	// Se abre o se crea un archivo html (index.html)
	fileHtml, e := os.Create(dir + "\\Actividad 5.3\\index.html")
	if e != nil {
		fmt.Println(e)
	}

	// Escribimos el head del archivo html
	fileHtml.WriteString("<!DOCTYPE html>\n")
	fileHtml.WriteString("<html>\n")
	fileHtml.WriteString("\t<head>\n")
	fileHtml.WriteString("\t\t<meta charset=\"utf-8\"/>\n")
	fileHtml.WriteString("\t\t<title>Resaltador de Sintaxix</title>\n")
	fileHtml.WriteString("\t\t<link rel=\"stylesheet\" href=\"style.css\">\n")
	fileHtml.WriteString("\t</head>\n")
	fileHtml.WriteString("\t<body>\n")

	// Definimos variables para manejar comentarios largos
	// comentarioLargo := false
	// posComentarioLargo := 0
	// originPos := 0

	/////////////////////////////////////// AQUÍ IRA LAS DEMAS FUNCIONES DE VERIFICACION DE SINTAXIX ///////////////////////////////////////
	// Lee cada enunciado del archivo de texto
	for nRow := 0; nRow < len(lista_sintaxis); nRow++ {
		// // Acumulador de la expresion
		// acumExp := ""
		// // Variables para la indentación
		// start = False
		// espacio := ""
		// // Variable para manejar espacios en literales de tipo string o char
		// nullSpace := false
		// // Variable para manejar "-" en los literales númericos
		// operadorOmit := false
		// // Variable para manejar las librerías
		// libreria := false

		// Lee cada caracter del enunciado
		for nTok := 0; nTok < len(lista_sintaxis[nRow]); nTok++ {
			expresion := lista_sintaxis[nRow]
			token := expresion[nTok : nTok+1]
			fmt.Print(token)
		}
		fmt.Println()
	}

	// Escribimos el final del archivo html
	fileHtml.WriteString("\t</body>\n")
	fileHtml.WriteString("</html>")

	// Cerramos el archivo HTML
	fileHtml.Close()
}
