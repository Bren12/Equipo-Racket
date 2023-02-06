# Instrucciones

Hacer un programa que reciba como entrada un archivo de texto que contenga expresiones aritméticas y comentarios, 
y nos regrese una tabla con cada uno de sus tokens encontrados, en el orden en que fueron encontrados e indicando de qué tipo son.



***



## Tipos de tokens

Las expresiones aritméticas sólo podrán contener los siguientes tipos de tokens:

<ul>
<li> Enteros </li>
<li> Flotantes (Reales) </li>
<li> Operadores:
  <ul>
    <li> Asignación </li>
    <li> Suma </li>
    <li> Resta </li>
    <li> Multiplicación </li>
    <li> División </li>
    <li> Potencia </li>
  </ul>
</li>
<li> Identificadores: 
  <ul>
    <li> Variables </li>
  </ul>
</li>
<li> Símbolos especiales: 
  <ul>
    <li> ( </li>
    <li> ) </li>
  </ul>
</li>
<li> Comentarios: 
  <ul>
    <li> // seguido de caracteres hasta que se acabe el renglón </li>
  </ul>
</li>
</ul>



***



## Función principal

El programa podrá estar formado con las funciones que requiera, pero la función principal tendrá la siguiente firma:

**def** lexerAritmetico(nombre_archivo)

donde **nombre_archivo** es el nombre del archivo que contiene las expresiones a ser analizadas 
(el nombre debe incluir la extensión, por ejemplo, **expresiones.txt**).



***



## Entrada

- Un archivo tipo texto que contenga una o más expresiones aritméticas, una por renglón.
- Los tokens no necesariamente deben estar separados por un blanco, o pueden tener separación de más de un blanco

Por ejemplo:

b=7

a = 32.4 *(-8.6 - b)/       6.1E-8

d = a ^ b // Esto es un comentario



***



## Salida

Debe entregar la siguiente salida:

| Token                    | Tipo                 |
| ------------------------ |:--------------------:|
| b                        | Variable             |
| =                        | Asignación           |
| 7                        | Entero               |
| a                        | Variable             |
| =                        | Asignación           |
| 32.4                     | Real                 |
| *                        | Multiplicación       |
| (                        | Paréntesis que abre  |
| -8.6                     | Real                 |
| -                        | Resta                |
| b                        | Variable             |
| )                        | Paréntesis que cierra|
| /                        | División             |
| 6.1E-8                   | Real                 |
| d                        | Variable             |
| =                        | Asignación           |
| ^                        | Potencia             |
| b                        | Variable             |
| // Esto es un comentario | Comentario           |



***



## Reglas de formación de algunos tokens

<ul>
<li> Variables:
  <ul>
    <li> Deben empezar con una letra (mayúscula o minúscula). </li>
    <li> Sólo están formadas por letras, números y underscore (‘_’). </li>
  </ul>
</li>
<li> Números reales (de punto flotante):
  <ul>
    <li> Pueden ser positivos o negativos </li>
    <li> Pueden o no tener parte decimal pero deben contener un punto (e.g. 10. o 10.0) </li>
    <li> Pueden usar notación exponencial con la letra E, mayúscula o minúscula, pero después de la letra E sólo puede ir un entero positivo o negativo (e.g. 2.3E3, 6.345e-5, -0.001E-3, .467E9). </li>
  </ul>
</li>
<li> Comentarios: 
  <ul>
    <li> ( </li>
    <li> ) </li>
  </ul>
</li>
<li> Comentarios: 
  <ul>
    <li> Inician con // y todo lo que sigue hasta que termina el renglón es un comentario </li>
  </ul>
</li>
</ul>



***



## Algoritmo

- El reconocimiento de tokens se debe hacer por medio de la tabla de transición de un Autómata Finito Determinístico.
- El diseño del autómata debe ser parte fundamental de la documentación (utilice alguna herramienta computacional para dibujarlo, no lo haga a mano).



***



## Documentación:

1. Manual del usuario, indicando cómo correr su programa y qué se obtiene de salida.
2. El autómata que resuelve el problema (como un anexo del punto 1).
