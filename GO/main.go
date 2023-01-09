package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrice [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		var row []int
		for _, s := range line {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}

		matrice = append(matrice, row)
	}

	// Return the matrix
	return matrice, nil
}

func main() { //Scanner fichier de la doc golang

	matone, err := read("mat1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(matone)
	mattwo, err := read("mat2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(mattwo)

}
