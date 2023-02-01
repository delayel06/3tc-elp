package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	//On recupère les matrices à calculer
	mat1 := os.Args[1] + ".txt"
	mat2 := os.Args[2] + ".txt"
	result := os.Args[3] + ".txt"

	fmt.Println()

	//mm port, localhost parce que c interne
	c, _ := net.Dial("tcp", "localhost:8000")

	file, _ := os.Open(mat1)

	data, _ := io.Copy(c, file) //https://pkg.go.dev/io#Copy
	//data = nombre de bytes envoyé

	fi, _ := file.Stat()
	if data != fi.Size() {
		fmt.Println("Erreur ! Je n'ai pas envoyé le bon nombre de bytes.")
		return
	}

	defer file.Close()
	fmt.Println("(1) Bytes envoyées: ", data) // test voir ce qu'on a send

	file2, _ := os.Open(mat2)

	data2, _ := io.Copy(c, file2) //https://pkg.go.dev/io#Copy
	//data = nombre de bytes envoyé

	fi2, _ := file2.Stat()
	if data2 != fi2.Size() {
		fmt.Println("Erreur ! Je n'ai pas envoyé le bon nombre de bytes.")
		return
	}

	defer file2.Close()
	fmt.Println("(2) Bytes envoyées: ", data2) // test voir ce qu'on a send

	receivedData := make([]byte, 2048)

	number, err := c.Read(receivedData)

	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	file3, err := os.Create(result)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file3.Close()

	_, err = file3.Write(receivedData[:number])
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file")

	defer c.Close()
}
