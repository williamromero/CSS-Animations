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
			nombre : "Verónica",
			direccion : "Calle 12",
			edad : 38,
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

}