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

	/////////////////////////////////////// AQUÍ IRA LAS DEMAS FUNCIONES DE VERIFICACION DE SINTAXIX ///////////////////////////////////////
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
			expresion := lista_sintaxis[i]
			token := expresion[j : j+1]
			fmt.Println(token)

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
			start = true

			if (token == " " && !nullSpace && !comentarioLargo && !libreria) || (j == len(lista_sintaxis[i])-1) {
				// Si acumExp no esta vacío, significa que no pertenece a ninguna categoría léxica
				if acumExp != "" {
					fileHtml.WriteString("\t\t<span class=\"error\">" + acumExp + "</span>\n")
					acumExp = ""
				}
				// fileHtml.WriteString("\t\t<span>" + acumExp + "</span>\n")
				// fileHtml.WriteString("\t\t<br>\n")
			} else if token != "\n" && token != "\t" {
				acumExp += token
			}

			if acumExp == "/" && expresion[j:j+2] == "/*" {
				originPos = i
				for k := 0; k < len(lista_sintaxis)-i; k++ {
					exp := lista_sintaxis[i+k]
					if (len(exp) >= 1 && strings.Index(exp[2:], "*/") != -1 && i == i+k) || (strings.Index(exp, "*/") != -1 && i != i+k) {
						if !comentarioLargo {
							posComentarioLargo = i + k
						}
						comentarioLargo = true
					}
				}
				if !comentarioLargo {
					fileHtml.WriteString("\t\t<span class=\"error\">" + expresion[j:] + "</span>\n")
					for k := 0; k < len(lista_sintaxis)-i-1; k++ {
						fileHtml.WriteString("\t\t<br>\n")
						exp := lista_sintaxis[i+k+1]
						fileHtml.WriteString("\t\t<span class=\"error\">" + exp + "</span>\n")
					}
					return
				}
				fmt.Println(originPos)
				fmt.Println(posComentarioLargo)
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
