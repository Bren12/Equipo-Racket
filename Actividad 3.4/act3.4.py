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
nombre_archivo = os.path.join(folder_actual, "sintaxis.txt")

# Ejemplo: int a = 5 // a es igual a 5

# Definición de función principal
def main():
    # Abrir el archivo de expresiones
    with open (nombre_archivo, mode = 'r', encoding = "utf-8") as archivo:
        # Leer todo el archivo y dejarlo en una lista de strings
        lista_expresion = archivo.readlines()
        
    for num, expresion in enumerate(lista_expresion):
        isComentario(expresion) # Brenda
        isLibreria(expresion) # Diego
        isReservada(expresion) # Jose Angel
        isLiteral(expresion) # Brenda
        isOperador(expresion) # Brenda
        isDelimitador(expresion) # Jose Angel
        isIdentificador(expresion) # Diego

# Se llama a la función principal
main()