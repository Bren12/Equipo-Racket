'''
Actividad 3.2 - Programando un DFA

Fecha: 22-04-2020

Equipo:
    - Diego Alberto Baños Lopez | A01275100
    - José Ángel Rentería Campos | A00832436
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
        adjunto = False
        floatE = False
        
        expresion = expresion.replace(" ","")
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
                    adjunto = False
                elif real != "":
                    print(real + "\tReal")
                    real = ""
                    floatBool = False
                    adjunto = False
                elif division != "":
                    print(division + "\tDivision")
                    division = ""
                    adjunto = True
                elif var != "":
                    print(var + "\tVariable")
                    var = ""
                    variableBool = False
                    adjunto = False
                
                if token == "=":
                    print(str(token) + "\tAsignación")
                    adjunto = True
                elif token == "*":
                    print(str(token) + "\tMultiplicación")
                    adjunto = False
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

            elif (token.isdigit() or token == "_") and variableBool == True:
                if division == "/":
                    print(var + "\tVariable")
                    var = ""
                    print(division + "\tDivision")
                    division = ""
                var = var + token
                adjunto = False
                
            elif token.isdigit():
                if floatBool == False:
                    if division == "/":
                        if (entero != ""):
                            print(str(entero) + '\tEntero')
                            entero = ""
                        print(division + "\tDivisión")
                        division = ""
                    if adjunto == True:
                        entero = resta + str(token)
                        resta = ""
                    else:
                        entero = entero + str(token)
                else:
                    if division == "/":
                        if (real != ""):
                            print(str(real) + '\tReal')
                            real = ""
                        print(division + "\tDivision")
                        division = ""
                    if adjunto == True:
                        real = resta + entero + real + str(token)
                        resta = ""
                    else:
                        real = entero + real + str(token)
                adjunto = False
                
            elif token == "." or (floatBool == True and (token == "E" or token == "e" or (token == "-" and floatE == True))):
                if token == "E" or token == "e":
                    floatE = True
                real = entero + real + str(token)
                entero = ""
                floatBool = True
                adjunto = True

            elif token.isalpha():
                if division == "/":
                    print(var + "\tVariable")
                    var = ""
                    print(division + "\tDivision")
                    division = ""
                var = var + token
                variableBool = True
                adjunto = False

            elif token == "-":
                resta = "-"
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
                    
            
            elif token == "/":
                division = division + "/"
                if division == "//":
                    comentario = division
                    division = ""
                    comentarioBool = True
                adjunto = False
                    
            if resta != "" and adjunto == False:
                print(resta + "\tResta")
                resta = ""
            
            
        if entero != "":
            print(entero + "\tEntero")
        elif real != "":
            print(real + "\tReal")
        elif var != "":
            print(var + "\tVariable")
            
        if comentario != "":
            print(comentario + "\tComentario")
            
        print("\n")



# Se llama a la función principal
lexerAritmetico("./Actividad 3.2/expresiones.txt")

'''
Falta:
- Diseño del autómata (Herramienta computacional para dibujarlo)
- Documentación (Manual del usuario, indicando cómo correr su programa y qué se obtiene de salida)
- Documentación (El autómata que resuelve el problema)
'''