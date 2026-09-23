package main

// ============================================================
//  main.go
//  TACHE 6 : le menu principal du jeu (avec un switch case)
//  C'est le fichier qui relie toutes les parties du projet.
// ============================================================

import "fmt"

func main() {
	// ---- l'ecran titre ----
	if !ecranTitre() {
		return // le joueur a choisi "Quitter"
	}

	// ---- TACHE 11 : le joueur cree son personnage ----
	// (cette fonction remplace l'initialisation de la tache 2)
	joueur := characterCreation()

	// ---- TACHE 6 : la boucle du menu principal ----
	menuPrincipal(&joueur)
}

// ------------------------------------------------------------
// L'ecran d'accueil du jeu
// ------------------------------------------------------------
func ecranTitre() bool {
	nettoyerEcran()

	fmt.Println()
	fmt.Println(Rouge + "   ██       ███████   ██████   ███████  ███    ██  ██████  " + Reset)
	fmt.Println(Rouge + "   ██       ██       ██        ██       ████   ██  ██   ██ " + Reset)
	fmt.Println(Rouge + "   ██       █████    ██   ███  █████    ██ ██  ██  ██   ██ " + Reset)
	fmt.Println(Rouge + "   ██       ██       ██    ██  ██       ██  ██ ██  ██   ██ " + Reset)
	fmt.Println(Rouge + "   ███████  ███████   ██████   ███████  ██   ████  ██████  " + Reset)
	fmt.Println()
	fmt.Println(Gras + centrer("O F   R E D", largeurCadre) + Reset)
	fmt.Println()
	titre("Bienvenue, aventurier !")
	fmt.Println()
	fmt.Println("  1. Nouvelle partie")
	fmt.Println("  2. Quitter")

	choix := lireChoix(1, 2)
	return choix == 1
}

// ------------------------------------------------------------
// TACHE 6 : le menu principal
// Chaque choix appelle une fonction ecrite dans un autre fichier.
// ------------------------------------------------------------
func menuPrincipal(c *Character) {
	for {
		nettoyerEcran()
		titre("MENU PRINCIPAL")

		// Un petit rappel de l'etat du personnage en haut du menu
		fmt.Println()
		fmt.Println("  " + Gras + c.Name + Reset + "  -  " + c.Class + "  -  niveau " + fmt.Sprint(c.Level))
		fmt.Println("  PV   : " + barreDeVie(c.CurrentHP, c.MaxHP))
		fmt.Println("  Mana : " + barreDeMana(c.Mana, c.MaxMana))
		fmt.Println("  Or   : " + Jaune + fmt.Sprint(c.Money) + " pieces" + Reset +
			"      Sac : " + fmt.Sprint(len(c.Inventory)) + "/" + fmt.Sprint(c.InventoryMax))
		fmt.Println()
		separateur()

		fmt.Println("  1. 👤  Voir mon personnage")
		fmt.Println("  2. 🎒  Inventaire")
		fmt.Println("  3. 🛡   Equipement")
		fmt.Println("  4. 🧙  Marchand")
		fmt.Println("  5. 🔨  Forgeron")
		fmt.Println("  6. ⚔   Entrainement (Gobelin)")
		fmt.Println("  7. 👑  Boss final (" + nomDuBoss + ")")
		fmt.Println("  8. ❔  Qui sont-ils ?")
		fmt.Println("  9. 🚪  Quitter")

		choix := lireChoix(1, 9)

		// Le switch case demande par la TACHE 6
		switch choix {

		case 1: // TACHE 3
			nettoyerEcran()
			titre("FICHE DE PERSONNAGE")
			displayInfo(*c)
			pause()

		case 2: // TACHE 4
			accessInventory(c)

		case 3: // TACHE 17
			menuEquipement(c)

		case 4: // TACHE 7 et 14
			marchand(c)

		case 5: // TACHE 15
			forgeron(c)

		case 6: // TACHE 22
			trainingFight(c)

		case 7: // BOSS FINAL (amelioration perso)
			bossFight(c)

		case 8: // MISSION 6
			quiSontIls()

		case 9: // TACHE 6 : quitter le programme
			nettoyerEcran()
			titre("A BIENTOT, " + c.Name)
			fmt.Println()
			message("Merci d'avoir joue a Legends of RED !")
			fmt.Println()
			return
		}
	}
}

// ------------------------------------------------------------
// MISSION 6 : "Qui sont-ils ?"
// Deux artistes sont caches dans les titres des taches
// des parties 2 et 3 du sujet.
// ------------------------------------------------------------
func quiSontIls() {
	nettoyerEcran()
	titre("QUI SONT-ILS ?")

	fmt.Println()
	fmt.Println("  " + Gras + "Partie 2 (l'economie) : le groupe ABBA" + Reset)
	fmt.Println(Gris + "    - Tache 13 : Money, Money, Money       (1976)" + Reset)
	fmt.Println(Gris + "    - Tache 14 : Two for the Price of One  (1981)" + Reset)
	fmt.Println(Gris + "    - Tache 15 : Gimme! Gimme! Gimme!      (1979)" + Reset)
	fmt.Println(Gris + "    - Tache 17 : Mamma Mia                 (1975)" + Reset)
	fmt.Println(Gris + "    - Tache 18 : On and On and On          (1980)" + Reset)

	fmt.Println()
	fmt.Println("  " + Gras + "Partie 3 (le combat) : le realisateur Steven Spielberg" + Reset)
	fmt.Println(Gris + "    - Tache 19 : La Chose  (Firelight / ses debuts)" + Reset)
	fmt.Println(Gris + "    - Tache 20 : A.I. Intelligence Artificielle (2001)" + Reset)
	fmt.Println(Gris + "    - Tache 21 : Ready Player One               (2018)" + Reset)
	fmt.Println(Gris + "    - Tache 22 : Fighter Squad (1961) / Duel (1971)" + Reset)

	pause()
}
