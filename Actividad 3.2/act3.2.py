'''
Actividad 3.2 - Programando un DFA

Fecha: 23-04-2020

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
nombre_archivo = os.path.join(folder_actual, "expresiones.txt")


# Definición de función principal
def lexerAritmetico(nombre_archivo):
    
    # Abrir el archivo de expresiones
    with open (nombre_archivo, mode = 'r', encoding = "utf-8") as archivo:
        # Leer todo el archivo y dejarlo en una lista de strings
        lista_expresion = archivo.readlines()
        
    #Imprime la cabeza de la tabla
    print ("Token\tTipo")
        
    # Lee cada enunciado del archivo
    for num, expresion in enumerate(lista_expresion):
        
        # Variable que almacena cada token seguido que sea entero
        entero = ""
        # Variable que almacena cada token seguido que sea real
        real = ""
        # Variable que almacena el simbolo de resta
        resta = ""
        # Variable que almacena cada token seguido que sea comentario
        comentario = ""
        # Variable que almacena el símbolo de división
        division = ""
        # Variable que almacena cada token seguido que sea variable
        var = ""
        
        # Variable booleana que nos ayuda a identificar si un entero paso a ser de tipo flotante
        floatBool = False
        # Variable booleana que nos ayuda a identificar si ya se empezo a leer comentarios
        comentarioBool = False
        # Variable booleana que nos ayuda a identificar si se estan leyendo variables
        variableBool = False
        # Variable booleana que nos ayuda a identificar si la resta esta despues de un "(", "/", "=", "^", "E" (float) o "e" (float).
        adjunto = False
        # Variable booleana que nos ayuda a identificar si la variable Real ya tiene una "E" o "e" incluida para validar que pueda seguir un "-"
        floatE = False
        
        # Lee cada caracter del enunciado
        for num2, token in enumerate(expresion):
            
            # Verifica que si ya se esta recibiendo token's que pertenecen a comentario
            if comentarioBool == True and token != "\n":
                comentario = comentario + token
                
            # Verifica si el token leído es un operador, "(" o ")", e imprime lo que se tenga almacenado en las variables entero, real o var,
            # para despues desplegar en pantalla dichos operadores y caracteres especiales.
            elif token == "=" or token == "*" or token == "+" or token == " " or token == "^" or token == "(" or token == ")" or (token == "-" and adjunto == False) or (comentarioBool == False and token == "/" and expresion[num2:num2+2] != "//" and expresion[num2-1:num2+1] != "//"):
                
                # Variables de tipo alfabeticas o numericas
                if entero != "":
                    print(entero + "\tEntero")
                    entero = ""
                    adjunto = False
                elif real != "":
                    print(real + "\tReal")
                    real = ""
                    floatBool = False
                    adjunto = False
                elif var != "":
                    print(var + "\tVariable")
                    var = ""
                    variableBool = False
                    adjunto = False
                
                # Operadores y caracteres especiales
                if token == "=":
                    print(str(token) + "\tAsignación")
                    adjunto = True
                elif token == "*":
                    print(str(token) + "\tMultiplicación")
                    adjunto = True
                elif token == "+":
                    print(str(token) + "\tSuma")
                    adjunto = False
                elif token == "^":
                    print(str(token) + "\tPotencia")
                    adjunto = True
                elif token == "(":
                    print(str(token) + "\tParéntesis que abre")
                    adjunto = True
                elif token == ")":
                    print(str(token) + "\tParéntesis que cierra")
                    adjunto = False
                elif token == "-":
                    print(str(token) + "\tResta")
                    adjunto = False
                elif token == "/":
                    print(str(token) + "\tDivision")
                    adjunto = True
                    
            # Verifica si ya se leyo anteriormente una letra, para poder comenzar a aceptar digitos y "_" que se reconocen tambien como variables
            elif (token.isdigit() or token == "_") and variableBool == True:
                var = var + token
                adjunto = False
                
            # Verifica si se esta leyendo un token de tipo númerico
            elif token.isdigit():
                # Si ya se estaba confirmado que se estaban leyendo enteros, sigue almacenandolos en dicha variable
                if floatBool == False:
                    # Si es valido recibir el signo "-", lo concatenamos a entero
                    if adjunto == True:
                        entero = resta + str(token)
                        resta = ""
                    # Si no, solo concatenamos el siguiente token recibido
                    else:
                        entero = entero + str(token)
                # Si ya se estaba confirmado que se estaban leyendo reales, sigue almacenandolos en dicha variable
                else:
                    if floatE == True:
                        floatE = False
                    # Si es valido recibir el signo "-", lo concatenamos a real
                    if adjunto == True:
                        real = resta + entero + real + str(token)
                        resta = ""
                    # Si no, solo concatenamos el siguiente token recibido
                    else:
                        real = entero + real + str(token)
                adjunto = False
                
            # Verifica si se esta leyendo un token que pertenece a los reales
            elif token == "." or (floatBool == True and (token == "E" or token == "e" or (token == "-" and floatE == True))):
                # Verificación para saber si el signo "-" es válido para la variable real o no (el "-" sigue de una "E" o "e")
                if token == "E" or token == "e":
                    floatE = True
                    adjunto = True
                real = entero + real + str(token)
                entero = ""
                floatBool = True
                
            # Verifica si se esta leyendo un token de tipo variable
            elif token.isalpha():
                var = var + token
                variableBool = True
                adjunto = False
                
            # Verifica si el token es un signo de resta y válido para adjuntarlos a las variables de entero o real
            elif token == "-":
                resta = token
            
            # Verifica si el token es un signo de división, el cuál sirve más que nada para activar la variable comentario
            elif token == "/":
                division = division + "/"
                adjunto = True
                if division == "//":
                    comentario = division
                    division = ""
                    comentarioBool = True
                    adjunto = False
            
        # En caso de que haya terminado de leer una expresión, pero no hayan quedado vacías ciertas variables 
        if entero != "":
            print(entero + "\tEntero")
        elif real != "":
            print(real + "\tReal")
        elif var != "":
            print(var + "\tVariable")
            
        # Despliega el comentario en caso de haber
        if comentario != "":
            print(comentario + "\tComentario")



# Se llama a la función principal
lexerAritmetico(nombre_archivo)