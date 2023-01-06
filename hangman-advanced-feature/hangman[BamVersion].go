package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	tHeGame(aGamePrepare())
}

func tHeGame(dataGame *HangManData) {
	var findletter, bl bool
	var cfiniNonOui bool
	aRuneToFindCHWord := []rune(dataGame.ToFind)
	aRuneCHWord := []rune(dataGame.Word)
	fmt.Println("Good Luck, you have 10 attempts")

	for {
		if dataGame.Attempts == 0 {
			fmt.Println("Perdu !")
			break
		}
		if dataGame.Attempts > 0 {
			if findletter != false || dataGame.Attempts == 10 {
				fmt.Println(dataGame.Word)
			}
			findletter = false
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("Choose: ")
			scanner.Scan()
			input := scanner.Text()
			////////////////////// IF HE INPUTED AN ALAREADY SHOWN LETTER ///////////////////////////////////////////
			if len(input) == 1 {
				for _, p := range aRuneCHWord {
					if rune(input[0]) == p {
						fmt.Println("Error ! already shown")
						bl = true
					}
				}
				if bl {
					continue
				}
			}
			//////////////////////////        F  I   N       /////////////////////////////////////////////////
			///////////// la condition de rentrer un mots, si elle est vari ça fini la game ,sinon t'as deux tentavies de moins
			if len(input) > 1 {
				if string(dataGame.ToFind) == string(input) {
					fmt.Println("Congrats !")
					break
				} else {
					dataGame.Attempts -= 2
					fmt.Println("The word is false now you just have ", dataGame.Attempts, "attempts remaining!")
					fmt.Println(dataGame.HangmanPositions[len(dataGame.HangmanPositions)-1-dataGame.Attempts])
					fmt.Println()
					continue
				}
			}
			///////////////////////// fin de la condition /////////////////////////////////////////////////////////////////////
			rStdin := []rune(input)
			for i, p := range aRuneToFindCHWord {
				if rStdin[0] == p {
					findletter = true
					aRuneCHWord[i] = p
				}
			}
		}
		if findletter == false {
			dataGame.Attempts -= 1
			fmt.Println("Not present in the word", dataGame.Attempts, "attempts remaining")
			fmt.Println(dataGame.HangmanPositions[len(dataGame.HangmanPositions)-1-dataGame.Attempts])
			fmt.Println()
		}
		if findletter == true {
			dataGame.Word = string(aRuneCHWord)
			cfiniNonOui = true
			for _, p := range aRuneCHWord {
				if p == '_' {
					cfiniNonOui = false
				}
			}
		}
		if cfiniNonOui {
			fmt.Println("Congrats !")
			break
		}

	}
}

func aGamePrepare() *HangManData {
	var a *HangManData
	words, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("You've A problem small dick")
	} else {
		theWordsList := aSliceofBytesToStringsSlice(words)
		rand.Seed(time.Now().Unix())
		r := rand.Intn(len(theWordsList) - 1)
		var tHangmaData HangManData
		tHangmaData.ToFind = theWordsList[r]
		if byte(theWordsList[r][len(theWordsList[r])-1]) == 13 {
			tHangmaData.ToFind = tHangmaData.ToFind[0 : len(tHangmaData.ToFind)-1]
		}
		tHangmaData.Attempts = 10
		for i := 0; i < len(tHangmaData.HangmanPositions); i++ {
			tHangmaData.HangmanPositions[i] = aHangmantxtToStringsSlice()[i]
		}
		arunechwo := []rune(theWordsList[r])
		wordstate := make([]rune, len(arunechwo))
		for i := range arunechwo {
			if i != len(arunechwo)/2-1 {
				wordstate[i] = '_'
			} else {
				wordstate[i] = arunechwo[i]
			}
		}
		wordstat := make([]rune, len(wordstate)-1)
		for i := range wordstat {
			wordstat[i] = wordstate[i]
		}
		tHangmaData.Word = string(wordstat)
		a = &tHangmaData
	}
	return a
}

func aSliceofBytesToStringsSlice(sob []byte) []string {
	var rsos []string
	for i := 0; i < len(sob); i++ {
		s := ""
		c := 0
		if sob[i] == 10 || i == len(sob)-1 {
			for j := (i - 1); j >= 0 && sob[j] != 10; j-- {
				if j == len(sob)-2 && c == 0 {
					j += 1
					c++
				}
				s = string(sob[j]) + s
			}
			rsos = append(rsos, s)
		}
	}
	return rsos
}

func aHangmantxtToStringsSlice() []string {
	var hsos []string
	s := " "
	hangposes, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println("you've got an error in hangaman thing jerk off")
	} else {
		for i := 0; i < len(hangposes); i++ {
			if hangposes[i] == 10 && hangposes[i+1] == 10 {
				i += 2
				hsos = append(hsos, s)
				s = " "
			} else {
				s += string(hangposes[i])
			}
		}
	}
	return hsos
}
