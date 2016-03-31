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