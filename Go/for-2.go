package main
import "fmt"
func main(){
	var nombres[5] string 	// Declarar arreglos
	amigos:= [5]string{"Raquel", "Luis", "Isabel", "Enrique", "José"}
	nombres = amigos 
	for i, nombre:= range nombres {
		fmt.Println(nombre, &nombres[i])
	}
}
