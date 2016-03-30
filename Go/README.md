CURSO GO:

	Para instalar Go:
		- brew install go

	Correr primer programa:

		- Crear archivo example.go
		- Copiar el siguiente código:

			package main
			import "fmt"
			func main() {
				fmt.Println("HI")
			}

		- Go run example.go

	Teoría:

		Es un lenguaje compilado, creado en 2007. Anunciado en 2009.
		Nació a partir de la necesidad de tener sistemas grandes eficientes, pero simples.
		Aprovecha las arquitecturas modernas (cores, CPUs)

		Lenguaje creado por Robert Griesemer, "Rob Pike" y Hen Thompson.

	¿Es Go el nuevo C?

		- Como C, es una herramienta para que los programadores hagan grandes cosas, con muy pocos recursos.
		- Pero es mucho más que eso
		- Muy poderoso como C, pero con sintaxis legible, como Python
		- Open Source. Portable.
		- Go está escrito en Go - LOL - Otros lenguajes, no.
		- Go es capaz de desarrollar todo tipo de programas, sistemas operativos y apps web.
		
	Simplicidad

		- Muy legible y compacto == menos código.
		- Infraestructura más eficiente.
		- Sistema de tipos más simple que el de la mayoría de otros lenguajes tipados.
		- Baterias incluidas: Biblioteca Standard(STD Library) robusta: más de 100 paquetes para tareas comunes.

	¿Para qué sirve Go?

		- Lenguajes de sistemas != Lenguajes de Aplicaciones
		- Porque es un lenguaje Turing completo (que puede hacer cualquier cosa).

		Sistemas

			- Software diseñado para controlar hardware (de computadores, teléfonos, servidores, etc).

			- El sistema sirve como base para que corran las aplicaciones.

			- Ejemplos: 

				OS X ----> Drivers -----> Compiladores

			- Usuario: la computadora o desarrollador.


		Aplicaciones

			- Software diseñado con base en tareas o actividades coordinadas para el beneficio del usiario.
			- Conocidas también como apps (y antes también applets).
			- Ejemplos: Word, Chrome, Whatsapp, etc.
			- Usuario: Usuarios finales / 

			- Es un lenguaje, por lo tanto, puede tener mas de 1 framework e IDE
			- Se puede crear aplicaciones, aplicaciones web y el backend de ellos y que funcion con lo mismo.

		FRAMEWORKS

			- Beego, Gorilla, Revel, Martini, Hugo

		Contras:
			- Aún son nuevos. El desarrollo es más lento que en otros frameworks establecidos como Django o Rails. Aún faltan features.

		Pros: Constante desarrollo, todo tu proyecto hecho en Go.

    Paradigmas:
    
      - Funciona con varios paradigmas (POO)

    Manejo de Datos:
    
      - Utiliza y aprovecha todos los núcleos.
      
ESTRUCTURA DE UN PROGRAMA EN GO:

	package main			   
	import "fmt"
	func main (){
		fmt.Println("Hello world")
    	}
    	
    	"fmt" -> format
    	
    	// package main -> Es un paquete de archivos. Es un package que hace algo entre otros paquetes que podemos crear. el package main 			   es para el paquete donde esta el fmt y otras librerias. package main simplemente es el nombre del paquete, la 				   funcion principal es func main(){}. El package es para evitar colisiones con otras funciones de otros 				   paquetes. "Es el paquete central o general del lenguaje".
    	// import "fmt" -> Format es parte de la librería standard que permite hacer impresiones en pantalla y otras cosas.
    	// func main()  -> El paquete main define un paquete de programa por si solo.  la función main dentro del package main es el punto de entrada para un ejecutable
    	
    	- Si se importan paquetes que no vamos a utilizar, el asunto crashea.
    	
    		package main
		import "fmt"
		import "net"
		func main (){
			fmt.Println("Hello world")
		}
	Respuestra en consola: 
		# command-line-arguments
		./test.go:3: imported and not used: "net"
		
	- Los imports van luego del package main

	FUNCIONES:
		nombre_función(parámetros) 
		{ 
		   código
		}
	
	TIPOS DE DATOS:
		- func
		- var
		- const
		- type
	
		Cadenas, Enteros, Flotantes, Booleanos, Numeros complejos.
	
			- Go no hace conversión de variables númericas.
			- 
	func main(){
		fmt.Println("go"+"lang") 	// implementación de cadenas
		fmt.Println("5+5 =", 5+5)	// sumar/restar enteros
		fmt.Println("7.5/3.2=", 5.5/3.2) //manejo de flotantes
		
		fmt.Println("resultado=",)
		fmt.Println(true && false)  // true and false
		fmt.Println(true || false ) // true or false
		fmt.Println(!true)
	}		

	VARIABLES
		Fuertemente Tipado: Que tienes que declarar el tipo de dato de la variable y no puede ser cambiado en tiempo de 				ejecución.
		
	func main(){
		var a string = "inicial" //declarando la variable inicial
		fmt.Println(a)
	
		var b, c int = 4,5		 //declaración de multiples variables
		fmt.Println(b,c)
	
		var d bool = true
		fmt.Println(d)
	
		var e int  				// declaración de variable e - tipo entera
		fmt.Println(e)
	
		var f = 10
		fmt.Println(f)			// Otra forma de declarar variables
	
		g:= "short"				// Declaración de una cadena
		fmt.Println(g)
	
	}
	
	Otro ejemplo: Manejo de Constantes
	
	package main
	import "fmt"
	import "math"
	
	func main(){
		var a string = "inicial" //declarando la variable inicial
		fmt.Println(a)
	
		var b, c int = 4,5		 //declaración de multiples variables
		fmt.Println(b,c)
	
		var d bool = true
		fmt.Println(d)
	
		var e int  				// declaración de variable e - tipo entera
		fmt.Println(e)
	
		var f = 10
		fmt.Println(f)			// Otra forma de declarar variables
	
		g:= "short"				// Declaración de una cadena
		fmt.Println(g)
	
		const s string = "constante"
		fmt.Println(s)
	
		const n = 500
		const k = 3e20 / n
		fmt.Println(k)
		fmt.Println(int64(k))
		fmt.Println(math.Sin(n))
	
	}
	
	
