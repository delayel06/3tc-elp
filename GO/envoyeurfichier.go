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

	file, _ := os.Open("mat1.txt")

	data, _ := io.Copy(c, file) //https://pkg.go.dev/io#Copy
	//data = nombre de bytes envoyé

	fi, _ := file.Stat()
	if data != fi.Size() {
		fmt.Println("Erreur ! Je n'ai pas envoyé le bon nombre de bytes.")
		return
	}

	defer file.Close()
	fmt.Println("(1) Bytes envoyées: ", data) // test voir ce qu'on a send

	file2, _ := os.Open("mat2.txt")

	data2, _ := io.Copy(c, file2) //https://pkg.go.dev/io#Copy
	//data = nombre de bytes envoyé

	fi2, _ := file2.Stat()
	if data2 != fi2.Size() {
		fmt.Println("Erreur ! Je n'ai pas envoyé le bon nombre de bytes.")
		return
	}

	defer file.Close()
	fmt.Println("(2) Bytes envoyées: ", data2) // test voir ce qu'on a send

	var stockage = make([]byte, 2048)
	newdata, erreur := c.Read(stockage) // lis les données et les stock dans stockage
	if erreur != nil {
		fmt.Println("Je n'arrive pas a lire les données. Les dimensions des matrices ne doivent pas être bonnes.")
	} else {
		fmt.Println("(client) Bytes reçues: ", newdata)
	}
	fmt.Println(string(stockage[:data]))

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Print("arrive pas a creer fichier")
	}

	file.Write(stockage[:data])
	defer file.Close()

	defer c.Close()
}
