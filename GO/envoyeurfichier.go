package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//mm port, localhost parce que c interne
	c, _ := net.Dial("tcp", "localhost:8000")

	// partout la j'ai mis _ pour erreur pour l'instant pour tester mais faudra peut etre mettre les if machin

	file, _ := os.Open("mat1.txt")

	data, _ := io.Copy(c, file) //https://pkg.go.dev/io#Copy
	//data = nombre de bytes envoyé

	var _ = file.Close()          //erreur negligée
	fmt.Println("envoyé: ", data) // test voir ce qu'on a send

	c.Close() //erreur negligée

}
