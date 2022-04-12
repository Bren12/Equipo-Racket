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
        print(expresion)
        print ("Token\tTipo")
        
        # Lee cada caracter del enunciado
        for num2, token in enumerate(expresion):
            if token.isalpha():
                print(str(token) + "\tVariable")
            elif token == "=":
                print(str(token) + "\tAsignación")
            elif token.isdigit():
                print(str(token) + "\tEntero")
            
            
        print("\n")
    
    
# Definición de la función main
def main():
    # Se llama a la función principal
    lexerAritmetico("expresiones.txt")
    
    
    
# Llamar a la funcion main
main()
