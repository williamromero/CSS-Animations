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

// puedes utilizar "texto"+variable+"texto", como en Java