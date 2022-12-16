package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangData struct {
	InputLetter string
	WordTF      string
	HiddenWord  string
	Attempts    int
	Level       string
	LevelSet    bool
}

// FONCTIONS D'INITIALISATIONS

func RandomWord(fileName []string, n int) string { //Permet de récupérer un mot aléatoirement dans un fichier
	content, _ := os.Open(fileName[n])
	rand.Seed(time.Now().UnixNano())
	randomLine := rand.Intn(CountLines(fileName, n))
	if randomLine == 0 {
		randomLine = 1
	}
	compt := 0
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		compt++
		if compt == randomLine {
			return scanner.Text()
		}
	}
	return "unreachable point"
}

func CountLines(FName []string, n int) int {
	// Compte le nombre de ligne dans le fichier envoyé en argument
	content, _ := os.Open(FName[n])
	defer content.Close()
	scanner := bufio.NewScanner(content)
	lineNbrs := 1
	for scanner.Scan() {
		lineNbrs++
	}
	return lineNbrs
}

func ChooseFile(Pts *HangData) int {
	fmt.Println(Pts.Level)
	switch Pts.Level {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	}
	return 0
}

func ChangeWord(Pts *HangData) string {
	// Nombre random de lettre en suivant le schéma : n/2 - 1
	rand.Seed(time.Now().UnixNano())
	n := len(Pts.WordTF)/2 - 1

	tabWTF := []rune(Pts.WordTF)
	if n != 0 { // Modifier le mot en underscore uniquement sur les lettres différentes des n random
		for i := 0; i < len(tabWTF)-n; i++ {
			index := rand.Intn(len(Pts.WordTF))
			if i != index {
				tabWTF[i] = '_'
			}
		}
	} else { // ajout des underscores pour les mots avec n = 0 (genre lit, ou mot ...)
		for i := 0; i < len(Pts.WordTF); i++ {
			tabWTF[i] = '_'
		}
	}

	// Modifier la valeur dans la variable Pts.WordTF après ajout des '_'
	return string(tabWTF)
}

// FONCTIONS DE JEUX

func CheckLetter(Pts *HangData) int {
	for i := 0; i < len(Pts.WordTF); i++ {
		if Pts.InputLetter[0] == Pts.WordTF[i] {
			return i
		}
	}
	return 999 // cas d'erreur
}

// Appeler pour vérifier l'input et modifier le mot
func ModifHW(Pts *HangData) {
	i := CheckLetter(Pts)
	stToR := []rune(Pts.HiddenWord)
	if i != 999 {
		for j := 0; j < len(Pts.HiddenWord); j++ {
			if stToR[j] == '_' && i == j {
				stToR[j] = rune(Pts.InputLetter[0])
			}
		}
	}
	Pts.HiddenWord = string(stToR)
}
