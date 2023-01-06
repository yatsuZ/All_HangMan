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

			pendu := &HangManData{ToFind: Choisirunmots(), Attempts: 10, HangmanPositions: ListeduPendu()}
			Devine(pendu)
			//	fmt.Println(pendu.ToFind)
			//	fmt.Println(string(pendu.Word))
			Jeux_le_pendu(pendu)
		} else {
			fmt.Println("il faut fair avec \"words.txt\" et pas autre chose >:(")
		}
	} else {
		fmt.Println("TU NA PAS D'ARGUEMENT PATATE")
	}
}

func Changement(s string, changement rune, p int) string {
	a := []rune(s)
	for i := range a {
		if p == i {
			a[i] = changement
		}
	}
	return string(a)
}

func Jeux_le_pendu(t *HangManData) {
	var scanner *bufio.Scanner
	var input string
	pastrouver := true
	position := 0
	for true {
		if pastrouver == false {
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
		fmt.Printf("Choose : ")
		pastrouver = false
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		for i, j := range t.ToFind {
			if j == rune(input[0]) {
				t.Word = Changement(t.Word, rune(j), i)
				pastrouver = true
			}
		}
	}

}
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

func Choisirunmots() string {
	t := time.Now().Nanosecond()
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
		m := time.Now().Nanosecond()
		fmt.Println(m - t)
		return tab[random]
	} else {
		fmt.Println("il y a un problème pas de document texte :/")
	}
	return ""
}

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

//////117 à

/*
ancien code
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	T := choisirunmots()
	fmt.Println(T)
	jeux_lependu(T)
}

func jeux_lependu(mots string) {
	fmt.Println("Good Luck, you have 10 attempts.")
	/////////////////////////////////////////////////////crée là liste S à deviné et Devine les proposition du joueur
	S := []rune(mots)
	Devine := make([]rune, len(S))
	//////////////////////////////////////////////////////j'ajoute les element dans devine et le reste c'est _
	rand.Seed(time.Now().Unix())
	random := rand.Intn(len(S))
	for i := range S {
		if i != random {
			Devine[i] = '_'
		} else {
			Devine[i] = S[i]
		}
	}
	/////////////////////////////////////////////////////// crée les variable dans ma boucle infinie
	var scanner *bufio.Scanner
	var input string
	pastrouver := true
	fini := false
	count := 0
	remainingAttempts := 10
	//////////////////////////////////////////////////////// je crée ma boucle infini les 1ier condition
	//////////// si pas trouvé et faux on mets 0 sinon si count = 10 perdu et si fini alors gagne
	for 1 > 0 {
		if pastrouver == false {
			count += 1
			remainingAttempts = 10 - count
			fmt.Println("Not present in the word, ", remainingAttempts, " attempts remaining")
			affichLpenduavecDOC(count)
		} else {
			fmt.Println(string(Devine))
		}
		////////////////////////////////////////////////////////metre fin
		if count == 10 {
			fmt.Println("perdu")
			break
		}
		if fini {
			fmt.Println("Congrats !")
			break
		}
		//////////////////////////////////////////////////////j'affiche Devine et  j'update les valeur
		fmt.Printf("Choose : ")
		pastrouver = false
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		fini = true
		/////////////////////////////////////////////////////// je crée une boucle for pour voir si mon inpout est egal
		for i, j := range S {
			if j == rune(input[0]) {
				Devine[i] = S[i]
				pastrouver = true
			}
			if Devine[i] == S[i] && fini {
				fini = true
			} else {
				fini = false
			}
		}
	}
}

func choisirunmots() string {
	file, _ := os.Open("words.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tab []string
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}
	rand.Seed(time.Now().Unix())
	random := rand.Intn(84)
	return tab[random]
}

func affichLpenduavecDOC(nvx int) {
	file, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tab string
	for scanner.Scan() {
		tab = tab + scanner.Text()
		tab = tab + "\n"
	}
	fmt.Println(tab[71*(nvx-1) : 71*(nvx-1)+70])
}

<<<<<<< HEAD
//////117
func affichagependu(niveau int) {
	if niveau == 0 {
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
	} else if niveau == 1 {
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("=========")
	} else if niveau == 2 {
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 3 {
		fmt.Println("  +---+  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 4 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 5 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 6 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println("  |   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 7 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 8 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 9 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println(" /    |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	} else if niveau == 10 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println(" / \\  |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	}
}
*/
