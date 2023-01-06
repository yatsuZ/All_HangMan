package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	t := "aok sa marche"
	var tab [][]string
	for _, v := range t {
		tab = append(tab, lettreassci(rune(v)))
	}
	affichageassci(tab)

}

func lettreassci(s rune) []string {
	file, err := os.Open("standard.txt")
	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var tab []string
		counter := 0
		position := int(s - 32)
		for scanner.Scan() {
			if counter >= 9*position+2 && counter < 9*position+9 {
				tab = append(tab, scanner.Text())
			}
			counter++
		}
		return tab
	}
	return nil
}
func affichageassci(TAB [][]string) {

	for i := range TAB[0] {

		for j := range TAB {
			fmt.Printf(TAB[j][i])
		}
		fmt.Println()
	}
}
