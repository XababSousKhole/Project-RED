package main

// ============================================================
//  equipment.go
//  TACHES 15, 16, 17
//  Le forgeron, la fabrication et le port des equipements.
// ============================================================

import "fmt"

// ------------------------------------------------------------
// TACHE 16 : la structure Equipment
// Trois emplacements : la tete, le torse, les pieds.
// Si la case est vide (""), c'est qu'on ne porte rien.
// ------------------------------------------------------------
type Equipment struct {
	Head string // tete
	Body string // torse
	Feet string // pieds
}

// Une ressource necessaire a une fabrication
type Ressource struct {
	Nom      string
	Quantite int
}

// Une recette du forgeron
type Recette struct {
	Nom         string
	Emplacement string // "tete", "torse" ou "pieds"
	BonusPV     int    // TACHE 17 : PV max en plus quand c'est equipe
	Ressources  []Ressource
}

// TACHE 15 : le tableau de craft demande par le sujet
var recettesForgeron = []Recette{
	{
		Nom:         "Chapeau de l'aventurier",
		Emplacement: "tete",
		BonusPV:     10,
		Ressources: []Ressource{
			{"Plume de Corbeau", 1},
			{"Cuir de Sanglier", 1},
		},
	},
	{
		Nom:         "Tunique de l'aventurier",
		Emplacement: "torse",
		BonusPV:     25,
		Ressources: []Ressource{
			{"Fourrure de Loup", 2},
			{"Peau de Troll", 1},
		},
	},
	{
		Nom:         "Bottes de l'aventurier",
		Emplacement: "pieds",
		BonusPV:     15,
		Ressources: []Ressource{
			{"Fourrure de Loup", 1},
			{"Cuir de Sanglier", 1},
		},
	},
}

const prixFabrication = 5 // TACHE 15 : 5 pieces d'or par fabrication

// trouverRecette cherche une recette a partir du nom d'un equipement
func trouverRecette(nom string) (Recette, bool) {
	for i := 0; i < len(recettesForgeron); i++ {
		if recettesForgeron[i].Nom == nom {
			return recettesForgeron[i], true
		}
	}
	return Recette{}, false
}

// ------------------------------------------------------------
// TACHE 15 : le menu du forgeron
// ------------------------------------------------------------
func forgeron(c *Character) {
	for {
		nettoyerEcran()
		titre("LE FORGERON")

		fmt.Println()
		fmt.Println(Gris + "  \"Apporte-moi des materiaux et je te fais une armure !\"" + Reset)
		fmt.Println()
		fmt.Println("  Votre or : " + Jaune + fmt.Sprint(c.Money) + " pieces" + Reset +
			"   (chaque fabrication coute " + fmt.Sprint(prixFabrication) + " or)")
		fmt.Println()

		for i := 0; i < len(recettesForgeron); i++ {
			recette := recettesForgeron[i]

			// on ecrit la liste des ressources necessaires
			texteRessources := ""
			for j := 0; j < len(recette.Ressources); j++ {
				if j > 0 {
					texteRessources = texteRessources + " + "
				}
				texteRessources = texteRessources +
					fmt.Sprint(recette.Ressources[j].Quantite) + " " + recette.Ressources[j].Nom
			}

			couleur := Rouge
			if hasResources(c, recette) {
				couleur = Vert // en vert si on peut le fabriquer
			}

			fmt.Printf("  %d. %s%-26s%s (+%d PV max)\n", i+1, couleur, recette.Nom, Reset, recette.BonusPV)
			fmt.Println("     " + Gris + texteRessources + Reset)
		}

		fmt.Println()
		separateur()
		fmt.Println("  0. Retour")

		choix := lireChoix(0, len(recettesForgeron))
		if choix == 0 {
			return
		}

		craftEquipment(c, recettesForgeron[choix-1])
		pause()
	}
}

// ------------------------------------------------------------
// TACHE 15 : hasResources
// Est-ce que le joueur a toutes les ressources de la recette ?
// ------------------------------------------------------------
func hasResources(c *Character, recette Recette) bool {
	for i := 0; i < len(recette.Ressources); i++ {
		if compterObjet(c, recette.Ressources[i].Nom) < recette.Ressources[i].Quantite {
			return false
		}
	}
	return true
}

// ------------------------------------------------------------
// TACHE 15 : craftEquipment
// On verifie l'or, les ressources et la place, puis on fabrique.
// Les ressources sont supprimees de l'inventaire.
// ------------------------------------------------------------
func craftEquipment(c *Character, recette Recette) {
	fmt.Println()

	// 1) l'argent
	if c.Money < prixFabrication {
		erreur("Le forgeron veut " + fmt.Sprint(prixFabrication) +
			" pieces d'or et vous n'en avez que " + fmt.Sprint(c.Money) + ".")
		return
	}

	// 2) les ressources
	if !hasResources(c, recette) {
		erreur("Il vous manque des materiaux pour fabriquer " + recette.Nom + " :")
		for i := 0; i < len(recette.Ressources); i++ {
			manque := recette.Ressources[i].Quantite - compterObjet(c, recette.Ressources[i].Nom)
			if manque > 0 {
				fmt.Println("     - " + fmt.Sprint(manque) + " " + recette.Ressources[i].Nom)
			}
		}
		return
	}

	// 3) la place dans le sac
	//    (on va retirer les ressources donc la place se libere, mais
	//     on verifie quand meme au cas ou le sac serait plein)
	if !checkInventoryCapacity(c) {
		erreur("Votre sac est plein ! Videz-le avant de fabriquer.")
		return
	}

	// 4) on paye, on consomme, on fabrique
	c.Money = c.Money - prixFabrication

	for i := 0; i < len(recette.Ressources); i++ {
		for j := 0; j < recette.Ressources[i].Quantite; j++ {
			removeInventory(c, recette.Ressources[i].Nom)
		}
	}

	addInventory(c, recette.Nom)

	fmt.Println(Jaune + "   *CLANG*  *CLANG*  *CLANG*" + Reset)
	succes(recette.Nom + " fabrique ! (-" + fmt.Sprint(prixFabrication) + " or)")
	message("Equipez-le depuis votre inventaire pour gagner +" +
		fmt.Sprint(recette.BonusPV) + " PV max.")
}

// ------------------------------------------------------------
// TACHE 17 : equipItem
// On porte l'equipement : il sort de l'inventaire et donne
// des PV max en plus. S'il y avait deja quelque chose a cet
// emplacement, l'ancien retourne dans l'inventaire.
// ------------------------------------------------------------
func equipItem(c *Character, nomEquipement string) {
	recette, trouve := trouverRecette(nomEquipement)
	if !trouve {
		erreur("Cet objet ne peut pas etre equipe.")
		return
	}

	if !possede(c, nomEquipement) {
		erreur("Vous n'avez pas " + nomEquipement + " dans votre sac.")
		return
	}

	// on regarde ce qu'il y a deja a cet emplacement
	ancien := equipementPorte(c, recette.Emplacement)

	// le nouvel equipement quitte l'inventaire
	removeInventory(c, nomEquipement)

	// l'ancien revient dans l'inventaire et on enleve son bonus
	if ancien != "" {
		ancienneRecette, _ := trouverRecette(ancien)
		c.MaxHP = c.MaxHP - ancienneRecette.BonusPV
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		addInventory(c, ancien)
		message("Vous rangez " + ancien + " dans votre sac.")
	}

	// on place le nouveau au bon endroit
	placerEquipement(c, recette.Emplacement, nomEquipement)
	c.MaxHP = c.MaxHP + recette.BonusPV

	succes("Vous equipez " + nomEquipement + " ! (+" + fmt.Sprint(recette.BonusPV) + " PV max)")
	fmt.Println("  PV : " + barreDeVie(c.CurrentHP, c.MaxHP))
}

// desequiper : on enleve un equipement et il retourne dans le sac
func desequiper(c *Character, emplacement string) {
	nom := equipementPorte(c, emplacement)
	if nom == "" {
		erreur("Vous ne portez rien a cet emplacement.")
		return
	}

	if !checkInventoryCapacity(c) {
		erreur("Votre sac est plein, impossible d'y ranger " + nom + ".")
		return
	}

	recette, _ := trouverRecette(nom)
	c.MaxHP = c.MaxHP - recette.BonusPV
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	placerEquipement(c, emplacement, "")
	addInventory(c, nom)

	message("Vous retirez " + nom + " (-" + fmt.Sprint(recette.BonusPV) + " PV max)")
}

// equipementPorte renvoie ce que l'on porte a un emplacement
func equipementPorte(c *Character, emplacement string) string {
	switch emplacement {
	case "tete":
		return c.Equip.Head
	case "torse":
		return c.Equip.Body
	case "pieds":
		return c.Equip.Feet
	}
	return ""
}

// placerEquipement range un equipement au bon endroit de la structure
func placerEquipement(c *Character, emplacement string, nom string) {
	switch emplacement {
	case "tete":
		c.Equip.Head = nom
	case "torse":
		c.Equip.Body = nom
	case "pieds":
		c.Equip.Feet = nom
	}
}

// afficherEquipement montre les 3 emplacements
func afficherEquipement(c Character) {
	fmt.Println("    Tete  : " + texteEmplacement(c.Equip.Head))
	fmt.Println("    Torse : " + texteEmplacement(c.Equip.Body))
	fmt.Println("    Pieds : " + texteEmplacement(c.Equip.Feet))
}

func texteEmplacement(nom string) string {
	if nom == "" {
		return Gris + "(vide)" + Reset
	}
	recette, _ := trouverRecette(nom)
	return Vert + nom + Reset + " (+" + fmt.Sprint(recette.BonusPV) + " PV max)"
}

// menuEquipement : un petit menu pour retirer ce que l'on porte
func menuEquipement(c *Character) {
	for {
		nettoyerEcran()
		titre("EQUIPEMENT")

		fmt.Println()
		afficherEquipement(*c)
		fmt.Println()
		fmt.Println("  PV : " + barreDeVie(c.CurrentHP, c.MaxHP))
		fmt.Println()
		separateur()
		fmt.Println("  1. Retirer l'equipement de tete")
		fmt.Println("  2. Retirer l'equipement de torse")
		fmt.Println("  3. Retirer l'equipement de pieds")
		fmt.Println("  0. Retour")

		choix := lireChoix(0, 3)

		switch choix {
		case 0:
			return
		case 1:
			desequiper(c, "tete")
		case 2:
			desequiper(c, "torse")
		case 3:
			desequiper(c, "pieds")
		}
		pause()
	}
}
