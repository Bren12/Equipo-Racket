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
	"strings"
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

func isLibreria(expresion []string) bool {
	//Busca el # que en C++ indica una libreria a incluir
	pos := indexOf("#", expresion)
	//Dado caso de que la encuentre marcalo como verdadero
	if pos == 0 {
		return true
	}
	return false
}

func isReservada(expresion string) bool {
	// Se definen las palabras reservadas como un diccionario
	/*reservada = {"int": True, "bool": True, "char": True, "void": True, "float": True, "double": True, "string": True, "cin": True, "cout": True, "while": True,
	  "as": True, "using": True, "namespace": True, "auto": True, "const": True, "asm": True, "dynamic_cast": True, "reinterpret_cast": True, "try": True,
	  "explicit": True, "new": True, "static_cast": True, "static": True, "typeid": True, "catch": True, "false": True, "operator": True, "template": True,
	  "typename": True, "class": True, "friend": True, "private": True, "this": True, "const_cast": True, "inline": True, "public": True, "throw": True,
	  "virtual": True, "delete": True, "enum": True, "goto": True, "else": True, "mutable": True, "protected": True, "true": True, "wchar_t": True, "endl": True,
	  "sizeof": True, "register": True, "unsigned": True, "break": True, "continue": True, "extern": True, "if": True, "return": True, "switch": True, "case": True,
	  "default": True, "short": True, "struct": True, "volatile": True, "do": True, "for": True, "long": True, "signed": True, "union": True, "std": True,}*/

	reservada := []string{"int", "bool", "char", "void", "float", "double", "string", "cin", "cout", "while",
		"as", "using", "namespace", "auto", "const", "asm", "dynamic_cast", "reinterpret_cast", "try",
		"explicit", "new", "static_cast", "static", "typeid", "catch", "false", "operator", "template",
		"typename", "class", "friend", "private", "this", "const_cast", "inline", "public", "throw",
		"virtual", "delete", "enum", "goto", "else", "mutable", "protected", "true", "wchar_t", "endl",
		"sizeof", "register", "unsigned", "break", "continue", "extern", "if", "return", "switch", "case",
		"default", "short", "struct", "volatile", "do", "for", "long", "signed", "union", "std"}
	// Verifica si existe en la expresión cualquier palabra reservada, si no, retorna falso
	for i := 0; i < len(reservada); i++ {
		// Si encontro el delimitador, se retorna verdadero
		if string(reservada[i]) == expresion || strings.Index(expresion, reservada[i]) != -1 {
			return true
		}
	}
	return false
}

func isOperador(expresion string, original string, pos int) bool {
	// Se definen los operadores como una lista
	operador := []string{"+", "+=", "++", "-", "-=", "--", "%", "%=", "*", "*=", "/=", "^", "<", "<<", ">", ">>", "<=", ">=", "=", "==", "!", "!=", "~", "?", "&", "&&", "||"}
	// Ciclo que itera cada operador de la lista definida
	for i := 0; i < len(operador); i++ {
		// Si encontro el operador, se retorna verdadero
		if operador[i] == expresion || strings.Index(expresion, operador[i]) != -1 || (expresion == "/" && string(original[pos:pos+2]) != "//") || (strings.Index(expresion, "/") != -1 && string(original[pos:pos+2]) != "//") {
			return true
		}
	}
	return false
}

func isOperadorUnique(expresion string, original string, pos int) bool {
	// Se definen los operadores como un diccionario
	operador := []string{"+", "+=", "++", "-", "-=", "--", "%", "%=", "*", "*=", "/=", "^", "<", "<<", ">", ">>",
		"<=", ">=", "=", "==", "!", "!=", "~", "?", "&", "&&", "||"}

	// Verifica si existe en la expresión cualquier operador, si no, retorna falso
	// Para eso, primero se verifica que no vaya a sobrepasar la longitud de la expresión original
	if pos+2 < len(original) {
		// Verifica que no sea un comentario lo que se esta leyendo
		if (expresion == "/" && string(original[pos:pos+2]) != "//") && (expresion == "/" && string(original[pos:pos+2]) != "/*") {
			return true
			// Si encuentra un operador retorna verdadero
		} else {
			//Expresion está ubicada en ooperador??
			for i := 0; i < len(operador); i++ {
				// Si encontro el delimitador, se retorna verdadero
				if operador[i] == expresion || strings.Index(expresion, operador[i]) != -1 {
					return true
				}
			}
		}
	}
	// En caso contrario, retorna falso
	return false
}

func isDelimitador(expresion string) bool {
	// Se definen los delimitadores como una lista
	delimitador := []string{"(", ")", "[", "]", "{", "}", ",", ";", "...", ":"}
	// Ciclo que itera cada delimitador de la lista definida
	for i := 0; i < len(delimitador); i++ {
		// Si encontro el delimitador, se retorna verdadero
		if delimitador[i] == expresion || strings.Index(expresion, delimitador[i]) != -1 {
			return true
		}
	}
	return false
}

func isIdentificador(expresion string, original string, pos int) bool {
	// Se crea una lista para checar todos los identificadores con letras
	alfabeto := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z"}

	// Se crea una lista para números
	numeros := []string{}
	// Añade a la lista los numeros del 0 al 9
	for i := 0; i < 10; i++ {
		numeros = append(numeros, (strconv.Itoa(i)))
	}

	// Ciclo que itera cada letra del alfabeto
	for i := 0; i < len(alfabeto); i++ {
		// Si encontro una letra del afabeto o un guión entra a la siguiente condicional
		if alfabeto[i] == expresion || strings.Index(expresion, alfabeto[i]) != -1 || strings.Contains(expresion, alfabeto[i]) || strings.Contains(expresion, "_") {
			// Checa casos de excepcion que indican que no es un identificador
			if containsArray(alfabeto, string(expresion[0])) && (!(strings.Contains(expresion, `"`) || strings.Contains(expresion, `'`) || strings.Contains(expresion, ".") || strings.Contains(expresion, "#"))) {
				return true
			}
		}
	}
	return false
}

func isLiteral(expresion string, original string, pos int, operador []bool) bool {
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
	if string(expresion[0]) == `'` || string(expresion[0]) == `"` {
		return true
	}

	// Añade los numeros del 0 al 9
	for i := 0; i < 10; i++ {
		numeros = append(numeros, (strconv.Itoa(i)))
	}

	// Guarda la posicion en la que se encuentra el inicio de la expresion

	pos2 := strings.Index(original, string(expresion[0]))

	// Ciclo que itera cada caracter de la expresion
	for i := 0; i < len(expresion); i++ {
		//Verifica que sea un número
		if (containsArray(numeros, string(expresion[i]))) && (!letter) {
			wait = true
		} else {
			if (string(expresion[i]) == ".") && (!punto) && (containsArray(numeros, string(original[pos2+i+1]))) {
				// Verifica si es real
				punto = true
			} else if (string(expresion[i]) == "L" || string(expresion[i]) == "l") && (!wait) && ((lReal < 2 && !uReal) || (uReal && lReal == 0) || (strings.Index(expresion[:i], "ul") != -1)) && (!fReal) {
				// Verifica si es un dato de tipo long o long long
				lReal = lReal + 1
				letter = true
			} else if (string(expresion[i]) == "U" || string(expresion[i]) == "u") && (!uReal) && (!punto) && (!eReal) && (string(original[pos2+i+1]) != ".") {
				// Verifica si es un dato de tipo unsigned
				uReal = true // Antes era uReal + 1. Aquí parece que la intención es hacer que cambie de false a true. (???)
				letter = true
			} else if (string(expresion[i]) == "E" || string(expresion[i]) == "e") && (!eReal) && ((containsArray(numeros, string(original[pos2+i+1]))) || string(original[pos2+i+1]) == "-") {
				// Verifica si se esta leyendo la "E" de la notación científica y que reciba un número o guión a continuación
				operador[0] = true
				eReal = true
				wait = true
			} else if string(expresion[i]) == "-" && (!guion) && eReal && (containsArray(numeros, string(original[pos2+i+1]))) {
				// Verifica si es un guión y que reciba un número a continuación
				guion = true
				wait = true
			} else if (string(expresion[i]) == "F" || string(expresion[i]) == "f") && (!wait) && (!uReal) && (lReal == 0) && eReal {
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
