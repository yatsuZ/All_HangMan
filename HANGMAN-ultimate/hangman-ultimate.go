package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Proposition      []string   // toute les lettre que j'ai proposé
}

func main() {
	//verifier que j'ai un argument sinon mon hang man ne marche pas
	if len(os.Args) > 1 {
		//si l'argument est words.txt je joue au pendu avec le document texte de words
		//les conditon pour fair un document teste est qu'il faut 1 mot par ligne
		if os.Args[1] == "words.txt" {
			//je crée mon objet pendu et je remplis ces argument
			pendu := &HangManData{ToFind: choisirunmots(), Attempts: 10, HangmanPositions: listeduPendu(), Proposition: nil}
			devine(pendu)
			//fmt.Println(pendu.ToFind)
			//jeux le pendu execute le jeux
			jeux_le_pendu(pendu)
			//si mon argument est --starWith cela signifie que mon joueur veux reprendre une partie
		} else if os.Args[1] == "--startWith" {
			//mon fonction reprend la sauvegarde et si ma document texte sauvegarde n'est
			//pas vide je peux exucuté mon sinon je retourne un message d'erreur
			sauv := Load()
			if sauv.Word != "" {
				fmt.Println("Welcome Back, you have", sauv.Attempts, "attempts remaining.")
				jeux_le_pendu(sauv)
			} else {
				fmt.Println("y a pas de sauvegarde bg")
			}
		} else {
			fmt.Println("il faut fair avec \"words.txt\" et pas autre chose ou reprendre une sauvegarde >:(")
		}
	} else {
		fmt.Println("TU NA PAS D'ARGUEMENT PATATE")
	}
}

func devine(t *HangManData) {
	//crée hangman .Word
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

func (t *HangManData) save() {

	// create and populate a map from dummy JSON data

	// convert map to Person struct

	jsonData, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	// sanity check

	// write to JSON file

	jsonFile, err := os.Create("./save.txt")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())

}

func Load() *HangManData {
	var t *HangManData
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("save.txt")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &t)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return &HangManData{Word: t.Word, ToFind: t.ToFind, Attempts: t.Attempts, HangmanPositions: t.HangmanPositions, Proposition: t.Proposition}
}

func verification(a string, L []string) bool {
	//verifier qu'une lettre que je propose n'a pas deja etais proposé
	//verifier que a n'est pas dans la liste L
	for _, i := range L {
		if i == a {
			return true
		}
	}
	return false
}

func changement(s string, changement rune, p int) string {
	// transforme un caractére de string en rune
	//en l'occurence là convontion veux que soit des lettre
	a := []rune(s)
	for i := range a {
		if p == i {
			a[i] = changement
		}
	}
	return string(a)
}

func jeux_le_pendu(t *HangManData) {
	//je definie plusieur variable que je vais utilisé dans une boucle
	var scanner *bufio.Scanner // le module qui me permet de lire ligne par ligne dans dans les doc .txt
	var input string           //la variable intermediere qui me permetra de stocké scanner
	position := 0
	pastrouver := true // mon bool qui pemetra de dire si j'ai trouvé une lettre
	moin2 := false     // si je propose un mot c'est cette variable qui pemetra de mettre moi 2
	gagne := false     // bool qui dis si j'ai gagné
	var tab [][]string // mot en assci art
	//fin des definition debut de la boucle infinie qui est divisé en 4 partie
	for true {
		///////////////////////////////////////////////////////////////////////////////////////////////////////////// partie 1 les resultat de nos action
		if t.ToFind == string(t.Word) || gagne {
			fmt.Println("Congrats !")
			break
		}
		if pastrouver == false {
			if moin2 {
				t.Attempts--
			}
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
		//////////////////////////////////////////////////////////////////////////////////////////////////////////////////part2 action du joueur
		fmt.Printf("Choose : ")
		pastrouver = false
		moin2 = false
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		if input == "STOP" {
			fmt.Println("Game Saved in save.txt.")
			t.save()
			return
		}
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////part3 verifier que notre chois est legite pour advenced fealure
		if len(input) == 1 && len(t.Proposition) != 0 {
			for verification(input, t.Proposition) {
				fmt.Println("tu à deja proposé là lettre,", input, "donc repropose une autre lettre svp bg")
				scanner = bufio.NewScanner(os.Stdin)
				scanner.Scan()
				input = scanner.Text()
				if input == "STOP" {
					fmt.Println("Game Saved in save.txt.")
					t.save()
					return
				}
			}
		}
		t.Proposition = append(t.Proposition, string(input[0]))
		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////part4 es que notre choix nous fais gagné ou perdre
		for i, j := range t.ToFind {
			if len(input) > 1 {
				if len(input) != len(t.ToFind) {
					moin2 = true
				} else {
					if string(j) != string(input[i]) {
						moin2 = true
						pastrouver = true
						gagne = false
						break
					} else {
						gagne = true
					}
				}
			} else {
				if j == rune(input[0]) {
					t.Word = changement(t.Word, rune(j), i)
					pastrouver = true
				}
			}
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////on recomence tant que j'ai pas gagné ou perdu ou sauvegardé
	}
}

func choisirunmots() string {
	//crée Hangman . To Find
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
	//crée les different affichage du pendu
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

func lettreassci(s rune) []string {
	//fonction qui permet de convertir une rune en caractére assci art
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
	//permeé d'affiché une chaine de caractére assci
	for i := range TAB[0] {

		for j := range TAB {
			fmt.Printf(TAB[j][i])
		}
		fmt.Println()
	}
}

//////117
