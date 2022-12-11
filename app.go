package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println("Hola mundo")
	/*Variables y tipos*/
	//var nombre_variable tipo_dato
	/*var X, Y, Z int
	X = 23
	var cadena string
	var bandera bool
	var cadenas []string
	*/

	/*---------------------------------------------------*/
	//Version tipado dinamico de Go
	X := 23
	fmt.Println(X)

	nombre := "nombre"
	nombre = "nuevo nombre" //Se debe asignar con = cuando ya esta declarada la variable
	fmt.Println(nombre)

	/*---------CONVERSION DE TIPOS--------------------*/
	edad := 22
	edad2 := "22"
	edad_str := strconv.Itoa(edad)
	edad_int, _ := strconv.Atoi(edad2) // Con el guion bajo descartamos los valores de retorno de las funciones
	fmt.Println("Mi edad es " + edad_str)
	fmt.Println(edad_int + 10)

	/*--------------INPUT AND OUTPUT-----------------*/
	fmt.Println("Funcion de output")

	edad3 := 22
	fmt.Printf("Mi edad es %d\n", edad3)
	bandera2 := true
	fmt.Printf("El valor de la bandera es %t\n", bandera2)
	precio := 14.3
	fmt.Printf("El precio es %f\n", precio)
	fmt.Printf("Mi nombre es %s\n", "Aldo Ayala")

	var edad4 int
	fmt.Println("Coloca tu edad:")
	fmt.Scanf("%d\n", &edad4)
	fmt.Printf("Mi edad es %d\n", edad4)

	/*--------USO DE BUFIO-----------------*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese tu nombre: ")
	nombre2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre2)
	}

	/*------------CONDICIONALES-------------*/

	/*
	 == igual a
	 != diferente de
	 < menor que
	 > mayor que
	 <= menor o igual que
	 >= mayor o igual que
	 && AND
	 || OR
	*/
	if true {
		fmt.Println("Entro por el true")
	}

	x := 10
	y := 5
	if x > y {
		fmt.Printf("%d es mayor que %d \n", x, y)
	} else if x < y {
		fmt.Printf("%d es menor que %d \n", x, y)
	} else {
		fmt.Printf("%d es igual que %d \n", x, y)
	}

	/*------------CICLOS-------------*/
	for i := 0; i < 10; i++ {
		fmt.Println("Iteracion " + strconv.Itoa(i))
	}

	indice := 0

	for indice < 10 {
		fmt.Println("Iteracion2 " + strconv.Itoa(indice))
		indice++
	}

	/*CICLOS INFINITOS Y QUIEBRES*/
	indiceBreak := 0
	for {
		fmt.Println(indiceBreak)
		indiceBreak++
		if indiceBreak >= 10 {
			break
		}
	}

	/*------------ARREGLOS-----------------*/
	//var areglo [3]int -> Se inicializa el arreglo con los elementos en cero
	arreglo := [3]int{1, 2, 3}
	fmt.Println(arreglo)
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(i)
		fmt.Println(arreglo[i])
	}

	/*----MATRICES----*/
	var matriz [2][2]int
	matriz[0][1] = 20
	fmt.Println(matriz)
	fmt.Println(matriz[0][1])

	/*---------------SLICES------------------*/
	//var slice []int  El tamaño puede variar en tiempo de ejecucion
	slice := []int{1, 2, 3, 4, 5}

	if slice != nil {
		fmt.Println(slice)
	}

	//Puntero del arreglo
	//Longitud del arreglo al que apunta
	//Capacidad del arreglo

	//Obtener un slice desde un arreglo - Slicing
	arreglo2 := [3]int{1, 2, 3}
	slice2 := arreglo2[1:2] //Rango a copiar del arreglo. No es inclusivo en los extremos
	fmt.Println(arreglo2)
	fmt.Println(slice2)

	/*-----------MAKE AND APPEND-----------------*/
	slice3 := make([]int, 3, 5)
	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println(cap(slice3)) //Capacidad del slice

}
