package main

import (
	"fmt"
	"hangman/hangman"
	"net/http"
	"text/template"
)

//CREATION DES TEMPLATES
//Le pointeur de la structure de jeux est donné en argument et executé dans la template pour afficher les données sur la page web
//et aussi récupérer les données du form / les modifier facilement sur les fonctions de jeux

func HandleHomePage(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	tmp, _ := template.ParseFiles("./page/index.html")
	tmp.Execute(rw, Pts)
}

func HandleGamePage(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	tmp, _ := template.ParseFiles("./page/game.html")
	tmp.Execute(rw, Pts)
}

func HandleCreditsPage(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	tmp, _ := template.ParseFiles("./page/credits.html")
	tmp.Execute(rw, Pts)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	fp := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fp))
	//CREATION D'UN POINTEUR DE STRUCTURE VIDE (qui servira tout au long du jeu)

	Pts := &hangman.HangData{}
	//Exactement c'est un pointeur d'une structure définie dans le package hangman
	//Le fait de passer par un pointeur de structure du package hangman permets de mettre en
	//argument des fonctions dans le package hangman un pointeur de cette meme structure
	//et ainsi de ne pas redéfinir les fonctions dans le serveur
	//ex : On retrouve dans le dossier hangman/hangman.go une fonction avec comme paramètres : function Test(Pts *HangData) {...}

	//PREMIERE INITIALISATION DE STRUCTURE
	InitialiseStruct(Pts)

	//CREATION DES ROUTES

	//Route de front de la page d'accueil
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		HandleHomePage(rw, r, Pts)
	})

	//Route de front, qui récupère le level de la route "/" et Initialise la structure
	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(Pts.Level)
		if Pts.Level != r.FormValue("level") && Pts.Level != "" {
			r.FormValue("letter")
			Pts.Level = r.FormValue("level")
			InitialiseStruct(Pts)
		}
		Pts.LevelSet = true
		HandleGamePage(rw, r, Pts)
	})

	//Route de back, juste pour lancer les fonctions qui modifient les données de la structure puis redirection
	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) {
		Pts.InputLetter = r.FormValue("letter")
		hangman.ModifHW(Pts)
		http.Redirect(rw, r, "/game", http.StatusFound)
	})

	//Route des crédits
	http.HandleFunc("/credits", func(rw http.ResponseWriter, r *http.Request) {
		HandleCreditsPage(rw, r, Pts)
	})

	// INTEGRATION DES ASSETS (CSS, IMAGES, FONTS)

	http.ListenAndServe(":8080", nil)
}

//INITIALISATION DE LA STRUCTURE DU JEU
func InitialiseStruct(Pts *hangman.HangData) {
	FilesName := []string{"hangman/words.txt", "hangman/words1.txt", "hangman/words2.txt"}
	Pts.Attempts = 10
	Pts.WordTF = hangman.RandomWord(FilesName, hangman.ChooseFile(Pts))
	Pts.HiddenWord = hangman.ChangeWord(Pts)
	Pts.LevelSet = false
}
