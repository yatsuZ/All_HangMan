package main

import (
	"fmt"
	"html/template"
	"lib"
	"net/http"
)

//Data
//Defining a data model that will be injected in the template
type Data struct {
	username string
	Score    int
	Pendu    *lib.HangManData
	Alert    string
}

//Initializing the Data
var data = Data{
	Score: 0,
	Pendu: nil,
	Alert: "",
}

func main() {
	//index
	http.HandleFunc("/", Index)
	http.HandleFunc("/start", IndexHandler)
	http.HandleFunc("/select", home)
	http.HandleFunc("/choice", HomeHandler)
	http.HandleFunc("/game", Game)
	//HTTP endpoint to handle POST requests
	http.HandleFunc("/hangman", GameHandler)

	//select

	//fileserver := http.FileServer(http.Dir("./templates"))
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	err := http.ListenAndServe("localhost:555", nil)
	if err != nil {
		return
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		return
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//Ensure that the endpoint only gets POST requests
	if r.Method != "POST" {
		// http.Error(w, "Bad request - Go away!", 405)
		tmpl2, _ := template.ParseGlob("./templates/*.gohtml")
		err := tmpl2.ExecuteTemplate(w, "br", nil)
		if err != nil {
			return
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			return
		}
		if r.FormValue("start") == "Start" {
			data.Alert = ""
			data.username = r.FormValue("username")
			fmt.Println(data.username)
			data.Alert = lib.GameStateText(lib.StateUserLoggedIn) + data.username
		}

		http.Redirect(w, r, "/select", http.StatusSeeOther)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "select", nil)
	if err != nil {
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//Ensure that the endpoint only gets POST requests
	if r.Method != "POST" {
		// http.Error(w, "Bad request - Go away!", http.StatusMethodNotAllowed)
		tmpl2, _ := template.ParseGlob("./templates/*.gohtml")
		err := tmpl2.ExecuteTemplate(w, "br", nil)
		if err != nil {
			return
		}
	} else {

		err := r.ParseForm()
		if err != nil {
			return
		}

		if r.FormValue("doc1") == "EASY peasy" {
			data.Alert = ""

			fmt.Println(data.username)
			data.Alert = lib.GameStateText(lib.StateUserLoggedIn) + data.username
			data.Pendu = lib.RemplirHangman("words.txt")
		} else if r.FormValue("doc2") == "NORMAL" {
			data.Alert = ""

			fmt.Println(data.username)
			data.Alert = lib.GameStateText(lib.StateUserLoggedIn) + data.username
			data.Pendu = lib.RemplirHangman("words2.txt")
		} else if r.FormValue("doc3") == "HARD" {
			data.Alert = ""

			fmt.Println(data.username)
			data.Alert = lib.GameStateText(lib.StateUserLoggedIn) + data.username
			data.Pendu = lib.RemplirHangman("words3.txt")
		} else {
			http.Error(w, "Unsupported POST param", 405)
		}

		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "game", data)
	if err != nil {
		return
	}

}

func GameHandler(w http.ResponseWriter, r *http.Request) {

	//Ensure that the endpoint only gets POST requests
	if r.Method != "POST" {
		// http.Error(w, "Bad request - Go away!", 405)
		tmpl2, _ := template.ParseGlob("./templates/*.gohtml")
		err := tmpl2.ExecuteTemplate(w, "br", nil)
		if err != nil {
			return
		}
	} else {

		err := r.ParseForm()
		if err != nil {
			return
		}
		if r.FormValue("guess") == "Guess" {
			data.Alert = ""
			lettre := r.FormValue("lettre")
			fmt.Println(lettre)
			gameState := jeuxLePenduWeb(data.Pendu, lettre)
			if gameState != lib.StateOK {
				data.Alert = lib.GameStateText(gameState)
				if gameState == lib.StateWon {
					if data.Score != 1 {
						data.Score = data.Score / 2
					}
					tmpl2, _ := template.ParseGlob("./templates/*.gohtml")
					err = tmpl2.ExecuteTemplate(w, "win", data)
					if err != nil {
						return
					}
				}
				if gameState == lib.StateLost {
					tmpl2, _ := template.ParseGlob("./templates/*.gohtml")
					err = tmpl2.ExecuteTemplate(w, "lose", data)
					if err != nil {
						return
					}
				}
			}
		} else if r.FormValue("replay") == "Play again !" {

			data.Pendu = nil
			data.Alert = lib.GameStateText(lib.StateReset)
			http.Redirect(w, r, "/select", http.StatusSeeOther)
		} else {
			http.Error(w, "Unsupported POST param", 405)
		}

		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

//Note: we might use Enums (rather than strings) as return type to model different states of the game
//This approach is simpler though xD
func jeuxLePenduWeb(pendu *lib.HangManData, lettre string) int {

	if pendu == nil && lettre == "" {
		return lib.StateError
	} else if lettre == "" {
		return lib.StateNothing
	} else if pendu.Attempts == 0 {
		return lib.StateLost
	}
	pastrouver := true
	for key, value := range pendu.ToFind {
		if value == rune(lettre[0]) {
			pendu.Word = lib.Changement(pendu.Word, rune(value), key)
			pastrouver = false
		}
	}
	if pastrouver {
		pendu.Attempts--
		pendu.Bonomme = pendu.HangmanPositions[10-pendu.Attempts-1]
	}
	if pendu.ToFind == pendu.Word {
		if data.Score != 1 {
			data.Score += pendu.Attempts
		}
		return lib.StateWon
	} else if pendu.Attempts == 0 && pastrouver {
		return lib.StateLost
	}
	return lib.StateOK
}

////question 1 pk il maffiche alor que je n'ai rien envoyé
////et pk le sever commence direct alors que j'ai rien envoyé
