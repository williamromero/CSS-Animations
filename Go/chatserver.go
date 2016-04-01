// CHAT SERVER
// Trabajo pesado

package main
import (
"bufio"
"fmt"
"log"
"net"
)

// 1. Estamos haciendo el componente que emite los mensajes
// Los mensajes SALIENTES

type client chan <- string
var ( 
	entrantes = make(chan client)
	salientes = make(chan client)
	mensajes = make(chan string) 
)

// 2. Función broadcaster
func broadcaster(){
	clients := make(map[client]bool)	
}
// Todos los clientes conectados
// CLI Comand line interaction
for { 
	select {
		case msg := <- mensajes:
		// Emitir el mensaje entrante a todos los clientes que están
		// conectados por medio de channels
		for cli := range clients {
			cli <- msg
		}

		case cli := <- entrantes:
		clients[cli] = true

			case cli := <- salientes:
			delete(clients, cli)
			close(cli)
		}
	}
}
// 3. Manejador de datos
func handleConn(conn net.Conn){
	// mensajes salientes
	ch := make(chan string)
	go clientWriter(conn, ch)

	quien := conn.RemoteAddr().String()
	ch <- "Tu eres "+quien
	mensajes <- quien + "se ha conectado a la sala"
	entrantes <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		mensajes <- quien + ":"+ input.Text()
	}

	salientes <- ch
	mensajes <- quien + "se ha ido"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string){
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// main: entrada a nuestra aplicación
func main (){
	listener, err := net.Listen("tcp", "localhost:8005")
	if err !=nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
