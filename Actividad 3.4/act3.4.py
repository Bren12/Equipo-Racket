'''
Actividad Integradora 3.4 - Resaltador de sintaxis
Fecha: 06-05-2020
Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
    - Brenda Elena Saucedo González | A00829855
En el presente programa se definen categorías léxicas pertenecientes al lenguaje de programación de C++
'''

#Librerias a usar
import os
import string

# Definimos las variables que usaremos para abrir el archivo con la ayuda de la libreria OS.
# Esto nos ayudara a evitar conflictos a la hora de abrirlo en equipos distintos.
folder_actual = os.path.dirname(os.path.abspath(__file__))
nombre_archivo_texto = os.path.join(folder_actual, "sintaxis.txt")
nombre_archivo_html = os.path.join(folder_actual, "index.html")

def isFile(expresion):

    # Busca si esta incluida la extensión del archivo en la expresión
    pos = expresion.find(".cpp")

    # Se definen algunos caracteres especiales que no son permitidos en los nombres de archivos
    caracteres = ["\\","/",":","*","?","<",">","|"]

    # Retorna verdadero si encontro la extensión en la expresión
    if (pos > 0):
        aux = expresion[pos+4:]
        for i, caract in enumerate(caracteres):
            for j, variable in enumerate(expresion[:pos]):
                if (variable == caract):
                    return False
        return True

    return False

def isComentario(expresion):
    # Busca si hay "//" en la expresión
    pos = expresion.find("//")

    # Retorna verdadero si encontro "//" en la expresión
    if (pos == 0):
        return True

    return False

def isLibreria(expresion):
    #Busca el # que en C++ indica una libreria a incluir
    pos = expresion.find("#")
    #Dado caso de que la encuentre marcalo como verdadero
    if (pos == 0):
        return True
    return False

def isReservada(expresion):
    # Se definen las palabras reservadas como un diccionario
    reservada = {"int": True, "bool": True, "char": True, "void": True, "float": True, "double": True, "string": True, "cin": True, "cout": True, "while": True, 
    "as": True, "using": True, "namespace": True, "auto": True, "const": True, "asm": True, "dynamic_cast": True, "reinterpret_cast": True, "try": True, 
    "explicit": True, "new": True, "static_cast": True, "static": True, "typeid": True, "catch": True, "false": True, "operator": True, "template": True, 
    "typename": True, "class": True, "friend": True, "private": True, "this": True, "const_cast": True, "inline": True, "public": True, "throw": True, 
    "virtual": True, "delete": True, "enum": True, "goto": True, "else": True, "mutable": True, "protected": True, "true": True, "wchar_t": True,
    "sizeof": True, "register": True, "unsigned": True, "break": True, "continue": True, "extern": True, "if": True, "return": True, "switch": True, "case": True,
    "default": True, "short": True, "struct": True, "volatile": True, "do": True, "for": True, "long": True, "signed": True, "union": True, "std": True,}
    
    # Verifica si existe en la expresión cualquier palabra reservada, si no, retorna falso
    try:
        return reservada[expresion]
    except:
        return False

def isOperador(expresion, original, pos):
    # Se definen los operadores como una lista
    operador = ["+", "+=", "++", "-", "-=", "--", "%", "%=", "*", "*=", "/=", "^", "<", "<<", ">", ">>", "<=", ">=", "=", "==", "!", "!=", "~", "?", ":", "&", "&&", "||"]

    # Ciclo que itera cada operador de la lista definida
    for i, op in enumerate(operador):
        # Si encontro el operador, se retorna verdadero
        if (op == expresion or expresion.find(op) != -1 or (expresion == "/" and original[pos:pos+2] != "//") or (expresion.find("/") != -1 and original[pos:pos+2] != "//")):
            return True

    return False

def isOperadorUnique(expresion, original, pos):
    # Se definen los operadores como un diccionario
    operador = {"+": True, "+=": True, "++": True, "-": True, "-=": True, "--": True, "%": True, "%=": True, "*": True, "*=": True, "/=": True, "^": True, "<": True, "<<": True, ">": True, ">>": True, 
    "<=": True, ">=": True, "=": True, "==": True, "!": True, "!=": True, "~": True, "?": True, ":": True, "&": True, "&&": True, "||": True}

    # Verifica si existe en la expresión cualquier operador, si no, retorna falso
    try:
        if (pos+2 < len(original)):
            if (expresion == "/" and original[pos:pos+2] != "//"):
                return True
            else:
                return operador[expresion]
    except:
        return False

def isDelimitador(expresion):
    delimitador = ["(", ")", "[", "]", "{", "}", ",", ";", "...", "*", "="]

    # Ciclo que itera cada operador de la lista definida
    for i, op in enumerate(delimitador):
        # Si encontro el operador, se retorna verdadero
        if (op == expresion or expresion.find(op) != -1):
            return True

    return False

def isIdentificador(expresion, original, pos):
    # Se crea una lista para checar todos los identificadores con letras
    alfabeto = list(string.ascii_letters)
    # Se crea una lista para números
    numeros = []
    #Añade los numeros del 0 al 9
    for x in range(0, 10):
        numeros.append(str(x))
    for i, op in enumerate(alfabeto):
        # Si encontro el identificador, se retorna verdadero
        if (op == expresion or expresion.find(op) != -1 or (expresion in numeros) or "_" in expresion):
            #Checa casos de excepcion que indican que no es un identificador
            if ((expresion[0] in alfabeto) and (not ("\"" in expresion or "\'" in expresion or "." in expresion))):
                return True
    return False

def isLiteral(expresion, original, pos, operador):
    numeros = []
    punto = False
    guion = False
    eReal = False
    fReal = False
    uReal = False
    lReal = 0
    wait = False
    letter = False
    
    if (expresion[0] == "\'" or expresion[0] == "\""):
        return True

    for x in range(0, 10):
        numeros.append(str(x))

    try:
        pos2 = original.find(expresion)
        for i, x in enumerate(expresion):
            if (expresion[i] in numeros and not letter):
                wait = False
            else:
                if (expresion[i] == "." and not punto):
                    punto = True
                elif ((expresion[i] == "L" or expresion[i] == "l") and not wait and ((lReal < 2 and not uReal) or (uReal and lReal == 0)) and not fReal):
                    lReal = lReal + 1
                    letter = True
                elif ((expresion[i] == "U" or expresion[i] == "u") and not uReal and not eReal and original[pos2+i+1] != "."):
                    uReal = uReal + 1
                    letter = True
                elif ((expresion[i] == "E" or expresion[i] == "e") and not eReal and (original[pos2+i+1] in numeros or original[pos2+i+1] == "-")):
                    operador[0] = True
                    eReal = True
                    wait = True
                elif (expresion[i] == "-" and not guion and eReal and original[pos2+i+1] in numeros):
                    guion = True
                    wait = True
                elif ((expresion[i] == "F" or expresion[i] == "f") and not wait and not uReal and lReal == 0 and eReal):
                    fReal = fReal + 1
                    letter = True
                else:
                    return False

        return True
    except:
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
            acumHTML = list()
            # Variables para la indentación
            start = False
            espacio = ""
            nullSpace = False
            operadorOmit = [False]

            # Lee cada caracter del enunciado
            for j, token in enumerate(expresion):

                # Condicional para realizar la indentación
                if (not start):
                    k = 0
                    while(expresion[k] == " "):
                        espacio = espacio + "&nbsp;"
                        k = k + 1
                        if (expresion[k] != " "):
                            file.write("\t\t<span>" + espacio + "</span>\n")
                start = True

                # Si el valor actual del token es un espacio en blanco, liberamos todas las variables que tienen almacenados algún valor
                if (token == " " and not nullSpace):
                    # Verifica que no haya quedado nada en la lista acumHTML, sino, las despliega todas en el archivo HTML
                    for x in range(len(acumHTML)):
                        file.write(acumHTML[x])
                    # Si acumExp no esta vacío, significa que no pertenece a ninguna categoría léxica
                    if (acumExp[:j] != ""):
                        file.write("\t\t<span class=\"error\">" + acumExp[:j] + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]
                # Concatenamos al acumulador los demás caracteres de la expresión, a excepción del salto de línea y del espacio en blanco
                elif (token != "\n"):
                    acumExp = acumExp + token

                if (acumExp != "" and isFile(acumExp)):
                    # Busca si esta incluida la extensión del archivo en la expresión
                    pos = expresion.find(".cpp") + 4
                    # Verifica que despues de validar que la extensión se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación
                    if (j == len(expresion)-1 or expresion[pos] == " " or expresion[pos] == "\n" or isDelimitador(expresion[pos]) or isOperadorUnique(expresion[pos],expresion,pos) or expresion[pos:pos+2] == "//"):
                        file.write("\t\t<span class=\"file\">" + acumExp + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]
                elif (acumExp != "" and isComentario(acumExp)):
                    file.write("\t\t<span class=\"comentario\">" + expresion[j-1:-1] + "</span>\n")
                    acumExp = ""
                    del acumHTML [:]
                    break
                elif (acumExp != "" and isLibreria(acumExp)):
                    file.write("\t\t<span class=\"libreria\">" + expresion + "</span>\n")
                    acumExp = ""
                    del acumHTML [:]
                    break
                elif (acumExp != "" and isReservada(acumExp)):
                    # Verifica que despues de validar que la palabra reservada se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isDelimitador(expresion[j+1]) or isOperadorUnique(expresion[j+1],expresion,j)):
                        file.write("\t\t<span class=\"reservada\">" + acumExp + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]

                elif (acumExp != "" and isOperador(acumExp,expresion,j) and not operadorOmit[0]):
                    enter = True
                    # Verifica que no hayan otros valores antes del operador en la expresión, sino, libera esa parte como un error de sintaxis a excepción del operador
                    if (len(acumExp) != 1 and acumExp[:j] != "" and not isOperadorUnique(acumExp,expresion,j)):
                        if (expresion[j:j+2] == "//"):
                            file.write("\t\t<span class=\"operador\">" + acumExp[:-1] + "</span>\n")
                            acumExp = expresion[j]
                            del acumHTML [:]
                            enter = False
                        elif (isOperadorUnique(acumExp[0],expresion,j) or acumExp[0].isdigit() or acumExp[0] == "."):
                            acumExp = acumExp
                        else:
                            file.write("\t\t<span class=\"error\">" + acumExp[:j] + "</span>\n")
                            acumExp = acumExp[j:]
                            del acumHTML [:]
                    # Verifica que a continuación se encuentre cualquier otro valor que no sea un operador
                    if (not isOperadorUnique(expresion[j+1],expresion,j) and enter):
                        file.write("\t\t<span class=\"operador\">" + acumExp + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]

                elif (acumExp != "" and isDelimitador(acumExp) and not operadorOmit[0]):
                    if (len(acumExp) > 1 and not isDelimitador(acumExp[0])):
                        file.write("\t\t<span class=\"error\">" + acumExp[:-1] + "</span>\n")
                        acumExp = acumExp[-1]
                        del acumHTML [:]
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isOperadorUnique(expresion[j+1],expresion,j) or isIdentificador(expresion[j+1],expresion,j) or isLiteral(expresion[j+1],expresion,j,operadorOmit)):
                        file.write("\t\t<span class=\"delimitador\">" + acumExp + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]

                elif (acumExp != "" and isIdentificador(acumExp, expresion, j)):
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isOperadorUnique(expresion[j+1],expresion,j) or isDelimitador(expresion[j+1])):
                        file.write("\t\t<span class=\"identificador\">" + acumExp + "</span>\n")
                        acumExp = ""
                        del acumHTML [:]
                
                elif (acumExp != "" and isLiteral(acumExp, expresion, j, operadorOmit)):
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isDelimitador(expresion[j+1]) or (isOperadorUnique(expresion[j+1],expresion,j) and expresion[j+1] != "-") or isComentario(acumExp[j+1:])):
                        if ((acumExp[0] == "\'" and expresion[j+1:].find("\'") != -1) or (acumExp[0] == "\"" and expresion[j+1:].find("\"") != -1)):
                            nullSpace = True
                            
                        if ((acumExp[0] == "\'" and acumExp[len(acumExp)-1] == "\'") or (acumExp[0] == "\"" and acumExp[len(acumExp)-1] == "\"")):
                            file.write("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
                            acumExp = ""
                            del acumHTML [:]
                            operadorOmit[0] = False
                            nullSpace = False

                        elif(((acumExp[0] != "\'") and (acumExp[len(acumExp)-1] != "\'")) and (acumExp[0] != "\"" and acumExp[len(acumExp)-1] != "\"")):
                            file.write("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
                            acumExp = ""
                            del acumHTML [:]
                            operadorOmit[0] = False
                            nullSpace = False

            for i in range(len(acumHTML)):
                file.write(acumHTML[i])

            # Si al final no se vacía el acumulador, es un syntax error, ya que no pertenece a ninguna categoría léxica
            if (acumExp != ""):
                file.write("\t\t<span class=\"error\">" + acumExp + "</span>\n")

            # Escribimos saltos de línea cuando termine de leer un renglón por completo, por cuestiones de diseño del html
            file.write("\t\t<br>\n")

        # Escribimos el final del archivo html
        file.write("\t</body>\n")
        file.write("</html>")

# Se llama a la función principal
main()