package main

// ============================================================
//  ui.go
//  Tout ce qui sert a AFFICHER des choses proprement dans le
//  terminal : cadres ASCII, barres de vie, couleurs, lecture
//  du clavier.
//  On met tout ca ici pour ne pas repeter le meme code partout.
// ============================================================

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Le lecteur de clavier. On le cree une seule fois pour tout le jeu.
var lecteur = bufio.NewReader(os.Stdin)

// Les couleurs ANSI (elles marchent dans presque tous les terminaux)
const (
	Reset  = "\033[0m"
	Gras   = "\033[1m"
	Rouge  = "\033[31m"
	Vert   = "\033[32m"
	Jaune  = "\033[33m"
	Bleu   = "\033[34m"
	Violet = "\033[35m"
	Cyan   = "\033[36m"
	Gris   = "\033[90m"
)

// Largeur interieure de nos cadres
const largeurCadre = 54

// nettoyerEcran efface le terminal
func nettoyerEcran() {
	fmt.Print("\033[H\033[2J")
}

// centrer met du vide a gauche et a droite pour centrer un texte
func centrer(texte string, largeur int) string {
	taille := len([]rune(texte)) // len([]rune) pour bien compter les accents
	if taille >= largeur {
		return texte
	}
	gauche := (largeur - taille) / 2
	droite := largeur - taille - gauche
	return strings.Repeat(" ", gauche) + texte + strings.Repeat(" ", droite)
}

// titre affiche un beau cadre avec un titre au milieu
// ATTENTION : pas d'emoji ici, sinon le cadre est decale
func titre(texte string) {
	fmt.Println(Cyan + "╔" + strings.Repeat("═", largeurCadre) + "╗" + Reset)
	fmt.Println(Cyan + "║" + Reset + Gras + centrer(texte, largeurCadre) + Reset + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╚" + strings.Repeat("═", largeurCadre) + "╝" + Reset)
}

// separateur affiche une simple ligne
func separateur() {
	fmt.Println(Gris + strings.Repeat("─", largeurCadre+2) + Reset)
}

// sousTitre affiche un petit titre entre deux tirets
func sousTitre(texte string) {
	fmt.Println()
	fmt.Println(Jaune + "──── " + texte + " " + strings.Repeat("─", 48-len([]rune(texte))) + Reset)
}

// barreDeVie construit une barre du style ████████░░░░  35/40
func barreDeVie(actuel int, max int) string {
	taille := 20
	if max <= 0 {
		max = 1
	}
	if actuel < 0 {
		actuel = 0
	}
	plein := actuel * taille / max
	if plein > taille {
		plein = taille
	}
	vide := taille - plein

	// La couleur change selon les PV restants
	couleur := Vert
	pourcent := actuel * 100 / max
	if pourcent <= 50 {
		couleur = Jaune
	}
	if pourcent <= 25 {
		couleur = Rouge
	}

	barre := couleur + strings.Repeat("█", plein) + Gris + strings.Repeat("░", vide) + Reset
	return barre + " " + fmt.Sprintf("%d/%d", actuel, max)
}

// barreDeMana : la meme chose mais en bleu
func barreDeMana(actuel int, max int) string {
	taille := 20
	if max <= 0 {
		max = 1
	}
	if actuel < 0 {
		actuel = 0
	}
	plein := actuel * taille / max
	if plein > taille {
		plein = taille
	}
	vide := taille - plein
	return Bleu + strings.Repeat("█", plein) + Gris + strings.Repeat("░", vide) + Reset +
		" " + fmt.Sprintf("%d/%d", actuel, max)
}

// pause attend que le joueur appuie sur Entree
func pause() {
	fmt.Print(Gris + "\nAppuyez sur Entree pour continuer..." + Reset)
	lecteur.ReadString('\n')
}

// lireTexte pose une question et renvoie ce que le joueur a tape
func lireTexte(question string) string {
	fmt.Print(question)
	texte, _ := lecteur.ReadString('\n')
	texte = strings.TrimSpace(texte)
	return texte
}

// lireChoix demande un nombre entre min et max et recommence si c'est faux
func lireChoix(min int, max int) int {
	for {
		fmt.Print(Vert + "\nVotre choix > " + Reset)
		texte, finDuClavier := lecteur.ReadString('\n')
		texte = strings.TrimSpace(texte)

		// securite : si le clavier est ferme (Ctrl+D) on arrete le jeu
		// proprement au lieu de boucler dans le vide
		if finDuClavier != nil && texte == "" {
			fmt.Println("\nA bientot !")
			os.Exit(0)
		}

		nombre, erreur := strconv.Atoi(texte)
		if erreur != nil {
			fmt.Println(Rouge + "Ce n'est pas un nombre, essayez encore." + Reset)
			continue
		}
		if nombre < min || nombre > max {
			fmt.Println(Rouge + "Choix impossible, tapez un nombre entre " +
				strconv.Itoa(min) + " et " + strconv.Itoa(max) + "." + Reset)
			continue
		}
		return nombre
	}
}

// message affiche un texte d'information
func message(texte string) {
	fmt.Println(Cyan + "> " + texte + Reset)
}

// succes affiche un texte en vert
func succes(texte string) {
	fmt.Println(Vert + "✔ " + texte + Reset)
}

// erreur affiche un texte en rouge
func erreur(texte string) {
	fmt.Println(Rouge + "✖ " + texte + Reset)
}
