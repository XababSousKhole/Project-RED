package main

// ============================================================
//  monster.go
//  TACHE 19 + BOSS FINAL
//  La structure Monster, le gobelin d'entrainement, le boss,
//  et tous les dessins ASCII.
// ============================================================

import "fmt"

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// >>> METTEZ ICI LE NOM DE VOTRE POTE (le boss final !)    <<<
const nomDuBoss = "LE POTE"

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// ------------------------------------------------------------
// TACHE 19 : la structure Monster
// ------------------------------------------------------------
type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int // MISSION 1
	XP         int // MISSION 2 : experience donnee au joueur
	EstBoss    bool
}

// ------------------------------------------------------------
// TACHE 19 : initGoblin
// Le gobelin d'entrainement : 40 PV, 5 d'attaque.
// ------------------------------------------------------------
func initGoblin() Monster {
	var gobelin Monster

	gobelin.Name = "Gobelin d'entrainement"
	gobelin.MaxHP = 40
	gobelin.CurrentHP = gobelin.MaxHP // PV actuels = PV max
	gobelin.Attack = 5
	gobelin.Initiative = 9 // MISSION 1
	gobelin.XP = 40        // MISSION 2
	gobelin.EstBoss = false

	return gobelin
}

// initBoss : le boss final, beaucoup plus costaud
func initBoss() Monster {
	var boss Monster

	boss.Name = nomDuBoss
	boss.MaxHP = 250
	boss.CurrentHP = boss.MaxHP
	boss.Attack = 12
	boss.Initiative = 15 // il joue presque toujours en premier
	boss.XP = 300
	boss.EstBoss = true

	return boss
}

// ============================================================
//  LES DESSINS ASCII
// ============================================================

// dessinGobelin : le petit gobelin vert
var dessinGobelin = []string{
	"          ,      ,",
	"         /(.-\"\"-.)\\",
	"     |\\  \\/      \\/  /|",
	"     | \\ / =.  .= \\ / |",
	"     \\( \\   o\\/o   / )/",
	"      \\_, '-/  \\-' ,_/",
	"        /   \\__/   \\",
	"        \\ \\__/\\__/ /",
	"      ___\\ \\|--|/ /___",
	"    /`    \\      /    `\\",
}

// ------------------------------------------------------------
// LE BOSS FINAL : sa tete change selon ses PV restants !
// C'est ici que vous pouvez coller l'ASCII art de votre pote.
// (Astuce : cherchez un convertisseur "photo vers ASCII art"
//  puis collez chaque ligne entre guillemets, comme ci-dessous)
// ------------------------------------------------------------

// PHASE 1 : plein de PV, il se moque de vous
var bossPhase1 = []string{
	"           .-\"\"\"\"\"\"\"\"-.",
	"         .'            '.",
	"        /   ___    ___   \\",
	"       |   (o o)  (o o)   |",
	"       |       ____       |",
	"       |      |    |      |",
	"        \\      \\__/      /",
	"         \\   \\______/   /",
	"          '.          .'",
	"            '-.____.-'",
}

// PHASE 2 : a moitie mort, il arrete de rigoler
var bossPhase2 = []string{
	"           .-\"\"\"\"\"\"\"\"-.",
	"         .'  //    \\\\  '.",
	"        /   ---    ---   \\",
	"       |   (> <)  (> <)   |",
	"       |       /\\/\\       |",
	"       |      |    |      |",
	"        \\      ----      /",
	"         \\   \\______/   /",
	"          '.          .'",
	"            '-.____.-'",
}

// PHASE 3 : dernier souffle, il ne rigole plus DU TOUT
var bossPhase3 = []string{
	"           .-\"\"\"\"\"\"\"\"-.",
	"         .'  \\\\ /\\ //  '.",
	"        /   \\__/  \\__/   \\",
	"       |    (x)    (x)    |",
	"       |       /\\/\\       |",
	"       |     /VVVVVV\\     |",
	"        \\    \\AAAAAA/    /",
	"         \\   \\______/   /",
	"          '.          .'",
	"            '-.____.-'",
}

// dessinerMonstre affiche le bon dessin selon le monstre et ses PV
func dessinerMonstre(m Monster) {
	dessin := dessinGobelin
	couleur := Vert

	if m.EstBoss {
		couleur = Violet
		pourcent := m.CurrentHP * 100 / m.MaxHP

		if pourcent > 60 {
			dessin = bossPhase1
		} else if pourcent > 20 {
			dessin = bossPhase2
		} else {
			dessin = bossPhase3
			couleur = Rouge
		}
	}

	fmt.Println()
	for i := 0; i < len(dessin); i++ {
		fmt.Println(couleur + dessin[i] + Reset)
	}
}

// phraseDuBoss : le boss parle selon ses PV restants (ambiance !)
func phraseDuBoss(m Monster) string {
	pourcent := m.CurrentHP * 100 / m.MaxHP

	if pourcent > 60 {
		return "\"Tu pensais vraiment pouvoir me battre ?\""
	}
	if pourcent > 20 {
		return "\"Ah... ca devient interessant.\""
	}
	return "\"OK. Maintenant... je ne rigole plus.\""
}

// dessinVictoire : l'ecran de victoire
func dessinVictoire() {
	fmt.Println()
	fmt.Println(Jaune + "        ★ ★ ★   V I C T O I R E   ★ ★ ★" + Reset)
	fmt.Println()
}

// dessinAttaque : un petit effet quand on tape fort
func dessinAttaque(texte string) {
	fmt.Println()
	fmt.Println(Jaune + "           *  *  *" + Reset)
	fmt.Println(Jaune + "        *   " + texte + "   *" + Reset)
	fmt.Println(Jaune + "           *  *  *" + Reset)
}
