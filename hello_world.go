// Comentario de una sola línea
/* Comentario
multilínea */

package main //Todoslos archivos Go deben pertenecer a un paquete

//Importar los paquetes necesarios
import "fmt" // Un paquete en la biblioteca estándar de Go

// Biblioteca de matemáticas con alias local m
/*
Python equivalencia: import math as m
*/

/*
Maneras de definir una variable:
Tipo ~ Python
Tipo Javascript
*/

// Tipo Javascript
var local string  // Las variales que inician con minuscula son privadas
var Global string // Las variales que inician con mayuscula son publicas

// Similar a Python

func main() { //El main es el punto de entrada asi como en C++

	saludo := "Hola Mundo!" //Tipado dinamico, similar a Python
	var otro_saludo = "Hola, otra vez"

	// Imprime
	fmt.Println(saludo, otro_saludo)
	fmt.Println("Hello World!")
}
