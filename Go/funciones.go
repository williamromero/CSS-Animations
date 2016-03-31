// FUNCIONES

// Declaramos un struct para guardar información de un usuario.
// Declarar la función que crea el valor y regresa apuntadores de este tipo y
// un error como valor.
// Llamar esta función desde main y regresar el valor

// Hacer una segunda llamada a la función, pero ignorando el valor y sólo
// probando el error como valor.

package main
import "fmt"

// el usuario representa un usuario el sistema

type usuario struct {
	nombre string
	email  string
}

// nuevoUsuario crea y regresa apuntadores de valore de tipo usuario

func nuevoUsuario() (*usuario, error){
	return &usuario{"Verónica", "veronica@mailficticio.com"}, nil
}

// main es la entrada de todos nuestros programas

func main(){
	// Crear valores del tipo usuario
	u, err := nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Imprimimos el valor
		fmt.Println(*u)
	// 2do llamado a la función y que solo coque el error en el regreso ()
	_, err  = nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}
}