package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func strtointsplice(s string) [][]int {
	rows := strings.Split(s, "\n")
	var splice [][]int
	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], " ")
		var row []int // creer un truc a ajouter apres
		for j := 0; j < len(cols); j++ {

			num, _ := strconv.Atoi(cols[j])

			row = append(row, num) // doit append parce que on a pas initializé de taille -> dynamique
			//on ajout a chaque row le num puis chaque row jusqu'a la fin
		}
		splice = append(splice, row)
	}

	return splice
}
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

var wg sync.WaitGroup

func calc(one [][]int, two [][]int, result [][]int, i int, j int) {
	for k := 0; k < len(two[0]); k++ {
		result[i][j] += one[i][k] * two[k][j]
	}
	wg.Done()
}

func process(mat1 [][]int, mat2 [][]int) [][]int {

	result := make([][]int, len(mat1))
	for i := range result {
		result[i] = make([]int, len(mat2[0]))
	}

	for i := 0; i < len(mat1); i++ { // calc the matrice numbers yes
		for j := 0; j < len(mat2[0]); j++ {
			wg.Add(1)
			go calc(mat1, mat2, result, i, j) // concurrence fait tout seul insh
			fmt.Print("fais calc\n")          // test
			// tableaux modifiés globallement comme java
		}
	}

	return result

}

func tcp(c net.Conn) {
	var stockage = make([]byte, 2048) // gros data jsp quelle taille il faut
	data, erreur := c.Read(stockage)  // lis les données et les stock dans stockage

	if erreur != nil {
		fmt.Println("Arrive pas a lire data")
	} else {
		fmt.Println("(serv) cool raoul jai recu: ", data)
	}

	matreceived := string(stockage[:data])

	result := process(strtointsplice(matreceived), read("mat2.txt"))
	fmt.Println(strtointsplice(matreceived), "et autre : ", read("mat2.txt"))
	fmt.Println("result:", result)
	wg.Wait()

	var outfile, errr = os.Create("result.txt")
	if errr != nil {
		return
	}

	w := bufio.NewWriter(outfile)

	for j := 0; j < len(result); j++ {
		for i := 0; i < len(result[0]); i++ {
			_, err23 := w.WriteString(strconv.Itoa(result[i][j]) + " ") // chiffre puis espace
			if err23 != nil {
				fmt.Println("arrive pas a convertir strconv")
			}
		}
		_, err24 := w.WriteString("\n") // newline
		if err24 != nil {
			fmt.Println("arrive pas a ecrire apres calc")
		}
	}

	w.Flush() // apparament il faut ca mais vu qu'on ecrit pas apres peut etre pas

	filetosend, _ := os.Open("result.txt")
	sentdata, _ := io.Copy(c, filetosend)
	fmt.Print("serv envoyé :", sentdata)

}

func main() { //Scanner fichier de la doc golang

	lst, err := net.Listen("tcp", "localhost:8000")
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
