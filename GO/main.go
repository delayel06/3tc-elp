package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func strtointsplice(s string) [][]int {
	rows := strings.Split(s, "\n")
	//rows marche
	var splice [][]int
	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], " ")
		var row []int // creer un truc a ajouter apres
		for j := 0; j < len(cols); j++ {

			re := regexp.MustCompile(`\r?`)
			cols[j] = re.ReplaceAllString(cols[j], "")
			// windows pourquoi ya \r a la fin de tes lignes
			// ca m'a pris 3h a comprendre ca

			num, _ := strconv.Atoi(cols[j])

			if j == 2 {

			}

			row = append(row, num) // doit append parce que on a pas initializé de taille -> dynamique
			//on ajout a chaque row le num puis chaque row jusqu'a la fin
		}
		splice = append(splice, row)
	}

	return splice
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
			fmt.Print("Calcul matrice..\n")   // test
			// tableaux modifiés globallement comme java
		}
	}

	return result

}

func tcp(c net.Conn) {
	var stockage = make([]byte, 4096)
	data, erreur := c.Read(stockage) // lis les données et les stock dans stockage

	if erreur != nil {
		fmt.Println("Arrive pas a lire data")
	} else {
		fmt.Println("(Server - data 1) Bytes reçues: ", data)
	}

	matreceived := string(stockage[:data])

	var stockage2 = make([]byte, 4096)
	data2, erreur2 := c.Read(stockage2) // lis les données et les stock dans stockage

	if erreur2 != nil {
		fmt.Println("Arrive pas a lire data")
	} else {
		fmt.Println("(Server - data 2) Bytes reçues: ", data)
	}

	matreceived2 := string(stockage2[:data2])

	matrice1 := strtointsplice(matreceived)
	matrice2 := strtointsplice(matreceived2)

	if len(matrice1[0]) != len(matrice2) {
		fmt.Println("Error ! Les dimensions des matrices ne sont pas bonnes à multiplier.")
		c.Close()
		return
	}

	fmt.Println("Matrice 1: ", matrice1, "\n", "et 2: ", matrice2)
	resulttemp := process(matrice1, matrice2)
	wg.Wait()
	result := resulttemp

	fmt.Println("Résultat:", result)

	var outfile, errr = os.Create("result.txt")
	if errr != nil {
		return
	}

	w := bufio.NewWriter(outfile)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
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

	w.Flush()

	filetosend, _ := os.Open("result.txt")
	sentdata, _ := io.Copy(c, filetosend)

	fi, _ := filetosend.Stat()
	if sentdata != fi.Size() {
		fmt.Println("Erreur ! Je n'ai pas envoyé le bon nombre de bytes.")
		return
	}

	fmt.Print("Bytes envoyées: ", sentdata, "\n")

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
