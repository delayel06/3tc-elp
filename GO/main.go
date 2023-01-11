package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		text := scanner.Text()

		temp := strings.Fields(text) //  Fields = split python mais en moins bien ici tout ca pour une boucle classique qui parcourt une ligne
		print(temp)
		var ligne []int
		for i := 0; i < len(temp); i++ { //
			//https://www.educative.io/answers/what-is-the-fields-function-in-go
			n, err := strconv.Atoi(temp[i]) // convert string en entier https://pkg.go.dev/strconvl
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

func calc(one [][]int, two [][]int, result [][]int, i int, j int) {

	for k := 0; k < len(two[0]); k++ {
		result[i][j] += one[i][k] * two[k][j]
	}

}

func main() { //Scanner fichier de la doc golang

	mat1 := read("mat1.txt")
	mat2 := read("mat2.txt") //peut pas utiliser mat2 encore

	var result [][]int
	for j := 0; j < len(mat1); j++ { // mat result taille mat 1 colonnes x mat 2 ligne

		var ligner []int
		for k := 0; k < len(mat2[0]); k++ {
			ligner = append(ligner) // on rempli de 0 pour avoir une matrice résultat pseudo vide
		}
		result = append(result, ligner)
	} // init result

	for i := 0; i < len(mat1); i++ { // calc the matrice numbers yes
		for j := 0; j < len(mat2[0]); j++ {
			// ** go calc(mat1, mat2, result, i, j) // concurrence fait tout seul insh
			// tableaux modifiés globallement comme java?
		}
	}

	for i := 0; i < len(mat2); i++ {

		for j := 0; j < len(mat2[0]); j++ {
			fmt.Printf("%d ", mat2[i][j]) //d format = int
		}

		fmt.Println()
	}

}
