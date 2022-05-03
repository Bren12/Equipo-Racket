'''
Actividad Integradora 3.4 - Resaltador de sintaxis

Fecha: 06-05-2020

Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
    - Brenda Elena Saucedo González | A00829855
'''

#Librerias a usar
import os

# Definimos las variables que usaremos para abrir el archivo con la ayuda de la libreria OS.
# Esto nos ayudara a evitar conflictos a la hora de abrirlo en equipos distintos.
folder_actual = os.path.dirname(os.path.abspath(__file__))
nombre_archivo_texto = os.path.join(folder_actual, "sintaxis.txt")
nombre_archivo_html = os.path.join(folder_actual, "index.html")

# Ejemplo: int a = 5 // a es igual a 5
def isFile(expresion):

    # Verifica si existen espacios
    pos = expresion.find(" ")

    # Si no hay espacios, procedemos a continuar
    if (pos == -1):

        # Busca si esta incluida la extensión del archivo en la expresión y retorna su debida posición
        pos = expresion.find(".cpp")
        if (pos != -1):
            print(expresion + " " + str(pos))
            return True

    return False

def isComentario(expresion):
    # Busca si hay "//" en la expresión y retorna su debida posición
    pos = expresion.find("//")
    if (pos == 0):
        return True

    return False

# Definición de función principal
def main():
    # Se abre el archivo html (index.html) y se empieza a escribir en el
    with open(nombre_archivo_html,"w") as file:

        # Escribimos el head del archivo html
        file.write("<!DOCTYPE html>\n")
        file.write("<html>\n")
        file.write("\t<head>\n")
        file.write("\t\t<meta charset=\"utf-8\"/>\n")
        file.write("\t\t<title>Resaltador de Sintaxix</title>\n")
        file.write("\t\t<link rel=\"stylesheet\" href=\"style.css\">\n")
        file.write("\t</head>\n")
        file.write("\t<body>\n")
        
        # Abre el archivo de texto (sintaxis.txt)
        with open (nombre_archivo_texto, mode = 'r', encoding = "utf-8") as archivo:
            # Lee todo el archivo y lo deja en una lista de strings
            lista_sintaxis = archivo.readlines()

        # Lee cada enunciado del archivo de texto
        for i, expresion in enumerate(lista_sintaxis):
            
            # Acumulador de la expresion 
            acumExp = ""
            # Acumulador de lo que se tiene que escribir en el html
            acumHTML = ""

            # Lee cada caracter del enunciado
            for j, token in enumerate(expresion):

                # Concatenamos al acumulador los demás caracteres de la expresión
                acumExp = acumExp + token

                if (isFile(acumExp)):
                    file.write("\t\t<span class=\"file\">" + acumExp + "</span>\n")
                    acumExp = ""
                    acumHTML = ""
                if (isComentario(acumExp)):
                    file.write("\t\t<span class=\"comentario\">" + expresion[j-1:-1] + "</span>\n")
                    acumExp = ""
                    acumHTML = ""
                    break
                    # isLibreria(expresion) # Diego
                    # isReservada(expresion) # Jose Angel
                    # isLiteral(expresion) # Brenda
                    # isOperador(expresion) # Brenda
                    # isDelimitador(expresion) # Jose Angel
                    # isIdentificador(expresion) # Diego

            # Si al final no se vacía el acumulador, es un syntax error
            if (acumExp != "" and "\n"):
                file.write("\t\t<span class=\"error\">" + acumExp[:-1] + "</span>\n")

            # Escribimos el salto de línea cuando termine de leer un renglón por completo
            file.write("\t\t<br>\n")

        # Escribimos el final del archivo html
        file.write("    </body>\n")
        file.write("</html>")

# Se llama a la función principal
main()