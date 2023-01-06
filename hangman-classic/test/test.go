package main

import (
	"bufio"
	"os"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func main() {
	listeduPendu()
	//	fmt.Println(T)
}

func listeduPendu() [10]string {
	var t [10]string
	file, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tab string
	for scanner.Scan() {
		tab = tab + scanner.Text()
		//		fmt.Print(tab)
		tab = tab + "\n"
	}
	//	fmt.Printf(tab)
	for i := 0; i < 10; i++ {
		t[i] = tab[71*i : 71*i+70]
		//		fmt.Println(i)
		//		fmt.Println(t[i])
	}
	return t
}
