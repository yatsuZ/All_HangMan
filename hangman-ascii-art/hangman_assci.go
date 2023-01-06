package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	if len(os.Args) > 1 {
		//fmt.Println(os.Args[1])
		if os.Args[1] == "words.txt" {

			pendu := &HangManData{ToFind: choisirunmots(), Attempts: 10, HangmanPositions: listeduPendu()}
			devine(pendu)
			/*			var tab [][]string
						for _, mot := range pendu.ToFind {
							tab = append(tab, lettreassci(rune(mot)))
						}
						affichageassci(tab)

						//	fmt.Println(string(pendu.Word))
			*/
			jeux_le_pendu(pendu)
		} else {
			fmt.Println("il faut fair avec \"words.txt\" et pas autre chose >:(")
		}
	} else {
		fmt.Println("TU NA PAS D'ARGUEMENT PATATE")
	}
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

func changement(s string, changement rune, p int) string {
	a := []rune(s)
	for i := range a {
		if p == i {
			a[i] = changement
		}
	}
	return string(a)
}

func jeux_le_pendu(t *HangManData) {
	var scanner *bufio.Scanner
	var input string
	pastrouver := true
	position := 0
	var tab [][]string
	for true {
		if pastrouver == false {
			t.Attempts--
			fmt.Println("Not present in the word, ", t.Attempts, " attempts remaining")
			fmt.Println(t.HangmanPositions[position])
			position++
		} else {
			tab = nil
			for _, mot := range t.Word {
				tab = append(tab, lettreassci(rune(mot)))
			}
			affichageassci(tab)
		}
		if t.Attempts == 0 {
			fmt.Println("perdu")
			break
		}
		if t.ToFind == string(t.Word) {
			fmt.Println("Congrats !")
			break
		}
		fmt.Printf("Choose : ")
		pastrouver = false
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		for i, j := range t.ToFind {
			if j == rune(input[0]) {
				t.Word = changement(t.Word, rune(j), i)
				pastrouver = true
			}
		}
	}

}
func devine(t *HangManData) {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(len(t.ToFind))
	for i := range t.ToFind {
		if i != random {
			t.Word = t.Word + "_"
		} else {
			t.Word = t.Word + string(t.ToFind[i])
		}
	}
}

func choisirunmots() string {
	file, err := os.Open(os.Args[1])
	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var tab []string
		for scanner.Scan() {
			tab = append(tab, scanner.Text())
		}
		rand.Seed(time.Now().Unix())
		random := rand.Intn(84)
		return tab[random]
	} else {
		fmt.Println("il y a un problème pas de document texte :/")
	}
	return ""
}

func listeduPendu() [10]string {
	var t [10]string
	file, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tab string
	for scanner.Scan() {
		tab = tab + scanner.Text()
		tab = tab + "\n"
	}
	for i := 0; i < 10; i++ {
		t[i] = tab[71*i : 71*i+70]
	}
	return t
}
