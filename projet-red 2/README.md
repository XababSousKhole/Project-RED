# ⚔️ Legends of RED

Mini RPG **100 % terminal (CLI)** écrit en Go, réalisé dans le cadre du **Projet RED** (Ymmersion — Aix Ynov Campus).

Le joueur crée son personnage, gère son inventaire, achète chez le marchand, fabrique son équipement chez le forgeron, puis affronte un gobelin d'entraînement… et un **boss final** au tour par tour.

---

## 🚀 Installation & lancement

Il faut simplement avoir **Go** installé (version 1.21 ou plus) :

```bash
go version          # vérifier que Go est installé
```

Puis, depuis la racine du projet :

```bash
go run ./src
```

Ou pour construire un exécutable :

```bash
go build -o legends-of-red ./src
./legends-of-red
```

> 💡 Le jeu utilise des **couleurs ANSI** et des **emojis**. Utilisez un terminal moderne (Terminal macOS, Windows Terminal, iTerm2, le terminal de VS Code…) pour un rendu correct.

---

## 🗂️ Structure du projet

```
projet-red/
├── src/
│   ├── main.go        → écran titre + menu principal (tâche 6)
│   ├── ui.go          → affichage ASCII, couleurs, barres de vie, lecture clavier
│   ├── character.go   → structure Character, création, infos, mort, XP (tâches 1-3, 8, 11)
│   ├── inventory.go   → inventaire, potions, sorts, limite (tâches 4, 5, 9, 10, 12)
│   ├── shop.go        → marchand, prix, achat, agrandissement du sac (tâches 7, 13, 14, 18)
│   ├── equipment.go   → forgeron, craft, équipement, bonus PV (tâches 15, 16, 17)
│   ├── monster.go     → structure Monster, gobelin, boss, dessins ASCII (tâche 19)
│   └── combat.go      → combat tour par tour, IA du gobelin (tâches 20, 21, 22)
├── docs/
│   └── gestion-projet.md
├── go.mod
└── README.md
```

---

## 🎮 Contenu du jeu

### Personnage
- 3 classes : **Humain** (100 PV), **Elfe** (80 PV), **Nain** (120 PV)
- PV de départ = 50 % des PV max, niveau 1, sort « Coup de poing »
- Nom saisi par le joueur (lettres uniquement, reformaté en `Alex`)
- Mort → résurrection automatique à 50 % des PV max

### Inventaire
- 10 emplacements au départ, extensible 3 fois (+10 à chaque fois)
- Potion de vie (+50 PV), Potion de poison (-10 PV/s pendant 3 s), Potion de mana (+30)
- Livre de Sort : Boule de Feu (apprend le sort, une seule fois)

### Économie
| Objet | Prix |
|---|---|
| Potion de vie | 3 or |
| Potion de poison | 6 or |
| Potion de mana | 5 or |
| Livre de Sort : Boule de Feu | 25 or |
| Fourrure de Loup | 4 or |
| Peau de Troll | 7 or |
| Cuir de Sanglier | 3 or |
| Plume de Corbeau | 1 or |
| Augmentation d'inventaire | 30 or |

### Forgeron (5 or par fabrication)
| Équipement | Ressources | Bonus |
|---|---|---|
| Chapeau de l'aventurier | 1 Plume de Corbeau + 1 Cuir de Sanglier | +10 PV max |
| Tunique de l'aventurier | 2 Fourrures de Loup + 1 Peau de Troll | +25 PV max |
| Bottes de l'aventurier | 1 Fourrure de Loup + 1 Cuir de Sanglier | +15 PV max |

### Combat
- Tour par tour : **Attaquer** (5 dégâts) / **Inventaire** / **Sorts**
- Gobelin d'entraînement : 40 PV, 5 d'attaque, **200 % de dégâts tous les 3 tours**
- Le combat s'arrête dès que le joueur ou le monstre tombe à 0 PV

---

## ⭐ Missions bonus réalisées

| Mission | État | Où ? |
|---|---|---|
| M1 — Initiative | ✅ | `character.go` / `combat.go` — l'Elfe (12) commence avant le Gobelin (9) |
| M2 — Expérience & niveaux | ✅ | `character.go` — `gagnerExperience`, excédent d'XP reporté, +20 PV / +10 mana par niveau |
| M3 — Sorts en combat | ✅ | `combat.go` — Coup de poing (8) et Boule de Feu (18) |
| M4 — Mana | ✅ | Coup de poing 5 mana, Boule de Feu 15 mana, potion de mana chez le marchand |
| M5 — Améliorations | ✅ | Boss final ASCII à 3 phases, barres de vie colorées, potion de poison jetable sur l'ennemi, butin en or, menu d'équipement |

---

## 👑 Personnaliser le boss final

Le boss final, c'est **votre pote**. Deux choses à modifier dans `src/monster.go` :

1. **Son nom**, tout en haut du fichier :

```go
const nomDuBoss = "LE POTE"   // ← mettez son prénom ici
```

2. **Sa tête en ASCII**, dans les variables.

Pour transformer une vraie photo en ASCII : passez la photo dans un convertisseur « image → ASCII art » (largeur ≈ 40 caractères), puis collez chaque ligne entre guillemets.

⚠️ En Go, pensez à échapper les antislashs (`\\`) et les guillemets (`\"`) dans les chaînes.

---

## 🎨 Aperçu

```
╔══════════════════════════════════════════════════════╗
║                  COMBAT  -  TOUR 3                   ║
╚══════════════════════════════════════════════════════╝

          ,      ,
         /(.-""-.)\
     |\  \/      \/  /|
     | \ / =.  .= \ / |
     \( \   o\/o   / )/
      \_, '-/  \-' ,_/
        /   \__/   \
        \ \__/\__/ /
      ___\ \|--|/ /___
    /`    \      /    `\

  Gobelin d'entrainement
  PV : ██████████████░░░░░░ 28/40
──────────────────────────────────────────────────────
  Alex  (niveau 1)
  PV   : ████████████░░░░░░░░ 60/100
  Mana : ██████████████░░░░░░ 35/50
──────────────────────────────────────────────────────
  Que voulez-vous faire ?

  1. Attaquer      (attaque basique, 5 degats)
  2. Inventaire    (utiliser un objet)
  3. Sorts         (MISSION 3 : magie)
```

---

## 👥 Équipe

| Pôle | Fichiers |
|---|---|
| 👤 Personnage + inventaire | `character.go`, `inventory.go` |
| 💰 Économie + craft | `shop.go`, `equipment.go` |
| ⚔️ Combat | `combat.go`, `monster.go` |
| 🎨 Commun | `main.go`, `ui.go` |

Voir `docs/gestion-projet.md` pour la répartition détaillée et l'organisation Git.
