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

# Definimos una función que verifica si el parametro recibido es un archivo .cpp
# Complejidad: O(n), porque aunque se manejan 2 ciclos, uno de ellos siempre van a ser constantes (caracteres), su valor no es variable, ni dependen de un parametro.
#              Por otro lado, el 2do ciclo maneja un valor que no es constante, que depende de la longitud del parametro recibido, el cual si es variable.
def isFile(expresion):
    # Busca si esta incluida la extensión del archivo en la expresión
    pos = expresion.find(".cpp")
    # Se definen algunos caracteres especiales que no son permitidos en los nombres de archivos
    caracteres = ["\\","/",":","*","?","<",">","|"]
    # Retorna verdadero si encontro la extensión en la expresión
    if (pos > 0):
        # Un ciclo para recorrer la lista de caracteres
        for i, caract in enumerate(caracteres):
            # Un ciclo para recorrer la expresión antes del ".cpp"
            for j, variable in enumerate(expresion[:pos]):
                # Verifica que no este incluido un caracter especial no permitido en el nombre del archivo
                if (variable == caract):
                    return False
        # Retorna verdadero ya que no encontro un caracter especial
        return True
    # Retorna falso si no encontro la extensión en la expresión
    return False

# Definimos una función que verifica si el parametro recibido es un comentario
# Complejidad: O(1), porque no realiza ningun ciclo o recursión, lee las instrucciones 1 sola vez.
def isComentario(expresion):
    # Busca si hay "//" en la expresión
    pos = expresion.find("//")
    # Retorna verdadero si encontro "//" en la expresión
    if (pos == 0):
        return True
    # Retorna falso, en caso contrario
    return False

# Definimos una función que verifica si el parametro recibido es una librería
# Complejidad: O(1), porque no realiza ningun ciclo o recursión, lee las instrucciones 1 sola vez.
def isLibreria(expresion):
    #Busca el # que en C++ indica una libreria a incluir
    pos = expresion.find("#")
    #Dado caso de que la encuentre marcalo como verdadero
    if (pos == 0):
        return True
    return False

# Complejidad: O(1), porque no realiza ningun ciclo o recursión, lee las instrucciones 1 sola vez.
def isReservada(expresion):
    # Se definen las palabras reservadas como un diccionario
    reservada = {"int": True, "bool": True, "char": True, "void": True, "float": True, "double": True, "string": True, "cin": True, "cout": True, "while": True, 
    "as": True, "using": True, "namespace": True, "auto": True, "const": True, "asm": True, "dynamic_cast": True, "reinterpret_cast": True, "try": True, 
    "explicit": True, "new": True, "static_cast": True, "static": True, "typeid": True, "catch": True, "false": True, "operator": True, "template": True, 
    "typename": True, "class": True, "friend": True, "private": True, "this": True, "const_cast": True, "inline": True, "public": True, "throw": True, 
    "virtual": True, "delete": True, "enum": True, "goto": True, "else": True, "mutable": True, "protected": True, "true": True, "wchar_t": True, "endl": True,
    "sizeof": True, "register": True, "unsigned": True, "break": True, "continue": True, "extern": True, "if": True, "return": True, "switch": True, "case": True,
    "default": True, "short": True, "struct": True, "volatile": True, "do": True, "for": True, "long": True, "signed": True, "union": True, "std": True,}
    # Verifica si existe en la expresión cualquier palabra reservada, si no, retorna falso
    try:
        return reservada[expresion]
    except:
        return False

# Definimos una función que verifica si el parametro recibido (expresion) contiene un operador.
# Complejidad: O(1), porque aunque se maneja 1 ciclo, este siempre va a ser constante, su valor no es variable, ni depende de un parametro.
def isOperador(expresion, original, pos):
    # Se definen los operadores como una lista
    operador = ["+", "+=", "++", "-", "-=", "--", "%", "%=", "*", "*=", "/=", "^", "<", "<<", ">", ">>", "<=", ">=", "=", "==", "!", "!=", "~", "?", "&", "&&", "||"]
    # Ciclo que itera cada operador de la lista definida
    for i, op in enumerate(operador):
        # Si encontro el operador, se retorna verdadero
        if (op == expresion or expresion.find(op) != -1 or (expresion == "/" and original[pos:pos+2] != "//") or (expresion.find("/") != -1 and original[pos:pos+2] != "//")):
            return True
    return False

# Definimos una función que verifica si el parametro recibido (expresion) es un operador.
# Complejidad: O(1), porque no realiza ningun ciclo o recursión, lee las instrucciones 1 sola vez.
def isOperadorUnique(expresion, original, pos):
    # Se definen los operadores como un diccionario
    operador = {"+": True, "+=": True, "++": True, "-": True, "-=": True, "--": True, "%": True, "%=": True, "*": True, "*=": True, "/=": True, "^": True, "<": True, "<<": True, ">": True, ">>": True, 
    "<=": True, ">=": True, "=": True, "==": True, "!": True, "!=": True, "~": True, "?": True, "&": True, "&&": True, "||": True}
    # Verifica si existe en la expresión cualquier operador, si no, retorna falso
    try:
        # Verifica que no vaya a sobrepasar la longitud de la expresión original
        if (pos+2 < len(original)):
            # Verifica que no sea un comentario lo que se esta leyendo
            if ((expresion == "/" and original[pos:pos+2] != "//") and (expresion == "/" and original[pos:pos+2] != "/*")):
                return True
            # Si encuentra un operador retorna verdadero
            else:
                return operador[expresion]
    # En caso contrario, retorna falso
    except:
        return False

# Definimos una función que verifica si el parametro recibido (expresion) es un delimitador.
# Complejidad: O(1), porque aunque se maneja 1 ciclo, este siempre va a ser constante, su valor no es variable, ni depende de un parametro.
def isDelimitador(expresion):
    # Se definen los delimitadores como una lista
    delimitador = ["(", ")", "[", "]", "{", "}", ",", ";", "...", ":"]
    # Ciclo que itera cada delimitador de la lista definida
    for i, delim in enumerate(delimitador):
        # Si encontro el delimitador, se retorna verdadero
        if (delim == expresion or expresion.find(delim) != -1):
            return True
    return False

# Definimos una función que verifica si el parametro recibido (expresion) es un identificador.
# Complejidad: O(1), porque aunque se manejan 2 ciclos, estos siempre van a ser constantes, sus valores no son variables, ni dependen de un parametro.
def isIdentificador(expresion, original, pos):
    # Se crea una lista para checar todos los identificadores con letras
    alfabeto = list(string.ascii_letters)
    # Se crea una lista para números
    numeros = []
    # Añade a la lista los numeros del 0 al 9
    for x in range(0, 10):
        numeros.append(str(x))
    # Ciclo que itera cada letra del alfabeto
    for i, letra in enumerate(alfabeto):
        # Si encontro una letra del afabeto o un guión entra a la siguiente condicional
        if (letra == expresion or expresion.find(letra) != -1 or (expresion in numeros) or "_" in expresion):
            # Checa casos de excepcion que indican que no es un identificador
            if ((expresion[0] in alfabeto) and (not ("\"" in expresion or "\'" in expresion or "." in expresion or "#" in expresion))):
                return True
    return False

# Definimos una función que verifica si el parametro recibido (expresion) es una literal.
# Complejidad: O(n), porque se tiene un ciclo que maneja un parametro que es variable, puesto que depende de la longitud de la expresion.
def isLiteral(expresion, original, pos, operador):
    # Se declaran variables para manejar los casos de excepción que nos indican que no son literales
    numeros = []
    punto = False
    guion = False
    eReal = False
    fReal = False
    uReal = False
    lReal = 0
    wait = False
    letter = False
    
    # Si se encontro una o doble comilla, retorna verdadero, ya que se esta por leer un string o char
    if (expresion[0] == "\'" or expresion[0] == "\""):
        return True

    # Añade los numeros del 0 al 9
    for x in range(0, 10):
        numeros.append(str(x))

    try:
        # Guarda la posicion en la que se encuentra el inicio de la expresion
        pos2 = original.find(expresion)
        # Ciclo que itera cada caracter de la expresion
        for i, x in enumerate(expresion):
            # Verifica que sea un número
            if (expresion[i] in numeros and not letter):
                wait = False
            else:
                # Verifica si es real
                if (expresion[i] == "." and not punto and original[pos2+i+1] in numeros):
                    punto = True
                # Verifica si es un dato de tipo long o long long
                elif ((expresion[i] == "L" or expresion[i] == "l") and not wait and ((lReal < 2 and not uReal) or (uReal and lReal == 0) or expresion[:i].find("ul") != -1) and not fReal):
                    lReal = lReal + 1
                    letter = True
                # Verifica si es un dato de tipo unsigned
                elif ((expresion[i] == "U" or expresion[i] == "u") and not uReal and not punto and not eReal and original[pos2+i+1] != "."):
                    uReal = uReal + 1
                    letter = True
                # Verifica si se esta leyendo la "E" de la notación científica y que reciba un número o guión a continuación
                elif ((expresion[i] == "E" or expresion[i] == "e") and not eReal and (original[pos2+i+1] in numeros or original[pos2+i+1] == "-")):
                    operador[0] = True
                    eReal = True
                    wait = True
                # Verifica si es un guión y que reciba un número a continuación
                elif (expresion[i] == "-" and not guion and eReal and original[pos2+i+1] in numeros):
                    guion = True
                    wait = True
                # Verifica si es un fast int
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

        # Definimos variables para manejar comentarios largos
        comentarioLargo = [False]
        posComentarioLargo = 0
        originPos = 0

        # Lee cada enunciado del archivo de texto
        # Complejidad: O(r)
        for i, expresion in enumerate(lista_sintaxis):
            # Acumulador de la expresion 
            acumExp = ""
            # Variables para la indentación
            start = False
            espacio = ""
            # Variable para manejar espacios en literales de tipo string o char
            nullSpace = False
            # Variable para manejar "-" en los literales númericos
            operadorOmit = [False]
            # Variable para manejar las librerías
            libreria = [False]

            # Lee cada caracter del enunciado
            # Complejidad: O(r), donde r es la cantidad de renglones en el archivo
            for j, token in enumerate(expresion):

                # Condicional para realizar la indentación
                if (not start):
                    k = 0
                    # Complejidad: O(n), ya que depende de cuantos espacios en blanco tenga la expresión.
                    while(expresion[k] == " "):
                        espacio = espacio + "&nbsp;"
                        k = k + 1
                        if (expresion[k] != " "):
                            file.write("\t\t<span>" + espacio + "</span>\n")
                # Nos hace conocer que ya se leyó los primeros espacios en blanco de la expresión
                start = True

                # Si el valor actual del token es un espacio en blanco, a excepción de que se manejen comentarios largos
                # o literales de tipo string o char. Liberamos todas las variables que se tienen almacenadas si entra en la condicional.
                if (token == " " and not nullSpace and not comentarioLargo[0] and not libreria[0]):
                    # Si acumExp no esta vacío, significa que no pertenece a ninguna categoría léxica
                    if (acumExp[:j] != ""):
                        file.write("\t\t<span class=\"error\">" + acumExp[:j] + "</span>\n")
                        acumExp = ""
                # En caso contrario, concatenamos al acumulador los demás caracteres de la expresión, 
                # a excepción del salto de línea, del tab, y del espacio en blanco en caso de requerirlo
                elif (token != "\n" and token != "\t"):
                    acumExp = acumExp + token

                # Verificamos si se leerá a continuación comentarios largos
                if (acumExp == "/" and expresion[j:j+2] == "/*"):
                    # Almacenamos la posición de la lista de la expresión en donde se comenzó
                    originPos = i
                    # Ciclo que itera la cantidad de renglones del archivo
                    # Complejidad: O(r)
                    for k in range(len(lista_sintaxis)- i):
                        exp = lista_sintaxis[i+k]
                        # Busca el cierre del comentario largo
                        if ((exp[2:].find("*/") != -1 and i == i+k) or (exp.find("*/") != -1 and i != i+k)):
                            if (not comentarioLargo[0]):
                                # Almacena la posición de la lista en la que se encontro el cierre en la expresión
                                posComentarioLargo = i+k
                            # Marcamos que se encontro el cierre
                            comentarioLargo[0] = True
                    # Si no encontramos el cierre, marcamos todo como error hasta el final del archivo
                    if (not comentarioLargo[0]):
                        file.write("\t\t<span class=\"error\">" + expresion[j:-1] + "</span>\n")
                        # Ciclo que itera la cantidad de renglones del archivo
                        # Complejidad: O(r)
                        for k in range(len(lista_sintaxis) - i - 1):
                            # Despliega cada renglón como error de sintaxis
                            file.write("\t\t<br>\n")
                            exp = lista_sintaxis[i+k+1]
                            if (exp[-1] == "\n"):
                                exp[:-1]
                            file.write("\t\t<span class=\"error\">" + exp + "</span>\n")
                        # Da por terminado la lectura del archivo
                        return

                # Verifica si se activo la lectura de comentarios largos - O(1)
                if (comentarioLargo[0]):
                    # Líneas de comentarios sin el cierre
                    if (originPos != i and i != posComentarioLargo and j == len(expresion)-1):
                        file.write("\t\t<span class=\"comentario\">" + expresion[:-1] + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                    # Primer línea de comentario y además cierra en la misma línea
                    elif (acumExp[2:].find("*/") != -1 and originPos == i and i == posComentarioLargo):
                        file.write("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                        comentarioLargo[0] = False
                        posComentarioLargo = 0
                        originPos = 0
                    # Línea de comentario diferente a la línea de apertura que encontro el cierre
                    elif (acumExp.find("*/") != -1 and i !=  originPos and posComentarioLargo == i):
                        file.write("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                        comentarioLargo[0] = False
                        posComentarioLargo = 0
                        originPos = 0
                    # En caso de ser la primera línea y que no tenga cierre
                    elif (j == len(expresion)-1):
                        file.write("\t\t<span class=\"comentario\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                # Verifica si es un archivo - O(n)
                elif (acumExp != "" and not comentarioLargo[0] and isFile(acumExp) and not libreria[0]):
                    # Busca si esta incluida la extensión del archivo en la expresión
                    pos = expresion.find(".cpp") + 4
                    # Verifica que despues de validar que la extensión se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación o un caracter válido para desplegarla
                    if (j == len(expresion)-1 or expresion[pos] == " " or expresion[pos] == "\n" or isDelimitador(expresion[pos]) or isOperadorUnique(expresion[pos],expresion,pos) or expresion[pos:pos+2] == "//" or expresion[pos:pos+2] == "/*"):
                        file.write("\t\t<span class=\"file\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                # Verifica si es un comentario normal - O(1)
                elif ((acumExp != "" and isComentario(acumExp)) and not comentarioLargo[0] and not libreria[0]):
                        file.write("\t\t<span class=\"comentario\">" + expresion[j-1:-1] + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                        break
                # Verifica si es una librería - O(n^2)
                elif (acumExp != "" and not comentarioLargo[0] and isLibreria(acumExp)):
                    # Variables para identificar que este correctamente la librería declarada
                    libreria[0] = True
                    noEs = False
                    posI = expresion.find("include")
                    posC = expresion.find("\"")
                    posCC = expresion[posC+1:].find("\"")
                    posF = expresion.find("<")
                    posFF = expresion.find(">")
                    posExp = expresion.find(acumExp)
                    expSub = ""
                    # Extrae de la expresión original una copia de la librería
                    if (posFF != -1):
                        expSub = expresion[posExp:posFF+1]
                    elif (posCC != -1):
                        expSub = expresion[posExp:posCC+2+posC]
                        posCC = posCC+1+posC
                    # Verifica que no haya errores
                    for l in range(posI-1):
                        if (expresion[l+1] != " "):
                            noEs = True
                            libreria[0] = False
                    # Verifica que no haya errores
                    if (posC > posF):
                        for l in range(posI+7,posC):
                            if (expresion[l] != " "):
                                noEs = True
                                libreria[0] = False
                    else:
                        for l in range(posI+7,posF):
                            if (expresion[l] != " "):
                                noEs = True
                                libreria[0] = False
                    # Verifica que todo este correcto
                    if (posI > 0 and ((posC > posI and posCC > posC+1 and expresion[posCC+1] != " ") or (posF > posI and posFF > posF+1 and expresion[posFF+1] != " "))):
                        if (expSub == acumExp):
                            if (posFF != -1):
                                file.write("\t\t<span class=\"libreria\">" + acumExp[:posF+1] + " " + acumExp[posF+1:] + "</span>\n")
                            else:
                                file.write("\t\t<span class=\"libreria\">" + acumExp + "</span>\n")
                            acumExp = ""
                            nullSpace = False
                            libreria[0] = False
                    # En caso contrario es un error
                    elif (posCC == -1 or posFF == -1 or posI == -1):
                        file.write("\t\t<span class=\"error\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                        libreria[0] = False
                # Verifica si es una palabra reservada - O(1)
                elif (acumExp != "" and not comentarioLargo[0] and isReservada(acumExp)):
                    # Verifica que despues de validar que la palabra reservada se encuentre en la expresión, haya un espacio en blanco o un salto de línea a continuación
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isDelimitador(expresion[j+1]) or isOperadorUnique(expresion[j+1],expresion,j) or expresion[j+1] == "\"" or expresion[j+1] == "\'"):
                        file.write("\t\t<span class=\"reservada\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                # Verifica si es un operador - O(n)
                elif (acumExp != "" and not comentarioLargo[0] and isOperador(acumExp,expresion,j) and not operadorOmit[0] and not nullSpace):
                    # Casos de excepción para marcar que son syntax error
                    if (len(acumExp) > 1 and (not isOperadorUnique(acumExp[0],expresion,expresion.find(acumExp)) or isComentario(expresion[j:]))):
                        file.write("\t\t<span class=\"error\">" + acumExp[:-1] + "</span>\n")
                        acumExp = acumExp[-1]
                        nullSpace = False
                        operadorOmit[0] = False
                    # Verifica una vez retirado la expresión erronea, hay un operador válido
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isOperadorUnique(expresion[j+1],expresion,j) or isIdentificador(expresion[j+1],expresion,j) or isLiteral(expresion[j+1],expresion,j,operadorOmit)):
                        if (acumExp.find("_") == -1):
                            file.write("\t\t<span class=\"operador\">" + acumExp + "</span>\n")
                            acumExp = ""
                        else:
                            file.write("\t\t<span class=\"operador\">" + acumExp[:acumExp.find("_")] + "</span>\n")
                            acumExp = acumExp[acumExp.find("_"):]
                        nullSpace = False
                        operadorOmit[0]
                # Verifica si es un delimitador - O(n)
                elif (acumExp != "" and not comentarioLargo[0] and isDelimitador(acumExp) and not nullSpace):
                    if (len(acumExp) > 1 and not isDelimitador(acumExp[0])):
                        file.write("\t\t<span class=\"error\">" + acumExp[:-1] + "</span>\n")
                        acumExp = acumExp[-1]
                        nullSpace = False
                    # Verifica que a continuación haya un caracter diferente válido para desplegarlo
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isOperadorUnique(expresion[j+1],expresion,j) or isIdentificador(expresion[j+1],expresion,j) or isLiteral(expresion[j+1],expresion,j,operadorOmit) or expresion[j+1] == "."):
                        if (acumExp.find("_") == -1):
                            file.write("\t\t<span class=\"delimitador\">" + acumExp + "</span>\n")
                            acumExp = ""
                        else:
                            file.write("\t\t<span class=\"delimitador\">" + acumExp[:acumExp.find("_")] + "</span>\n")
                            acumExp = acumExp[acumExp.find("_"):]
                        nullSpace = False
                # Verifica si es un identificador - O(1)
                elif (acumExp != "" and not comentarioLargo[0] and isIdentificador(acumExp, expresion, j)):
                    # Verifica que a continuación haya un caracter diferente válido para desplegarlo
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isOperadorUnique(expresion[j+1],expresion,j) or isDelimitador(expresion[j+1])):
                        file.write("\t\t<span class=\"identificador\">" + acumExp + "</span>\n")
                        acumExp = ""
                        nullSpace = False
                # Verifica si es una literal - O(n)
                elif (acumExp != "" and not comentarioLargo[0] and isLiteral(acumExp, expresion, j, operadorOmit)):
                    # Verifica que a continuación haya un caracter diferente válido para desplegarlo
                    if (j == len(expresion)-1 or expresion[j+1] == " " or expresion[j+1] == "\n" or isDelimitador(expresion[j+1]) or (isOperadorUnique(expresion[j+1],expresion,j) and expresion[j+1] != "-") or isComentario(acumExp[j+1:]) or (acumExp[0] == "\'" and acumExp[len(acumExp)-1] == "\'") or (acumExp[0] == "\"" and acumExp[len(acumExp)-1] == "\"")):
                        # Verifica que esten ambas comillas en la literal para poder habilitar la opción de leer espacios en blanco y otros caracteres
                        if ((acumExp[0] == "\'" and expresion[j+1:].find("\'") != -1) or (acumExp[0] == "\"" and expresion[j+1:].find("\"") != -1)):
                            nullSpace = True
                        # Verifica que sea una literal de tipo string o char
                        if ((acumExp[0] == "\'" and acumExp[len(acumExp)-1] == "\'" and len(acumExp) != 1) or (acumExp[0] == "\"" and acumExp[len(acumExp)-1] == "\"" and len(acumExp) != 1)):
                            file.write("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
                            acumExp = ""
                            operadorOmit[0] = False
                            nullSpace = False
                        # Verifica que sea una literal de tipo númerica
                        elif(((acumExp[0] != "\'") and (acumExp[len(acumExp)-1] != "\'")) and (acumExp[0] != "\"" and acumExp[len(acumExp)-1] != "\"")):
                            file.write("\t\t<span class=\"literal\">" + acumExp + "</span>\n")
                            acumExp = ""
                            operadorOmit[0] = False
                            nullSpace = False

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