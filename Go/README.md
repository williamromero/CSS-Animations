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
	
COMANDOS: 
		- go build --> compila el programa y genera los ejecutables 
		- go run --> compila y ademas ejecuta el programa
		- 
		El compilador toma cada una de las lineas y si uno quiere imprimir otra cosa en el mismo archivo, solo lee las lineas 		que se agregaron posteriormente.

		- Crear código:
		
		package main
		import "fmt"
		func main () {
		        fmt.Println("Ejemplo-> go build ejemplos")
		}
	
		- Darle un go run file.go
		
		Resultado de Consola: "Ejemplo-> go build ejemplos"
		
		- go build file.go // Esto meterá en caché las lineas que hay de programación
		
		package main
		import "fmt"
		func main () {
		        fmt.Println("Ejemplo-> go build ejemplos")
			fmt.Println("Segunda lineas")
		}

		- Darle go run file.go  // ya solo leerá la última instrucción añadida y eso hará más rápida la respuesta.
		[Nota] Build: compila todo lo nuevo en el script (.go) (de hecho lo relee todo, y para no demorar crea una cache de lo builds anteriores)

	ESTRUCTURA DE ITERACIÓN

	package main
	import "fmt"
	func main() {
		alumnos := 1
		for alumnos <= 3 {
		fmt.Println(alumnos)
		alumnos = alumnos + 1
		}
		for calificaciones := 7; calificaciones <= 9; calificaciones++ { 
		fmt.Println(calificaciones)
		}
	}
	
	
WEBSERVER
	package main
	
	import (
	"fmt"
	"log"           // Libreria para sacar mensajes
	"net/http"      // Libreria para manejar peticiones
	)
	
	func main(){ // función main
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // uso de la libreria log y hay problema. tambien indica que nuestro localhost será esa.
	}
	
	func handler(w http.ResponseWriter, r *http.Request){ // Envía y recibe peticiones
	        fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
	
	// func handler (w http.ResponseWriter, r *http.ReponseWriter)
	// Es la función que manejará las peticiones de entrada y de salida.


ARREGLOS

package main
import "fmt"
func main(){
	var nombres[5] string 	// Declarar arreglos
	amigos:= [5]string{'Raquel', 'Luis', 'Isabel', 'Enrique', 'José '}
	nombres = amigos 
	for i, nombres:= range nombres {
		fmt.Println(nombres, &nombres[i])
	}
}

SLICES 

- Tipo de datos que proporcionan secuencias mejores que los arreglos

	// DECLARAR UN SLICE DE ENTEROS VACIO
	// CREAR UN LOOP QUE META 10 VALORES AL SLICE
	// ITERAR SOBRE EL SLICE Y MOSTRAR CADA VALOR
	// DECLARAR UN SLICE DE 5 STRING E INICIALIZAR
	// DESPLEGAR O IMPRIMIR EL DEGUNDO SLICE
	
	package main //Es la edntrada a cualquier programa
	import "fmt"
	func main(){
		// crear array enteros vacíos
		var numeros []int
		// meter 10 numieros al slice
		for i:= 0; i < 10; i++ {
			numeros = append(numeros, i*10)
		}
		// mostrar cada valor
		for _, numero := range numeros {
		fmt.Println(numero)
		}
		// Declarar un slice de strings
		frutas := []string{"manzana", "naranja", "pera", "sandía", "aguacate"}
		// Mostrar cada índice o posición y cada nombre
		
		for i, fruta:= range frutas {
			fmt.Printf("Index: %d Fruta %s\n", i, fruta)
		}
		// Tomar un slice de índice 1 y 2
		slice := frutas[1:3]
		//Mostrar el valor del nuevo slice
		for i, fruta:= range slice {
		fmt.Printf("Index %d Fruta %s\n", i, fruta)
		}
	}

//el _ es un identificador especial para ignorar la variable donde normalmente te pasan el índice de la iteración
//for _, val := range foo ignora el key por el _ y toma el valor val de foo, que puede ser un mapa


MAPS - RELACIÓN SINTÁCTICA EN GO (Como hashes)

// Declarar y hacer un mapa de valores enteros con un string como llave
// Llenar el mapa con 5 valores e iterar sobre el mapa para mostrar los pares llave/valor

	package main
	import "fmt"
	
	// Main es la entrada de nuestro programa
	func main() {
		// Declarar y hacer el mapa
		departamentos := make(map[string]int)
		// Inicializar datos en el mapa
		departamentos["Devs"] = 25
		departamentos["Marketing"] = 50
		departamentos["Ejecutivos"] = 4
		departamentos["Ventas"] = 60
		departamentos["Mantimiento"] = 8
	
		// Desplegar por medio de iteracion el valor de cada par llave/valor
		for key, value := range departamentos {
			fmt.Printf("Depts: %s Personas %d\n", key, value)
		}
	}

RANGE 
	
	package main
	import "fmt"
	func main(){
		numeros := []int{2,4,6}
		suma := 0
		for _, numero := range numeros {
			suma += numeros
		}
		fmt.Println("suma:" suma)
	
		for i, numero := range numeros {
			if numero == 3
			fmt.Println("index:", i)
		}
	}
	
	//algo := map[string]string{"a":auto, "b":bebé}


STRUCTS
- ESTRUCTURA DE GO QUE TE PERMITE GUARDAR PROPIEDADES DE UN TIPO DE DATOS
	// Ejercicios:
	// Declarar un struct para mantener info. de un usuario (nombre, dirección, 
	// edad)
	// Crea un valor y lo vamos a inicializar con valores
	// Mostrar cada campo

	package main
	import "fmt"
	// usuario representa un usuario en el sistema

	type usuario struct {
		nombre string
		dirección string
		edad init
	}

	// punto de entrada de nuestra app
	func main(){
		// Declrar la variable ususario y la iniciamos usando un struct
		vero := usuario {
			nombre: "Verónica",
			direccion = "Calle 12",
			edad = 38,
		}
	}

	// Mostrar los valores de cada campo

	fmt.Println("Nombre", vero.nombre)
	fmt.Println("Dirección", vero.direccion)
	fmt.Println("Age", vero.edad)

	// Declarar otro struct anano

	nicole := struct {
		nombre string
		direccion string
		edad init
	} {

	nombre: "William",
	direccion: "Calle 13",
	edad: "22" 
	}

	// Imprimir datos de Nicole

	fmt.Println("Nombre", nicole.nombre)
	fmt.Println("Dirección", nicole.direccion)
	fmt.Println("Edad", nicole.edad)


METODOS

// Declarar un struct que representa a un jugador de soccer (nombre, goles, 
// partidos)
// Declarar un método que calcule el promedio de goles de un jugador
// formula: partido/goles

// Declarar un slice que inicialice un slice con varios jugadores.
// Iterar sobre el slice mostrando los jugadores y su promedio de lanzamiento/goles


package main
import "fmt"
// jugador representa una persona en el juego
type jugador struct {
	nombre string
	goles int
	partidos int
}

// Promedio de Goles

	func (j *jugador) average()float64{
	return float64(j.partidos) / float64(j.goles)
	}

	func main(){
	jugadores := []jugador {
		{"carlos", 20, 60}
		{"fernando", 17, 30}
		{"alonso", 34, 60}
	}
}

INTERFACES

// Declarar un interface llamada speaker con un método llamado speak
// Declarar un struct llamado ingles que represente a una persona que hable inglés
// Declarar un struct llamado chino que represente a una persona que hable chino.
// Implementar la interface speaker para cada struct usando un valor y strings: "Hello World" y "a23452634"
