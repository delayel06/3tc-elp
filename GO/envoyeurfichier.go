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

	var stockage = make([]byte, 2048)   // gros data jsp quelle taille il faut
	newdata, erreur := c.Read(stockage) // lis les données et les stock dans stockage
	if erreur != nil {
		fmt.Println("Arrive pas a lire data cli")
	} else {
		fmt.Println("cool raoul jai recu: ", newdata)
	}

	defer c.Close() //erreur negligée// erreurs on fera plus tard
}
