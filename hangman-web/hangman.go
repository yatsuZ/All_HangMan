package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//Here we're using enumerations to decouple the state of the game with web implementation
const (
	StateError = 100
	StateOK    = 101

	StateWon     = 200
	StateLost    = 201
	StateReset   = 202
	StateNothing = 203

	StateUserLoggedIn = 300
)

var gameStateText = map[int]string{
	StateError: "An Error must have occured !",
	StateOK:    "Everthing is okay",

	StateWon:          "You won, Want to play again ?",
	StateLost:         "You lost, Play again !",
	StateReset:        "Good luck !",
	StateNothing:      "You must provide a valide input",
	StateUserLoggedIn: "Hello ",
}

func GameStateText(code int) string {
	return gameStateText[code]
}

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Bonomme          string     //position exatce
}

func RemplirHangman(txt string) *HangManData {
	pendu := &HangManData{ToFind: Choisirunmots(txt), Attempts: 10, HangmanPositions: ListeduPendu()}
	Devine(pendu)
	pendu.Bonomme = ""
	return pendu
}

//Game
/*cette fonction execute le jeux le pendu
et defenis les variable de HangManData
si pas d'argument affiche une erreur */
func Game() {
	if len(os.Args) > 1 {
		pendu := RemplirHangman("")
		if pendu.ToFind == "" {
			return
		}
		fmt.Println("\nsi tu apuis sur entré avec aucune valeur tu qui la partie")
		fmt.Println("on accepte que les lettre en minuscule et pas de caractére speciaux")
		fmt.Println()
		JeuxLePendu(pendu)

	} else {
		fmt.Println("TU NA PAS D'ARGUEMENT PATATE")
	}
}

//Changement
//change un string
//s= le string changement = la rune qu'on remplace et p= la position
func Changement(s string, changement rune, p int) string {
	a := []rune(s)
	for i := range a {
		if p == i {
			a[i] = changement
		}
	}
	return string(a)
}

//JeuxLePendu
//cette fonction execute le jeux et sarrete
//apres  tentative ou quand le joueur trouve le mot
func JeuxLePendu(t *HangManData) {
	/////////////////////////////////////////////////////////////////////////////////////////////
	var scanner *bufio.Scanner
	var input string
	pastrouver := true
	position := 0
	///////////////////////////////////////////////////////////////////////////////////////////
	for {
		if !pastrouver {
			t.Attempts--
			fmt.Println("Not present in the word, ", t.Attempts, " attempts remaining")
			fmt.Println(t.HangmanPositions[position])
			position++
		} else {
			fmt.Println(string(t.Word))
		}
		if t.Attempts == 0 {
			fmt.Println("perdu")
			break
		}
		if t.ToFind == string(t.Word) {
			fmt.Println("Congrats !")
			break
		}
		/////////////////////////////////////////////////////////////////////////////////////////////
		fmt.Printf("Choose : ")
		pastrouver = false
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		if len(input) == 0 {
			fmt.Println("tu quite la partie aurevoir")
			return
		}
		//////////////////////////////////////////////////////////////////////////////////////////////
		for i, j := range t.ToFind {
			if j == rune(input[0]) {
				t.Word = Changement(t.Word, rune(j), i)
				pastrouver = true
			}
		}
	}

}

//Devine
//definie Word
func Devine(t *HangManData) {
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

//Choisirunmots
//definie ToFind
func Choisirunmots(doc string) string {
	if doc == "" {
		doc = os.Args[1]
	}
	file, err := os.Open(doc)
	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var tab []string
		for scanner.Scan() {
			tab = append(tab, scanner.Text())
		}
		if tab == nil {
			fmt.Println("document texte invalide")
			return ""
		}
		rand.Seed(time.Now().Unix())
		for _, ligne := range tab {
			for _, lettre := range ligne {
				if string(lettre) == " " {
					fmt.Println("document texte invalide")
					return ""
				}

			}
		}
		random := rand.Intn(len(tab))
		return tab[random]
	} else {
		fmt.Println("il y a un problème pas de document texte :/")
	}
	return ""
}

//ListeduPendu
//defnier Hangman Position
func ListeduPendu() [10]string {
	var t [10]string
	file, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tab string
	for scanner.Scan() {
		tab = tab + scanner.Text()
		tab = tab + "\n"
	}
	//	fmt.Printf(tab)
	for i := 0; i < 10; i++ {
		t[i] = tab[71*i : 71*i+70]
	}
	return t
}
