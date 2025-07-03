package main

import "fmt"

// funcion variatica solo numeros
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s la suma es: %d \n", name, total)
	return total
}

// funcion variatica con alfanumerico
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v \n", dato, dato)
	}
}

// vamos a crear una funcion para probar la recursividad de una funcion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(suma("juan", 12, 45, 78, 56))
	fmt.Println("###############################")
	fmt.Println(suma("daniel", 10, 20, 30, 40, 50))
	fmt.Println("###############################")
	imprimirDatos("Hola", 28, true, 3.14)
	fmt.Println("###############################")
	fmt.Println(factorial(3))
	fmt.Println(factorial(5))
	fmt.Println("###############################")

	// vamos a crear una funcion anonima
	func() {
		fmt.Println("hola soy una funcion anonima...!!!")
	}()
	fmt.Println("###############################")
	//funcion anonima con parametros y asignada a una variable
	saludo := func(name string) {
		fmt.Printf("hola soy una funcion anonima... con parametro %s!!!\n", name)
	}
	saludo("name:juandaniel")
	fmt.Println("###############################")
	saludar("daniel", saludo)
	fmt.Println("###############################")

	// vamos a pasar una funcion dentro de otra funcion
	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 10)
	fmt.Println(r1, r2)
	fmt.Println("###############################")

	// segundo ejemplo
	r := double(addOne, 3)
	fmt.Println("resultado: ", r)
	fmt.Println("###############################")

	//implementacion de closures
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("###############################")

}

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func saludar(name string, f func(string)) {
	f(name)
}

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}
