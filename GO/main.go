package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func read(filename string) [][]int {
	file, err := os.Open(filename) //on ouvre le fichier qui contient la matrice !!! attention renvoie fichier ET erreur
	if err != nil {                //gestion de l'erreur pour voir si le fichier existe
		fmt.Println("erreur ! j'arrive pas a ouvrir la matrice")
		return nil
	}

	scanner := bufio.NewScanner(file)
	var matrice [][]int
	for scanner.Scan() {
		temp := strings.Fields(scanner.Text()) //  Fields = split python mais en moins bien ici tout ca pour une boucle classique qui parcourt une ligne
		var ligne []int
		for _, i := range temp { //
			//https://www.educative.io/answers/what-is-the-fields-function-in-go
			n, err := strconv.Atoi(i) // convert string en entier https://pkg.go.dev/strconvl
			if err != nil {
				return nil
			}

			ligne = append(ligne, n) //ajout nombre a ligne
		}

		matrice = append(matrice, ligne) //ajout ligne a matrice
	}

	// Return the matrix
	return matrice
}

var wg sync.WaitGroup // pourquoi tu as mis ce truc ici en particulier

func calc(one [][]int, two [][]int, result [][]int, i int, j int) {
	for k := 0; k < len(two[0]); k++ {
		result[i][j] += one[i][k] * two[k][j]
	}
	wg.Done()
}

func convert(res [][]int) []byte {

	var buffer bytes.Buffer
	//boucle parcours simple
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {

			binary.Write(&buffer, binary.LittleEndian, int32(j)) //https://pkg.go.dev/encoding/binary#Write chelou un peu
		}
	}

	return buffer.Bytes()
}

func tcp(c net.Conn) {
	var stockage = make([]byte, 9999) // gros data jsp quelle taille il faut
	data, erreur := c.Read(stockage)  // lis les données et les stock dans stockage
	if erreur != nil {
		fmt.Println("Arrive pas a lire data")
	}
	file, err := os.Create("mat1new.txt")
	if err != nil {
		fmt.Print("arrive pas a creer fichier")
	}

	file.Write(stockage[:data]) // ecrit que les données recues, -> s'arrete a data parce que c'est la longueur
	file.Close()                // erreurs on fera plus tard

	var outfile, errr = os.Create("result.txt")
	if errr != nil {
		return
	}

	// pris le ancien code de main et mis ici
	mat1 := read("mat1new.txt")
	mat2 := read("mat2.txt") //peut pas utiliser mat2 encore

	result := make([][]int, len(mat1))
	for i := range result {
		result[i] = make([]int, len(mat2[0]))
	}

	for i := 0; i < len(mat1); i++ { // calc the matrice numbers yes
		for j := 0; j < len(mat2[0]); j++ {
			wg.Add(1)
			go calc(mat1, mat2, result, i, j) // concurrence fait tout seul insh
			// tableaux modifiés globallement comme java
		}
	}

	wg.Wait()

	resultconverted := convert(result) // pas encore fait conversion  !

	outfile.Write(resultconverted)
	outfile.Close()

	outgoingFile, err := os.Open("result.txt")
	if err != nil {
		fmt.Println("aled j'arrive pas a ouvrir un fichier")
	}
	defer outgoingFile.Close()

	outfilescan := bufio.NewScanner(outfile)
	outfilescan.Split(bufio.ScanBytes)

	for outfilescan.Scan() { // retourne vrai tant que ya des trucs à lire donc s'arrete quand plus rien a envoyer niquel
		c.Write(outfilescan.Bytes())
	}

	c.Close()
}

func main() { //Scanner fichier de la doc golang

	lst, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("a l'aide j'arrive pas a ecouter tcp !")

	}

	for { // pour chaque demande je fais groutine
		c, err201 := lst.Accept()
		if err201 != nil {
			fmt.Println("a l'aide j'arrive pas a accepter tcp !")
			os.Exit(1)
		}
		go tcp(c) // GROUTINE
	}

}
