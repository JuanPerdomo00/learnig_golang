package main

import (
	"fmt"
)

func main() {
	// Numeros enteros
	// int = Depende del OS (32 o 64 bits)
	// int8 = 8bits = -128 a 127
	// int16 = 16bits = -2^15 a 2^15-1
	// int32 = 32bits = -2^31 a 2^31-1
	// int64 = 64bits = -2^63 a 2^63-1
	var entero int
	var entero8 int8
	var entero16 int16
	var entero32 int32
	var entero64 int64

	// Numeros enteros sin signo
	// uint = Depende del OS (32 o 64 bits)
	// uint8 = 8bits = 0 a 127
	// uint16 = 16bits = 0 a 2^15-1
	// uint32 = 32bits = 0 a 2^31-1
	// uint64 = 64bits = 0 a 2^63-1
	var uintero uint
	var uintero8 uint8
	var uintero16 uint16
	var uintero32 uint32
	var uintero64 uint64

	// Numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308
	var decimal32 float32
	var decimal64 float64

	// Textos y booleanos
	// string = ""
	// bool = true or false
	var texto string
	var booleano bool

	// Numeros complejos
	// Complex64 = Real e Imaginario float32
	// Complex128 = Real e Imaginario float64
	var complejo64 complex64
	var complejo128 complex128

	// Imprimir información sobre los tipos de datos
	fmt.Println("Tipos de datos y rangos:")
	fmt.Printf("int: %T\n", entero)
	fmt.Printf("int8: %T\n", entero8)
	fmt.Printf("int16: %T\n", entero16)
	fmt.Printf("int32: %T\n", entero32)
	fmt.Printf("int64: %T\n", entero64)
	fmt.Printf("uint: %T\n", uintero)
	fmt.Printf("uint8: %T\n", uintero8)
	fmt.Printf("uint16: %T\n", uintero16)
	fmt.Printf("uint32: %T\n", uintero32)
	fmt.Printf("uint64: %T\n", uintero64)
	fmt.Printf("float32: %T\n", decimal32)
	fmt.Printf("float64: %T\n", decimal64)
	fmt.Printf("string: %T\n", texto)
	fmt.Printf("bool: %T\n", booleano)
	fmt.Printf("complex64: %T\n", complejo64)
	fmt.Printf("complex128: %T\n", complejo128)
}


// %v: Muestra el valor de la variable en un formato predeterminado
// %T: Muestra el tipo de la variable
// %d: Muestra el valor de la variable como un número entero decimal
// %f: Muestra el valor de la variable como un número de punto flotante
// %s: Muestra el valor de la variable como una cadena de texto
// %t: Muestra el valor de la variable como un booleano (true o false)
// %b: Muestra el valor de la variable en formato binario
// %x: Muestra el valor de la variable en formato hexadecimal
// %o: Muestra el valor de la variable en formato octal
