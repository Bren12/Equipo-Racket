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
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func containsArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isFile(expresion string) bool {
	// Busca si hay "//" en la expresión
	pos := strings.Index(expresion, ".cpp")
	// Se definen algunos caracteres especiales que no son permitidos en los nombres de archivos
	caracteres := []string{"\\", "/", ":", "*", "?", "<", ">", "|"}
	// Retorna verdadero si encontro la extensión en la expresión
	if pos > 0 {
		// Un ciclo para recorrer la lista de caracteres
		for i, caract := range caracteres {
			// Un ciclo para recorrer la expresión antes del ".cpp"
			for j := 0; j < len(expresion[:pos]); j++ {
				// Verifica que no este incluido un caracter especial no permitido en el nombre del archivo
				if expresion[j:j+1] == caract {
					fmt.Println(i)
					return false
				}
			}
		}
		// Retorna verdadero ya que no encontro un caracter especial
		return true
	}
	// Retorna falso si no encontro la extensión en la expresión
	return false
}

func isComentario(expresion string) bool {
	// Busca si hay "//" en la expresión
	pos := strings.Index(expresion, "//")
	// Retorna verdadero si encontro "//" en la expresión
	if pos == 0 {
		return true
	}
	// Retorna falso, en caso contrario
	return false
}

func isLibreria(expresion string) bool {
	//Busca el # que en C++ indica una libreria a incluir
	pos := strings.Index(expresion, "#")
	//Dado caso de que la encuentre marcalo como verdadero
	if pos == 0 {
		return true
	}
	return false
}

func isReservada(expresion string) bool {
	// Se definen las palabras reservadas como un array
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
		if reservada[i] == expresion {
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
		if operador[i] == expresion || strings.Index(expresion, operador[i]) != -1 || (expresion == "/" && original[pos:pos+2] != "//") || (strings.Index(expresion, "/") != -1 && original[pos:pos+2] != "//") {
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
		if expresion == "/" && original[pos:pos+2] != "//" && original[pos:pos+2] != "/*" {
			return true
			// Si encuentra un operador retorna verdadero
		} else {
			//Expresion está ubicada en ooperador??
			for i := 0; i < len(operador); i++ {
				// Si encontro el delimitador, se retorna verdadero
				if operador[i] == expresion {
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
	numeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Ciclo que itera cada letra del alfabeto
	for i := 0; i < len(alfabeto); i++ {
		// Si encontro una letra del afabeto o un guión entra a la siguiente condicional
		if alfabeto[i] == expresion || strings.Index(expresion, alfabeto[i]) != -1 || containsArray(numeros, expresion) || strings.Contains(expresion, "_") {
			// Checa casos de excepcion que indican que no es un identificador
			if containsArray(alfabeto, string(expresion[0])) && (!(strings.Contains(expresion, "\"") || strings.Contains(expresion, "'") || strings.Contains(expresion, ".") || strings.Contains(expresion, "#"))) {
				return true
			}
		}
	}
	return false
}

func isLiteral(expresion string, original string, pos int, operador *bool) bool {
	// Se declaran variables para manejar los casos de excepción que nos indican que no son literales
	numeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	punto := false
	guion := false
	eReal := false
	fReal := false
	uReal := false
	lReal := 0
	wait := false
	letter := false

	// Si se encontro una o doble comilla, retorna verdadero, ya que se esta por leer un string o char
	if string(expresion[0]) == "'" || string(expresion[0]) == "\"" {
		return true
	}

	// Guarda la posicion en la que se encuentra el inicio de la expresion
	pos2 := strings.Index(original, expresion)

	// Ciclo que itera cada caracter de la expresion
	for i := 0; i < len(expresion); i++ {
		//Verifica que sea un número
		if (containsArray(numeros, string(expresion[i]))) && (!letter) {
			wait = false
		} else {
			// Verifica si es real
			if (string(expresion[i]) == ".") && (!punto) && (containsArray(numeros, string(original[pos2+i+1]))) {
				punto = true
			} else if (string(expresion[i]) == "L" || string(expresion[i]) == "l") && (!wait) && ((lReal < 2 && !uReal) || (uReal && lReal == 0) || (strings.Index(expresion[:i], "ul") != -1)) && (!fReal) {
				// Verifica si es un dato de tipo long o long long
				lReal = lReal + 1
				letter = true
			} else if (string(expresion[i]) == "U" || string(expresion[i]) == "u") && (!uReal) && (!punto) && (!eReal) && (string(original[pos2+i+1]) != ".") {
				// Verifica si es un dato de tipo unsigned
				uReal = true
				letter = true
			} else if (string(expresion[i]) == "E" || string(expresion[i]) == "e") && (!eReal) && ((containsArray(numeros, string(original[pos2+i+1]))) || string(original[pos2+i+1]) == "-") {
				// Verifica si se esta leyendo la "E" de la notación científica y que reciba un número o guión a continuación
				*operador = true
				eReal = true
				wait = true
			} else if string(expresion[i]) == "-" && (!guion) && eReal && (containsArray(numeros, string(original[pos2+i+1]))) {
				// Verifica si es un guión y que reciba un número a continuación
				guion = true
				wait = true
			} else if (string(expresion[i]) == "F" || string(expresion[i]) == "f") && (!wait) && (!uReal) && (lReal == 0) && eReal {
				// Verifica si es un fast int
				fReal = true
				letter = true
			} else {
				return false
			}
		}
	}
	return true
}

func resaltador(file string, dir string, iFile int) {
	// Lista que guardará el contenido del archivo TXT
	lista_sintaxis := []string{}

	// Abre el archivo de texto (sintaxis.txt)
	fileTxt, ferr := os.Open(dir + "\\Actividad 5.3\\" + file)
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(fileTxt)
	for scanner.Scan() {
		lista_sintaxis = append(lista_sintaxis, scanner.Text())
	}

	// Se abre o se crea un archivo html (index.html)
	fileHtml, e := os.Create(dir + "\\Actividad 5.3\\index" + fmt.Sprint(iFile) + ".html")
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
	comentarioLargo := false
	posComentarioLargo := 0
	originPos := 0

	// Lee cada enunciado del archivo de texto
	for i := 0; i < len(lista_sintaxis); i++ {
		// Acumulador de la expresion
		acumExp := ""
		// Variables para la indentación
		start := false
		espacio := ""
		// Variable para manejar espacios en literales de tipo string o char
		nullSpace := false
		// Variable para manejar "-" en los literales númericos
		operadorOmit := false
		// Variable para manejar las librerías
		libreria := false

		// Lee cada caracter del enunciado
		for j := 0; j < len(lista_sintaxis[i]); j++ {
			// Almacena el renglón
			expresion := lista_sintaxis[i]
			// Almacena un caracter del renglón
			token := expresion[j : j+1]
			// fmt.Println(token) // Imprime cada token

			// Condicional para realizar la indentación
			if !start {
				k := 0
				for expresion[k:k+1] == " " {
					espacio += "&nbsp;"
					k += 1
					if expresion[k:k+1] != " " {
						fileHtml.WriteString("\t\t<span>" + espacio + "</span>\n")
					}
				}
			}
			// Nos hace conocer que ya se leyó los primeros espacios en blanco de la expresión
			start = true

			// Si el valor actual del token es un espacio en blanco, a excepción de que se manejen comentarios largos
			// o literales de tipo string o char. Liberamos todas las variables que se tienen almacenadas si entra en la condicional.
			if token == " " && !nullSpace && !comentarioLargo && !libreria {
				// Si acumExp no esta vacío, significa que no pertenece a ninguna categoría léxica
				if acumExp != "" {
					fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
					acumExp = ""
				}
				// En caso contrario, concatenamos al acumulador los demás caracteres de la expresión,
				// a excepción del salto de línea, del tab, y del espacio en blanco en caso de requerirlo
			} else if token != "\n" && token != "\t" {
				acumExp += token
			}

			// Verificamos si se leerá a continuación comentarios largos
			if acumExp == "/" && expresion[j:j+2] == "/*" {
				// Almacenamos la posición de la lista de la expresión en donde se comenzó
				originPos = i
				// Ciclo que itera la cantidad de renglones del archivo
				for k := 0; k < len(lista_sintaxis)-i; k++ {
					exp := lista_sintaxis[i+k]
					// Busca el cierre del comentario largo
					if (len(exp) > 1 && strings.Index(exp[2:], "*/") != -1 && i == i+k) || (strings.Index(exp, "*/") != -1 && i != i+k) {
						if !comentarioLargo {
							// Almacena la posición de la lista en la que se encontro el cierre en la expresión
							posComentarioLargo = i + k
						}
						// Marcamos que se encontro el cierre
						comentarioLargo = true
					}
				}
				// Si no encontramos el cierre, marcamos todo como error hasta el final del archivo
				if !comentarioLargo {
					fileHtml.WriteString("\t\t<span class=\"error\">" + expresion[j:] + "</span>\n")
					// Ciclo que itera la cantidad de renglones del archivo
					for k := 0; k < len(lista_sintaxis)-i-1; k++ {
						// Despliega cada renglón como error de sintaxis
						fileHtml.WriteString("\t\t<br>\n")
						exp := lista_sintaxis[i+k+1]
						fileHtml.WriteString("\t\t<span class=\"error\">" + exp + "</span>\n")
					}
					// Da por terminado la lectura del archivo
					return
				}
			}

			//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

			// Verifica si se activo la lectura de comentarios largos
			if comentarioLargo {
				// Líneas de comentarios sin el cierre
				if originPos != i && i != posComentarioLargo && j == len(expresion)-1 {
					fileHtml.WriteString("\t\t<span class=\"comentario\">" + expresion + "</span>\n")
					acumExp = ""
					nullSpace = false
					// Primer línea de comentario y además cierra en la misma línea
				} else if len(acumExp) > 1 && strings.Index(acumExp[2:], "*/") != -1 && originPos == i && i == posComentarioLargo {
					fileHtml.WriteString("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
					comentarioLargo = false
					posComentarioLargo = 0
					originPos = 0
					// Línea de comentario diferente a la línea de apertura que encontro el cierre
				} else if strings.Index(acumExp, "*/") != -1 && i != originPos && posComentarioLargo == i {
					fileHtml.WriteString("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
					comentarioLargo = false
					posComentarioLargo = 0
					originPos = 0
					// En caso de ser la primera línea y que no tenga cierre
				} else if j == len(expresion)-1 {
					fileHtml.WriteString("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
				}
				// Verifica si es un comentario normal
			} else if acumExp != "" && isFile(acumExp) && !comentarioLargo && !libreria {
				// Busca si esta incluida la extensión del archivo en la expresión
				pos := strings.Index(expresion, ".cpp") + 4
				// Verifica que despues de validar que la extensión se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación o un caracter válido para desplegarla
				if j == len(expresion)-1 || expresion[pos:pos+1] == " " || pos == len(expresion)-1 || isDelimitador(expresion[pos:pos+1]) || isOperadorUnique(expresion[pos:pos+1], expresion, pos) || expresion[pos:pos+2] == "//" || expresion[pos:pos+2] == "/*" {
					fileHtml.WriteString("\t\t<span class=\"file\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
				}
				// Verifica si es un comentario normal
			} else if (acumExp != "" && isComentario(acumExp)) && !comentarioLargo && !libreria {
				fileHtml.WriteString("\t\t<span class=\"comentario\">" + expresion[j-1:] + "</span>\n")
				acumExp = ""
				nullSpace = false
				break
				// Verifica si es una librería - O(n^2)
			} else if acumExp != "" && !comentarioLargo && isLibreria(acumExp) {
				libreria = true
				posI := strings.Index(expresion, "include")
				posC := strings.Index(expresion, "\"")
				posCC := strings.Index(expresion[posC+1:], "\"")
				posF := strings.Index(expresion, "<")
				posFF := strings.Index(expresion, ">")
				posExp := strings.Index(expresion, acumExp)
				expSub := ""

				// Extrae de la expresión original una copia de la librería
				if posFF != -1 {
					expSub = expresion[posExp : posFF+1]
				} else if posCC != -1 {
					expSub = expresion[posExp : posCC+2+posC]
					posCC = posCC + 1 + posC
				}

				// Verifica que no haya errores
				for l := 0; l < posI-1; l++ {
					if expresion[l+1:l+2] != " " {
						libreria = false
					}
				}

				// Verifica que no haya errores
				if posC > posF {
					for l := posI + 7; l < posC; l++ {
						if expresion[l:l+1] != " " {
							libreria = false
						}
					}
				} else {
					for l := posI + 7; l < posF; l++ {
						if expresion[l:l+1] != " " {
							libreria = false
						}
					}
				}

				// Verifica que todo este correcto
				if posI > 0 && ((posC > posI && posCC > posC+1) || (posF > posI && posFF > posF+1)) {
					if expSub == acumExp {
						if posFF != -1 {
							fileHtml.WriteString("\t\t<span class=\"libreria\">" + acumExp[:posF+1] + " " + acumExp[posF+1:] + "</span>\n")
						} else {
							fileHtml.WriteString("\t\t<span class=\"libreria\">" + acumExp + "</span>\n")
						}
						acumExp = ""
						nullSpace = false
						libreria = false
					}
				} else if posCC == -1 || posFF == -1 || posI == -1 {
					fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
					libreria = false
				}
				// Verifica si es una palabra reservada
			} else if acumExp != "" && !comentarioLargo && isReservada(acumExp) {
				// Verifica que despues de validar que la palabra reservada se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación
				if j == len(expresion)-1 || expresion[j+1:j+2] == " " || isDelimitador(expresion[j+1:j+2]) || isOperadorUnique(expresion[j+1:j+2], expresion, j) || expresion[j+1:j+2] == "\"" || expresion[j+1:j+2] == "'" {
					fileHtml.WriteString("\t\t<span class=\"reservada\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
				}
				// Verifica si es un operador
			} else if acumExp != "" && !comentarioLargo && isOperador(acumExp, expresion, j) && !operadorOmit && !nullSpace {
				// Casos de excepción para marcar que son syntax error
				if len(acumExp) > 1 && (!isOperadorUnique(string(acumExp[0]), expresion, strings.Index(expresion, acumExp)) || isComentario(expresion[j:])) && !isOperadorUnique(acumExp, expresion, j) {
					fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
					operadorOmit = false
				}
				// Verifica una vez retirado la expresión erronea, hay un operador válido
				if j == len(expresion)-1 || string(expresion[j+1]) == " " || isOperadorUnique(string(expresion[j+1]), expresion, j) || isDelimitador(string(expresion[j+1])) || isIdentificador(string(expresion[j+1]), expresion, j) || isLiteral(string(expresion[j+1]), expresion, j, &operadorOmit) {
					if strings.Index(acumExp, "_") == -1 {
						fileHtml.WriteString("\t\t<span class=\"operador\">" + acumExp + "</span>\n")
						acumExp = ""
					} else {
						fileHtml.WriteString("\t\t<span class=\"operador\">" + acumExp[:strings.Index(acumExp, "_")] + "</span>\n")
						acumExp = acumExp[strings.Index(acumExp, "_"):]
					}
					nullSpace = false
				}
				// Verifica si es un delimitador
			} else if acumExp != "" && !comentarioLargo && isDelimitador(acumExp) && !nullSpace {
				if len(acumExp) > 1 && !isDelimitador(string(acumExp[0])) {
					fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
				}
				// Verifica que a continuación haya un caracter diferente válido para desplegarlo
				if j == len(expresion)-1 || string(expresion[j+1]) == " " || isOperadorUnique(string(expresion[j+1]), expresion, j) || isIdentificador(string(expresion[j+1]), expresion, j) || isLiteral(string(expresion[j+1]), expresion, j, &operadorOmit) || string(expresion[j+1]) == "." {
					if strings.Index(acumExp, "_") == -1 {
						fileHtml.WriteString("\t\t<span class=\"delimitador\">" + acumExp + "</span>\n")
						acumExp = ""
					} else {
						fileHtml.WriteString("\t\t<span class=\"delimitador\">" + acumExp[:strings.Index(acumExp, "_")] + "</span>\n")
						acumExp = acumExp[strings.Index(acumExp, "_"):]
					}
					nullSpace = false
				}
				// Verifica si es un identificador
			} else if acumExp != "" && !comentarioLargo && isIdentificador(acumExp, expresion, j) {
				// Verifica que a continuación haya un caracter diferente válido para desplegarlo
				if j == len(expresion)-1 || string(expresion[j+1]) == " " || isOperadorUnique(string(expresion[j+1]), expresion, j) || isDelimitador(string(expresion[j+1])) {
					fileHtml.WriteString("\t\t<span class=\"identificador\">" + acumExp + "</span>\n")
					acumExp = ""
					nullSpace = false
				}
				// Verifica si es una literal
			} else if acumExp != "" && !comentarioLargo && isLiteral(acumExp, expresion, j, &operadorOmit) {
				// Verifica que a continuación haya un caracter diferente válido para desplegarlo
				if j == len(expresion)-1 || string(expresion[j+1]) == " " || isDelimitador(string(expresion[j+1])) || (isOperadorUnique(string(expresion[j+1]), expresion, j) && string(expresion[j+1]) != "-") || (j+1 <= len(acumExp) && isComentario(string(acumExp[j+1:]))) || (string(acumExp[0]) == "'" && string(acumExp[len(acumExp)-1]) == "'") || (string(acumExp[0]) == "\"" && string(acumExp[len(acumExp)-1]) == "\"") {
					// Verifica que esten ambas comillas en la literal para poder habilitar la opción de leer espacios en blanco y otros caracteres
					if (string(acumExp[0]) == "'" && strings.Index(expresion[j+1:], "'") != -1) || (string(acumExp[0]) == "\"" && strings.Index(expresion[j+1:], "\"") != -1) {
						nullSpace = true
					}
					// Verifica que sea una literal de tipo string o char
					if (string(acumExp[0]) == "'" && string(acumExp[len(acumExp)-1]) == "'" && len(acumExp) != 1) || (string(acumExp[0]) == "\"" && string(acumExp[len(acumExp)-1]) == "\"" && len(acumExp) != 1) {
						fileHtml.WriteString("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
						acumExp = ""
						operadorOmit = false
						nullSpace = false
						// Verifica que sea una literal de tipo númerica
					} else if ((string(acumExp[0]) != "'") && (string(acumExp[len(acumExp)-1]) != "'")) && (string(acumExp[0]) != "\"" && string(acumExp[len(acumExp)-1]) != "\"") {
						fileHtml.WriteString("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
						acumExp = ""
						operadorOmit = false
						nullSpace = false
					}
				}
			}
		}
		// Si al final no se vacía el acumulador, es un syntax error, ya que no pertenece a ninguna categoría léxica
		if acumExp != "" {
			fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
		}

		// Escribimos saltos de línea cuando termine de leer un renglón por completo, por cuestiones de diseño del html
		fileHtml.WriteString("\t\t<br>\n")
	}
	// Escribimos el final del archivo html
	fileHtml.WriteString("\t</body>\n")
	fileHtml.WriteString("</html>")

	// Cerramos el archivo HTML
	fileHtml.Close()
	defer wg.Done()
}

func main() {
	// Definimos las variables que usaremos para abrir el archivo con la ayuda de la libreria OS.
	// Esto nos ayudará a evitar conflictos a la hora de abrirlo en equipos distintos.
	start := time.Now()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}

	// Abre el directorio en donde se encuentra este programa
	file, err := os.Open(dir + "\\Actividad 5.3")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	// Array que guardará los nombres de los archivos txt
	lista_file := []string{}

	// Para buscar los archivos txt en la carpeta y almacenarlos en el array
	list, _ := file.Readdirnames(0)
	for _, name := range list {
		if strings.Contains(name, ".txt") {
			lista_file = append(lista_file, name)
		}
	}

	// Ciclo para que procese varios archivos
	for iFile := 0; iFile < len(lista_file); iFile++ {
		wg.Add(1)
		go resaltador(lista_file[iFile], dir, iFile+1)
	}

	wg.Wait()
	sinceStart := time.Since(start)

	log.Printf("%s", sinceStart)
}
