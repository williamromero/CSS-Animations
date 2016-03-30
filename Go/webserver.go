package main

import (
"fmt"
"log"		// Libreria para sacar mensajes 
"net/http"	// Libreria para manejar peticiones
)

func main(){ // función main
http.HandleFunc("/", handler)
log.Fatal(http.ListenAndServe("localhost:8000", nil)) // uso de la libreria log y hay problem
// tambien indica que nuestro localhost será esa.
}

func handler(w http.ResponseWriter, r *http.Request){ // Envía y recibe peticiones
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// func handler (w http.ResponseWriter, r *http.ReponseWriter)
// Es la función que manejará las peticiones de entrada y de salida.
