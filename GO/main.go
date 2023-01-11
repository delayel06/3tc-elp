package main

import (
	"bufio"
	"fmt"
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

			n, err := strconv.Atoi(s) // https://pkg.go.dev/strconv on convertit text en int parce que c'est plus cool

			if err != nil {
				return nil
			}
			row = append(row, n)
		}

		matrice = append(matrice, row)
	}

	// Return the matrix
	return matrice
}

func main() { //Scanner fichier de la doc golang

	matone := read("mat1.txt")

	fmt.Print(matone)
	mattwo := read("mat2.txt")

	fmt.Print(mattwo)

}
