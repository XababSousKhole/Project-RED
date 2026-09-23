package main

// ============================================================
//  combat.go
//  TACHES 20, 21, 22 + MISSIONS 1, 2, 3, 4
//  Le combat tour par tour entre le joueur et un monstre.
// ============================================================

import (
	"fmt"
	"time"
)

// ------------------------------------------------------------
// TACHE 22 : trainingFight
// Lance le combat d'entrainement contre le gobelin.
// ------------------------------------------------------------
func trainingFight(c *Character) {
	gobelin := initGoblin()

	nettoyerEcran()
	titre("ENTRAINEMENT")
	dessinerMonstre(gobelin)
	fmt.Println()
	message("Un " + gobelin.Name + " surgit des buissons !")
	pause()

	lancerCombat(c, &gobelin)
}

// bossFight : le combat contre le boss final (amelioration perso)
func bossFight(c *Character) {
	boss := initBoss()

	nettoyerEcran()
	fmt.Println()
	fmt.Println(Rouge + centrer("!!!  ATTENTION, AVENTURIER  !!!", largeurCadre) + Reset)
	fmt.Println()
	fmt.Println(centrer("UN ENNEMI TRES PUISSANT APPROCHE", largeurCadre))
	dessinerMonstre(boss)
	fmt.Println()
	fmt.Println(centrer(Violet+phraseDuBoss(boss)+Reset, largeurCadre))
	fmt.Println()
	fmt.Println(centrer(Gras+"* "+boss.Name+" *", largeurCadre))
	fmt.Println(centrer("BOSS FINAL - NIVEAU ???", largeurCadre))
	fmt.Println()
	pause()

	lancerCombat(c, &boss)
}

// ------------------------------------------------------------
// TACHE 22 : lancerCombat
// La boucle du combat. Le joueur et le monstre jouent
// l'un apres l'autre jusqu'a ce que l'un tombe a 0 PV.
// ------------------------------------------------------------
func lancerCombat(c *Character, m *Monster) {
	tour := 1 // la variable qui dit a quel tour de combat on est

	// MISSION 1 : celui qui a le plus d'initiative commence
	joueurCommence := c.Initiative >= m.Initiative

	nettoyerEcran()
	if joueurCommence {
		message("Votre initiative (" + fmt.Sprint(c.Initiative) + ") depasse celle de " +
			m.Name + " (" + fmt.Sprint(m.Initiative) + ") : vous commencez !")
	} else {
		message(m.Name + " est plus rapide que vous (" + fmt.Sprint(m.Initiative) +
			" contre " + fmt.Sprint(c.Initiative) + ") : il commence !")
	}
	pause()

	for {
		// ----- le joueur joue en premier ou en second -----
		if joueurCommence {
			characterTurn(c, m, tour)
			if combatTermine(c, m) {
				break
			}
			goblinPattern(m, c, tour)
		} else {
			goblinPattern(m, c, tour)
			if combatTermine(c, m) {
				break
			}
			characterTurn(c, m, tour)
		}

		if combatTermine(c, m) {
			break
		}

		tour = tour + 1
	}

	finDuCombat(c, m)
}

// combatTermine : est-ce que quelqu'un est tombe a 0 PV ?
func combatTermine(c *Character, m *Monster) bool {
	return c.CurrentHP <= 0 || m.CurrentHP <= 0
}

// ------------------------------------------------------------
// L'ecran de combat : les deux barres de vie + le dessin
// ------------------------------------------------------------
func afficherEcranCombat(c *Character, m *Monster, tour int) {
	nettoyerEcran()
	titre("COMBAT  -  TOUR " + fmt.Sprint(tour))

	dessinerMonstre(*m)

	fmt.Println()
	fmt.Println("  " + Gras + m.Name + Reset)
	fmt.Println("  PV : " + barreDeVie(m.CurrentHP, m.MaxHP))

	if m.EstBoss {
		fmt.Println()
		fmt.Println("  " + Violet + phraseDuBoss(*m) + Reset)
	}

	separateur()
	fmt.Println("  " + Gras + c.Name + Reset + "  (niveau " + fmt.Sprint(c.Level) + ")")
	fmt.Println("  PV   : " + barreDeVie(c.CurrentHP, c.MaxHP))
	fmt.Println("  Mana : " + barreDeMana(c.Mana, c.MaxMana))
	separateur()
}

// ------------------------------------------------------------
// TACHE 21 : characterTurn
// Le tour du joueur : attaquer, ouvrir l'inventaire ou
// lancer un sort (MISSION 3).
// ------------------------------------------------------------
func characterTurn(c *Character, m *Monster, tour int) {
	tourJoue := false

	for !tourJoue {
		afficherEcranCombat(c, m, tour)

		fmt.Println("  Que voulez-vous faire ?")
		fmt.Println()
		fmt.Println("  1. Attaquer      (attaque basique, 5 degats)")
		fmt.Println("  2. Inventaire    (utiliser un objet)")
		fmt.Println("  3. Sorts         (MISSION 3 : magie)")

		choix := lireChoix(1, 3)

		switch choix {

		case 1:
			// ---- attaque basique : 5 degats (tache 21) ----
			degats := 5
			m.CurrentHP = m.CurrentHP - degats
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}

			fmt.Println()
			message("Vous utilisez : Attaque basique")
			fmt.Println(Jaune + "  " + c.Name + " inflige " + fmt.Sprint(degats) +
				" degats a " + m.Name + " !" + Reset)
			fmt.Println("  " + m.Name + " : " + barreDeVie(m.CurrentHP, m.MaxHP))
			tourJoue = true

		case 2:
			// ---- l'inventaire pendant le combat (tache 21) ----
			tourJoue = inventaireDeCombat(c, m)

		case 3:
			// ---- les sorts (MISSION 3 et 4) ----
			tourJoue = lancerUnSort(c, m)
		}

		if tourJoue {
			pause()
		}
	}
}

// ------------------------------------------------------------
// TACHE 21 : l'inventaire pendant le combat.
// Renvoie true si le joueur a bien utilise un objet
// (dans ce cas son tour est termine).
// ------------------------------------------------------------
func inventaireDeCombat(c *Character, m *Monster) bool {
	nettoyerEcran()
	titre("INVENTAIRE - COMBAT")

	objets := afficherListeInventaire(c)

	fmt.Println()
	separateur()
	fmt.Println("  0. Retour (ne pas utiliser d'objet)")

	choix := lireChoix(0, len(objets))
	if choix == 0 {
		return false // on n'a pas joue, on revient au menu de combat
	}

	objet := objets[choix-1]

	// AMELIORATION (MISSION 5) : en plein combat, la potion de
	// poison est jetee sur l'ennemi au lieu d'etre bue !
	if objet == "Potion de poison" {
		jeterPoison(c, m)
		return true
	}

	// Les equipements ne s'equipent pas en plein combat
	if objet == "Chapeau de l'aventurier" || objet == "Tunique de l'aventurier" ||
		objet == "Bottes de l'aventurier" {
		erreur("Pas le temps de s'habiller en plein combat !")
		pause()
		return false
	}

	utiliserObjet(c, objet)
	return true
}

// jeterPoison : la potion de poison lancee sur le monstre
func jeterPoison(c *Character, m *Monster) {
	removeInventory(c, "Potion de poison")

	fmt.Println()
	message("Vous jetez la Potion de poison sur " + m.Name + " !")

	for seconde := 1; seconde <= 3; seconde++ {
		time.Sleep(1 * time.Second)

		m.CurrentHP = m.CurrentHP - 10
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		fmt.Println(Violet + "  Le poison ronge " + m.Name + " ! -10 PV" + Reset)
		fmt.Println("  " + m.Name + " : " + barreDeVie(m.CurrentHP, m.MaxHP))
	}
}

// ------------------------------------------------------------
// MISSION 3 + 4 : lancer un sort (qui consomme du mana)
// ------------------------------------------------------------
func lancerUnSort(c *Character, m *Monster) bool {
	nettoyerEcran()
	titre("LIVRE DE SORTS")

	fmt.Println()
	fmt.Println("  Mana : " + barreDeMana(c.Mana, c.MaxMana))
	fmt.Println()

	for i := 0; i < len(c.Skills); i++ {
		sort := c.Skills[i]
		couleur := Vert
		if c.Mana < manaDuSort(sort) {
			couleur = Rouge // en rouge si on n'a pas assez de mana
		}
		fmt.Printf("  %d. %s%-20s%s %d degats  -  %d mana\n",
			i+1, couleur, sort, Reset, degatsDuSort(sort), manaDuSort(sort))
	}

	fmt.Println()
	separateur()
	fmt.Println("  0. Retour")

	choix := lireChoix(0, len(c.Skills))
	if choix == 0 {
		return false
	}

	sort := c.Skills[choix-1]
	cout := manaDuSort(sort)

	// MISSION 4 : pas assez de mana = sort impossible
	if c.Mana < cout {
		fmt.Println()
		erreur("Pas assez de mana pour lancer " + sort + " ! (" +
			fmt.Sprint(c.Mana) + "/" + fmt.Sprint(cout) + ")")
		pause()
		return false
	}

	c.Mana = c.Mana - cout
	degats := degatsDuSort(sort)

	m.CurrentHP = m.CurrentHP - degats
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}

	dessinAttaque(sort)
	fmt.Println(Jaune + "  " + c.Name + " inflige " + fmt.Sprint(degats) +
		" degats a " + m.Name + " !" + Reset)
	fmt.Println("  " + m.Name + " : " + barreDeVie(m.CurrentHP, m.MaxHP))
	fmt.Println("  Mana : " + barreDeMana(c.Mana, c.MaxMana))

	return true
}

// ------------------------------------------------------------
// TACHE 20 : goblinPattern
// Le tour du monstre : il tape 100% de son attaque, et tous
// les 3 tours il tape 200% (attaque speciale).
// ------------------------------------------------------------
func goblinPattern(m *Monster, c *Character, tour int) {
	fmt.Println()
	fmt.Println(Gris + "  ... Tour de " + m.Name + " ..." + Reset)
	time.Sleep(1 * time.Second)

	degats := m.Attack // 100% de son attaque

	// tous les 3 tours : 200% des degats
	if tour%3 == 0 {
		degats = m.Attack * 2
		fmt.Println(Rouge + "  /!\\ " + m.Name + " prepare une attaque puissante !" + Reset)
		time.Sleep(1 * time.Second)
	}

	c.CurrentHP = c.CurrentHP - degats
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	fmt.Println(Rouge + "  " + m.Name + " inflige " + fmt.Sprint(degats) +
		" degats a " + c.Name + " !" + Reset)
	fmt.Println("  " + c.Name + " : " + barreDeVie(c.CurrentHP, c.MaxHP))

	pause()
}

// ------------------------------------------------------------
// La fin du combat : victoire ou defaite
// ------------------------------------------------------------
func finDuCombat(c *Character, m *Monster) {
	nettoyerEcran()

	// ---- le monstre est mort : victoire ----
	if m.CurrentHP <= 0 {
		titre("COMBAT TERMINE")
		dessinVictoire()
		succes("Vous avez vaincu " + m.Name + " !")

		// MISSION 2 : le joueur gagne de l'experience
		gagnerExperience(c, m.XP)

		// une petite recompense en or (amelioration perso)
		or := 20
		if m.EstBoss {
			or = 200
		}
		c.Money = c.Money + or
		fmt.Println()
		succes("Butin : +" + fmt.Sprint(or) + " pieces d'or")

		if m.EstBoss {
			fmt.Println()
			fmt.Println(Jaune + centrer("LE BOSS FINAL EST VAINCU !", largeurCadre) + Reset)
			fmt.Println(centrer("Vous etes officiellement une legende.", largeurCadre))
		}

		pause()
		return
	}

	// ---- le joueur est mort ----
	titre("COMBAT TERMINE")
	isDead(c) // TACHE 8 : mort puis resurrection a 50% des PV max
	fmt.Println()
	message("Vous etes ramene au menu principal pour vous reposer.")
	pause()
}
