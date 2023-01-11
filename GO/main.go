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
		fmt.Print(text)              // error check, a enlever
		temp := strings.Fields(text) //  Fields = split python mais en moins bien ici tout ca pour une boucle classique qui parcourt une ligne
		fmt.Print(temp)              // error check, a enlever
		var ligne []int
		for i := 0; i < 3; i++ { //
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

func main() { //Scanner fichier de la doc golang

	mat1 := read("mat1.txt")
	mat2 := read("mat2.txt")         //peut pas utiliser mat2 encore
	for i := 0; i < len(mat1); i++ { //j'arrive pas a afficher?
		fmt.Println("%d ", mat1[i])
		fmt.Print(",")
	}

	var result [][]int
	for j := 0; j < len(mat1); j++ {

		var ligner []int
		for k := 0; k < len(mat2[0]); k++ {
			ligner = append(ligner, 0) // on rempli de 0 pour avoir une matrice résultat pseudo vide
		}
		result = append(result, ligner)
	}

}
