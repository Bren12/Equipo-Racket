'''
Actividad 3.2 - Programando un DFA

Fecha: 22-04-2020

Equipo:
    - 
    -
    - Brenda Elena Saucedo González | A00829855
'''

# Definición de función principal
def lexerAritmetico(nombre_archivo):
    
    # Abrir el archivo de expresiones
    with open (nombre_archivo, mode = 'r', encoding = "utf-8") as archivo:
        # Leer todo el archivo y dejarlo en una lista de strings
        lista_expresion = archivo.readlines()
        
    # Lee cada enunciado del archivo
    for num, expresion in enumerate(lista_expresion):
        
        entero = ""
        real = ""
        resta = ""
        comentario = ""
        division = ""
        var = ""
        
        floatBool = False
        comentarioBool = False
        variableBool = False
        
        print(expresion)
        print ("Token\tTipo")
        
        # Lee cada caracter del enunciado
        for num2, token in enumerate(expresion):

            if comentarioBool == True:
                comentario = comentario + token

            elif token == "=" or token == "*" or token == "+" or token == " " or token == "^" or token == "(" or token == ")":
                if entero != "":
                    print(entero + "\tEntero")
                    entero = ""
                elif real != "":
                    print(real + "\tReal")
                    real = ""
                    floatBool = False
                elif resta != "":
                    print(resta + "\tResta")
                    resta = ""
                elif division != "":
                    print(division + "\tDivision")
                    division = ""
                elif var != "":
                    print(var + "\tVariable")
                    var = ""
                    variableBool = False
                
                if token == "=":
                    print(str(token) + "\tAsignación")
                elif token == "*":
                    print(str(token) + "\tMultiplicación")
                elif token == "+":
                    print(str(token) + "\tSuma")
                elif token == "^":
                    print(str(token) + "\tPotencia")
                elif token == "(":
                    print(str(token) + "\tParéntesis que abre")
                elif token == ")":
                    print(str(token) + "\tParéntesis que cierra")

            elif (token.isdigit() or token == "_") and variableBool == True:
                if division == "/":
                    print(var + "\tVariable")
                    var = ""
                    print(division + "\tDivision")
                    division = ""
                var = var + token
                
            elif token.isdigit():
                
                if floatBool == False:
                    if division == "/":
                        print(entero + "\tEntero")
                        entero = ""
                        print(division + "\tDivision")
                        division = ""
                    entero = resta + entero + str(token)
                    resta = ""
                else:
                    if division == "/":
                        print(real + "\tReal")
                        real = ""
                        print(division + "\tDivision")
                        division = ""
                    real = entero + real + str(token)
                    floatBool = True
                    resta = ""
                
            elif (token == "." or (token == "E" and floatBool == True) or (token == "-" and floatBool == True) ):
                real = entero + real + str(token)
                floatBool = True
                entero = ""

            elif token.isalpha():
                if division == "/":
                    print(var + "\tVariable")
                    var = ""
                    print(division + "\tDivision")
                    division = ""
                var = var + token
                variableBool = True

            elif token == "-":
                resta = "-"
            
            elif token == "/":
                division = division + "/"

                if division == "//":
                    comentario = division
                    division = ""
                    comentarioBool = True
            
            
        if entero != "":
            print(entero + "\tEntero")
        elif real != "":
            print(real + "\tReal")

        if comentario != "":
            print(comentario + "\tComentario")

        if var != "":
            print(var + "\tVariable")
            
        print("\n")



# Se llama a la función principal
lexerAritmetico("expresiones.txt")

'''
Falta:
- La "E" de los flotantes puede ser minuscula o mayuscula, seguida de un entero positivo o negativo
- Diseño del autómata (Herramienta computacional para dibujarlo)
- Documentación (Manual del usuario, indicando cómo correr su programa y qué se obtiene de salida)
- Documentación (El autómata que resuelve el problema)
'''