package main

import (
	"bufio"
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

func listentcp(c net.Conn) {
	var stockage = make([]byte, 9999) // gros data jsp quelle taille il faut
	data, erreur := c.Read(stockage)  // lis les données et les stock dans stockage
	if erreur != nil {
		fmt.Println("Arrive pas a lire data")
	}
	file, err := os.Create("mat1.txt")
	if err != nil {
		fmt.Print("arrive pas a creer fichier")
	}

	file.Write(stockage[:data]) // ecrit que les données recues, -> s'arrete a data parce que c'est la longueur
	file.Close()                // erreurs on fera plus tard

}

func main() { //Scanner fichier de la doc golang

	mat1 := read("mat1.txt")
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

	// Wait for all the goroutines to finish
	wg.Wait()

	for i := 0; i < len(result); i++ {

		for j := 0; j < len(result[0]); j++ {
			fmt.Printf("%d ", result[i][j]) //d format = int
		}

		fmt.Println()
	}

}
