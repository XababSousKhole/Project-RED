package main

// ============================================================
//  inventory.go
//  TACHES 4, 5, 7, 9, 10, 12
//  L'inventaire du personnage : ajouter, retirer, utiliser.
// ============================================================

import (
	"fmt"
	"time"
)

// ------------------------------------------------------------
// TACHE 12 : checkInventoryCapacity
// Renvoie true s'il reste de la place dans le sac.
// ------------------------------------------------------------
func checkInventoryCapacity(c *Character) bool {
	return len(c.Inventory) < c.InventoryMax
}

// ------------------------------------------------------------
// TACHE 7 : addInventory
// Ajoute un objet dans l'inventaire (si il reste de la place).
// ------------------------------------------------------------
func addInventory(c *Character, objet string) bool {
	if !checkInventoryCapacity(c) {
		erreur("Votre sac est plein ! (" + fmt.Sprint(c.InventoryMax) + " objets maximum)")
		return false
	}
	c.Inventory = append(c.Inventory, objet)
	return true
}

// ------------------------------------------------------------
// TACHE 7 : removeInventory
// Retire UN exemplaire de l'objet demande.
// ------------------------------------------------------------
func removeInventory(c *Character, objet string) bool {
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i] == objet {
			// on recolle le debut et la fin de la liste sans la case i
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

// compterObjet compte combien de fois on possede un objet
func compterObjet(c *Character, objet string) int {
	total := 0
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i] == objet {
			total = total + 1
		}
	}
	return total
}

// possede renvoie true si le joueur a au moins un exemplaire de l'objet
func possede(c *Character, objet string) bool {
	return compterObjet(c, objet) > 0
}

// listeObjetsUniques renvoie la liste des objets SANS doublon
// (pour afficher "Potion de vie x3" au lieu de 3 lignes)
func listeObjetsUniques(c *Character) []string {
	liste := []string{}
	for i := 0; i < len(c.Inventory); i++ {
		dejaVu := false
		for j := 0; j < len(liste); j++ {
			if liste[j] == c.Inventory[i] {
				dejaVu = true
			}
		}
		if !dejaVu {
			liste = append(liste, c.Inventory[i])
		}
	}
	return liste
}

// afficherListeInventaire affiche l'inventaire numerote
func afficherListeInventaire(c *Character) []string {
	objets := listeObjetsUniques(c)

	fmt.Println()
	fmt.Println("  Sac : " + fmt.Sprint(len(c.Inventory)) + "/" + fmt.Sprint(c.InventoryMax) + " objets")
	fmt.Println()

	if len(objets) == 0 {
		fmt.Println(Gris + "  (votre sac est vide)" + Reset)
		return objets
	}

	for i := 0; i < len(objets); i++ {
		nombre := compterObjet(c, objets[i])
		fmt.Printf("  %d. %-32s x%d\n", i+1, objets[i], nombre)
	}
	return objets
}

// ------------------------------------------------------------
// TACHE 4 : accessInventory
// Le menu de l'inventaire hors combat.
// ------------------------------------------------------------
func accessInventory(c *Character) {
	for {
		nettoyerEcran()
		titre("INVENTAIRE")

		objets := afficherListeInventaire(c)

		fmt.Println()
		separateur()
		fmt.Println("  Tapez le numero d'un objet pour l'utiliser")
		fmt.Println("  0. Retour")

		choix := lireChoix(0, len(objets))
		if choix == 0 {
			return
		}

		utiliserObjet(c, objets[choix-1])
		pause()
	}
}

// ------------------------------------------------------------
// utiliserObjet : applique l'effet de l'objet choisi.
// C'est le "cerveau" de l'inventaire : chaque objet a son effet.
// ------------------------------------------------------------
func utiliserObjet(c *Character, objet string) {
	fmt.Println()
	switch objet {

	case "Potion de vie":
		takePot(c)

	case "Potion de poison":
		poisonPot(c)

	case "Potion de mana":
		potionDeMana(c)

	case "Livre de Sort : Boule de Feu":
		spellBook(c, "Boule de Feu")

	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
		equipItem(c, objet)

	default:
		erreur(objet + " ne s'utilise pas : c'est une ressource pour le Forgeron.")
	}
}

// ------------------------------------------------------------
// TACHE 5 : takePot (potion de vie)
// Rend 50 PV, se consomme, et on ne depasse jamais les PV max.
// ------------------------------------------------------------
func takePot(c *Character) {
	if !possede(c, "Potion de vie") {
		erreur("Vous n'avez plus de Potion de vie.")
		return
	}

	removeInventory(c, "Potion de vie")

	avant := c.CurrentHP
	c.CurrentHP = c.CurrentHP + 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP // on ne depasse jamais le maximum
	}

	message("Vous buvez une Potion de vie.")
	succes("+" + fmt.Sprint(c.CurrentHP-avant) + " PV")
	fmt.Println("  PV : " + barreDeVie(c.CurrentHP, c.MaxHP))
}

// ------------------------------------------------------------
// TACHE 9 : poisonPot (potion de poison)
// Inflige 10 degats par seconde pendant 3 secondes.
// ------------------------------------------------------------
func poisonPot(c *Character) {
	if !possede(c, "Potion de poison") {
		erreur("Vous n'avez plus de Potion de poison.")
		return
	}

	removeInventory(c, "Potion de poison")

	message("Vous buvez la Potion de poison... mauvaise idee.")

	for tour := 1; tour <= 3; tour++ {
		time.Sleep(1 * time.Second) // la bibliotheque "time" pour attendre

		c.CurrentHP = c.CurrentHP - 10
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}

		fmt.Println(Violet + "  Le poison agit ! -10 PV" + Reset)
		fmt.Println("  PV : " + barreDeVie(c.CurrentHP, c.MaxHP))
	}

	isDead(c) // TACHE 8 : si on tombe a 0 PV
}

// MISSION 4 : potion de mana (meme principe que la potion de vie)
func potionDeMana(c *Character) {
	if !possede(c, "Potion de mana") {
		erreur("Vous n'avez plus de Potion de mana.")
		return
	}

	removeInventory(c, "Potion de mana")

	avant := c.Mana
	c.Mana = c.Mana + 30
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}

	message("Vous buvez une Potion de mana.")
	succes("+" + fmt.Sprint(c.Mana-avant) + " Mana")
	fmt.Println("  Mana : " + barreDeMana(c.Mana, c.MaxMana))
}

// ------------------------------------------------------------
// TACHE 10 : spellBook
// Apprend un nouveau sort. Un sort ne s'apprend qu'une fois.
// ------------------------------------------------------------
func spellBook(c *Character, sort string) {
	// est-ce qu'on connait deja ce sort ?
	for i := 0; i < len(c.Skills); i++ {
		if c.Skills[i] == sort {
			erreur("Vous connaissez deja le sort " + sort + " ! Le livre reste dans votre sac.")
			return
		}
	}

	removeInventory(c, "Livre de Sort : "+sort)
	c.Skills = append(c.Skills, sort)

	fmt.Println(Violet + "  ~*~ Le livre s'ouvre tout seul et les pages s'envolent ~*~" + Reset)
	succes("Nouveau sort appris : " + sort + " !")
}

// degatsDuSort : combien de degats fait chaque sort (MISSION 3)
func degatsDuSort(sort string) int {
	switch sort {
	case "Coup de poing":
		return 8
	case "Boule de Feu":
		return 18
	}
	return 5
}

// manaDuSort : combien de mana coute chaque sort (MISSION 4)
func manaDuSort(sort string) int {
	switch sort {
	case "Coup de poing":
		return 5
	case "Boule de Feu":
		return 15
	}
	return 0
}
