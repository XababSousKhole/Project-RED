package main

// ============================================================
//  character.go
//  TACHES 1, 2, 3, 8, 10, 11, 13, 16 + MISSION 2 et 4
//  La structure du personnage et tout ce qui le concerne.
// ============================================================

import (
	"fmt"
	"strings"
	"unicode"
)

// ------------------------------------------------------------
// TACHE 1 : la structure Character
// On y ajoute au fur et a mesure les attributs des taches
// suivantes : skills, money, equipement, mana, experience...
// ------------------------------------------------------------
type Character struct {
	Name         string    // son nom
	Class        string    // Humain / Elfe / Nain
	Level        int       // son niveau
	MaxHP        int       // points de vie maximum
	CurrentHP    int       // points de vie actuels
	Inventory    []string  // la liste de ses objets
	InventoryMax int       // TACHE 12 : taille max de l'inventaire
	Upgrades     int       // TACHE 18 : nombre d'agrandissements deja utilises
	Skills       []string  // TACHE 10 : la liste de ses sorts
	Money        int       // TACHE 13 : ses pieces d'or
	Equip        Equipment // TACHE 16 : son equipement porte
	Mana         int       // MISSION 4
	MaxMana      int       // MISSION 4
	XP           int       // MISSION 2 : experience actuelle
	XPMax        int       // MISSION 2 : experience a atteindre
	Initiative   int       // MISSION 1 : qui commence le combat
}

// ------------------------------------------------------------
// TACHE 2 : initCharacter
// Cette fonction cree un personnage et renvoie la structure
// remplie. C'est elle qui donne les valeurs de depart.
// ------------------------------------------------------------
func initCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int,
	inventaire []string, sorts []string) Character {

	var c Character

	c.Name = nom
	c.Class = classe
	c.Level = niveau
	c.MaxHP = pvMax
	c.CurrentHP = pvActuels
	c.Inventory = inventaire
	c.InventoryMax = 10 // TACHE 12 : 10 objets maximum au depart
	c.Upgrades = 0
	c.Skills = sorts
	c.Money = 100 // TACHE 13 : 100 pieces d'or au depart
	c.Equip = Equipment{Head: "", Body: "", Feet: ""}
	c.MaxMana = 50 // MISSION 4
	c.Mana = 50
	c.XP = 0 // MISSION 2
	c.XPMax = 100
	c.Initiative = initiativeDeLaClasse(classe) // MISSION 1

	return c
}

// MISSION 1 : chaque classe a sa propre initiative
func initiativeDeLaClasse(classe string) int {
	switch classe {
	case "Elfe":
		return 12 // les elfes sont rapides
	case "Humain":
		return 10
	case "Nain":
		return 8 // les nains sont lents mais costauds
	}
	return 10
}

// ------------------------------------------------------------
// TACHE 11 : characterCreation
// Le joueur choisit lui-meme son nom et sa classe.
// Cette fonction remplace l'initialisation en dur de la tache 2.
// ------------------------------------------------------------
func characterCreation() Character {
	nettoyerEcran()
	titre("CREATION DU PERSONNAGE")

	// ---- le nom (uniquement des lettres) ----
	nom := ""
	for nom == "" {
		saisie := lireTexte("\nQuel est votre nom, aventurier ? > ")

		if saisie == "" {
			erreur("Il faut bien ecrire quelque chose !")
			continue
		}
		if !queDesLettres(saisie) {
			erreur("Le nom ne doit contenir que des lettres.")
			continue
		}
		nom = formaterNom(saisie) // Alex, aLEX, ALEX -> Alex
	}

	// ---- la classe ----
	fmt.Println()
	fmt.Println("Choisissez votre classe :")
	fmt.Println()
	fmt.Println("  1. Humain  -> 100 PV  (equilibre)")
	fmt.Println("  2. Elfe    ->  80 PV  (rapide, plus d'initiative)")
	fmt.Println("  3. Nain    -> 120 PV  (resistant, lent)")

	choix := lireChoix(1, 3)

	classe := "Humain"
	pvMax := 100

	switch choix {
	case 1:
		classe = "Humain"
		pvMax = 100
	case 2:
		classe = "Elfe"
		pvMax = 80
	case 3:
		classe = "Nain"
		pvMax = 120
	}

	// Les PV de depart = 50% des PV max (consigne de la tache 11)
	pvDepart := pvMax / 2

	// Inventaire de depart : 3 potions de vie (tache 2)
	inventaireDepart := []string{"Potion de vie", "Potion de vie", "Potion de vie"}

	// Sort de base : Coup de poing (tache 10)
	sortsDepart := []string{"Coup de poing"}

	perso := initCharacter(nom, classe, 1, pvMax, pvDepart, inventaireDepart, sortsDepart)

	nettoyerEcran()
	titre("BIENVENUE " + strings.ToUpper(perso.Name))
	displayInfo(perso)
	pause()

	return perso
}

// queDesLettres verifie qu'il n'y a que des lettres dans le texte
func queDesLettres(texte string) bool {
	for _, lettre := range texte {
		if !unicode.IsLetter(lettre) {
			return false
		}
	}
	return true
}

// formaterNom : premiere lettre en majuscule, le reste en minuscule
func formaterNom(texte string) string {
	texte = strings.ToLower(texte)
	lettres := []rune(texte)
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// ------------------------------------------------------------
// TACHE 3 : displayInfo
// Affiche toutes les informations du personnage.
// ------------------------------------------------------------
func displayInfo(c Character) {
	fmt.Println()
	fmt.Println("  Nom        : " + Gras + c.Name + Reset)
	fmt.Println("  Classe     : " + c.Class)
	fmt.Println("  Niveau     : " + fmt.Sprint(c.Level))
	fmt.Println("  Or         : " + Jaune + fmt.Sprint(c.Money) + " pieces" + Reset)
	fmt.Println()
	fmt.Println("  PV         : " + barreDeVie(c.CurrentHP, c.MaxHP))
	fmt.Println("  Mana       : " + barreDeMana(c.Mana, c.MaxMana))
	fmt.Println("  Experience : " + fmt.Sprint(c.XP) + "/" + fmt.Sprint(c.XPMax))
	fmt.Println("  Initiative : " + fmt.Sprint(c.Initiative))
	fmt.Println()

	// Les sorts connus (tache 10)
	fmt.Println("  Sorts connus :")
	for i := 0; i < len(c.Skills); i++ {
		fmt.Println("    - " + c.Skills[i] + " (" + fmt.Sprint(degatsDuSort(c.Skills[i])) +
			" degats, " + fmt.Sprint(manaDuSort(c.Skills[i])) + " mana)")
	}

	// L'equipement porte (tache 16 et 17)
	fmt.Println()
	fmt.Println("  Equipement :")
	afficherEquipement(c)

	fmt.Println()
	fmt.Println("  Inventaire : " + fmt.Sprint(len(c.Inventory)) + "/" + fmt.Sprint(c.InventoryMax) + " objets")
}

// ------------------------------------------------------------
// TACHE 8 : isDead ("Wasted")
// Si le personnage tombe a 0 PV il meurt, puis il est
// ressuscite avec 50% de ses PV max.
// Renvoie true si le joueur est bien mort.
// ------------------------------------------------------------
func isDead(c *Character) bool {
	if c.CurrentHP > 0 {
		return false
	}

	c.CurrentHP = 0

	fmt.Println()
	fmt.Println(Rouge + "        ██  ██  ██  ██   WASTED   ██  ██  ██  ██" + Reset)
	fmt.Println()
	message(c.Name + " est tombe au combat...")

	// Resurrection a 50% des PV max
	c.CurrentHP = c.MaxHP / 2
	succes("Une lumiere chaude vous entoure : vous revenez a la vie !")
	fmt.Println("  PV : " + barreDeVie(c.CurrentHP, c.MaxHP))

	return true
}

// ------------------------------------------------------------
// MISSION 2 : experience et montee de niveau
// ------------------------------------------------------------
func gagnerExperience(c *Character, xp int) {
	fmt.Println()
	succes("Vous gagnez " + fmt.Sprint(xp) + " points d'experience !")
	c.XP = c.XP + xp

	// On utilise "for" et pas "if" : on peut monter plusieurs niveaux d'un coup
	for c.XP >= c.XPMax {
		c.XP = c.XP - c.XPMax // on garde l'experience en trop
		monterDeNiveau(c)
	}

	fmt.Println("  Experience : " + fmt.Sprint(c.XP) + "/" + fmt.Sprint(c.XPMax))
}

func monterDeNiveau(c *Character) {
	c.Level = c.Level + 1
	c.XPMax = c.XPMax + 50 // il faut de plus en plus d'XP a chaque niveau

	// Les gains de statistiques
	c.MaxHP = c.MaxHP + 20
	c.MaxMana = c.MaxMana + 10
	c.CurrentHP = c.MaxHP // on soigne completement le joueur
	c.Mana = c.MaxMana

	fmt.Println()
	fmt.Println(Jaune + "  ╔══════════════════════════════════════╗" + Reset)
	fmt.Println(Jaune + "  ║      NIVEAU SUPERIEUR ! Niveau " + fmt.Sprintf("%-5d", c.Level) + " ║" + Reset)
	fmt.Println(Jaune + "  ╚══════════════════════════════════════╝" + Reset)
	fmt.Println("   +20 PV max   +10 Mana max   PV et Mana restaures")
}
