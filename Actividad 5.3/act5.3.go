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
	"strings"
)

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

func isOperadorUnique(expresion string, original string, pos int) bool {
	// Se definen los operadores como un diccionario
	operador := []string{"+", "+=", "++", "-", "-=", "--", "%", "%=", "*", "*=", "/=", "^", "<", "<<", ">", ">>",
		"<=", ">=", "=", "==", "!", "!=", "~", "?", "&", "&&", "||"}

	// Verifica si existe en la expresión cualquier operador, si no, retorna falso
	// Para eso, primero se verifica que no vaya a sobrepasar la longitud de la expresión original
	if pos+2 < len(original) {
		// Verifica que no sea un comentario lo que se esta leyendo
		if (expresion == "/" && string(original[pos:pos+2]) != "//") && (expresion == "/" && original[pos:pos+2] != "/*") {
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
		// // Variable para manejar "-" en los literales númericos
		// operadorOmit := false
		// Variable para manejar las librerías
		libreria := false

		// Lee cada caracter del enunciado
		for j := 0; j < len(lista_sintaxis[i]); j++ {
			// Almacena el renglón
			expresion := lista_sintaxis[i]
			// Almacena un caracter del renglón
			token := expresion[j : j+1]
			fmt.Println(token)

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
					if (len(exp) >= 1 && strings.Index(exp[2:], "*/") != -1 && i == i+k) || (strings.Index(exp, "*/") != -1 && i != i+k) {
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
			}

		}
		fileHtml.WriteString("\t\t<br>\n")
	}
	// Escribimos el final del archivo html
	fileHtml.WriteString("\t</body>\n")
	fileHtml.WriteString("</html>")

	// Cerramos el archivo HTML
	fileHtml.Close()
}
