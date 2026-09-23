package main

// ============================================================
//  shop.go
//  TACHES 7, 9, 10, 13, 14, 18
//  Le marchand : il vend des potions, des livres de sort,
//  des ressources et l'augmentation d'inventaire.
// ============================================================

import "fmt"

// ------------------------------------------------------------
// Un article du marchand : un nom et un prix.
// On met tout dans une liste, comme ca ajouter un objet au
// magasin = ajouter une ligne ici (et rien d'autre a changer).
// ------------------------------------------------------------
type Article struct {
	Nom  string
	Prix int
}

// TACHE 14 : les prix demandes par le sujet
var catalogueMarchand = []Article{
	{"Potion de vie", 3},
	{"Potion de poison", 6},
	{"Potion de mana", 5}, // MISSION 4
	{"Livre de Sort : Boule de Feu", 25},
	{"Fourrure de Loup", 4},
	{"Peau de Troll", 7},
	{"Cuir de Sanglier", 3},
	{"Plume de Corbeau", 1},
	{"Augmentation d'inventaire", 30}, // TACHE 18
}

// ------------------------------------------------------------
// TACHE 7 + 14 : le menu du marchand
// ------------------------------------------------------------
func marchand(c *Character) {
	for {
		nettoyerEcran()
		titre("LE MARCHAND")

		fmt.Println()
		fmt.Println(Gris + "  \"Bienvenue l'ami ! Tout est frais, tout est bon marche.\"" + Reset)
		fmt.Println()
		fmt.Println("  Votre or : " + Jaune + fmt.Sprint(c.Money) + " pieces" + Reset +
			"     Sac : " + fmt.Sprint(len(c.Inventory)) + "/" + fmt.Sprint(c.InventoryMax))
		fmt.Println()

		for i := 0; i < len(catalogueMarchand); i++ {
			fmt.Printf("  %d. %-32s %s%3d or%s\n",
				i+1, catalogueMarchand[i].Nom, Jaune, catalogueMarchand[i].Prix, Reset)
		}

		fmt.Println()
		separateur()
		fmt.Println("  0. Retour")

		choix := lireChoix(0, len(catalogueMarchand))
		if choix == 0 {
			return
		}

		buyItem(c, catalogueMarchand[choix-1])
		pause()
	}
}

// ------------------------------------------------------------
// TACHE 14 : buyItem
// On verifie l'argent, puis la place, puis on achete.
// ------------------------------------------------------------
func buyItem(c *Character, article Article) {
	fmt.Println()

	// 1) est-ce qu'on a assez d'argent ?
	if c.Money < article.Prix {
		erreur("Pas assez d'or ! Il vous manque " +
			fmt.Sprint(article.Prix-c.Money) + " pieces.")
		return
	}

	// 2) cas special : l'augmentation d'inventaire n'est pas un objet
	//    du sac, elle agrandit le sac directement (TACHE 18)
	if article.Nom == "Augmentation d'inventaire" {
		if !upgradeInventorySlot(c) {
			return // le sac est deja au maximum : on ne paye pas
		}
		c.Money = c.Money - article.Prix
		fmt.Println("  Or restant : " + Jaune + fmt.Sprint(c.Money) + Reset)
		return
	}

	// 3) est-ce qu'il reste de la place dans le sac ? (TACHE 12)
	if !checkInventoryCapacity(c) {
		erreur("Votre sac est plein, le marchand refuse de vous vendre quoi que ce soit.")
		return
	}

	// 4) achat
	c.Money = c.Money - article.Prix
	addInventory(c, article.Nom)

	succes("Vous achetez : " + article.Nom + " (-" + fmt.Sprint(article.Prix) + " or)")
	fmt.Println("  Or restant : " + Jaune + fmt.Sprint(c.Money) + Reset)
}

// ------------------------------------------------------------
// TACHE 18 : upgradeInventorySlot
// +10 places dans le sac, 3 fois maximum.
// ------------------------------------------------------------
func upgradeInventorySlot(c *Character) bool {
	if c.Upgrades >= 3 {
		erreur("Vous avez deja agrandi votre sac 3 fois, c'est le maximum !")
		return false
	}

	c.Upgrades = c.Upgrades + 1
	c.InventoryMax = c.InventoryMax + 10

	succes("Votre sac s'agrandit ! Capacite : " + fmt.Sprint(c.InventoryMax) +
		" objets (" + fmt.Sprint(c.Upgrades) + "/3 ameliorations)")
	return true
}
