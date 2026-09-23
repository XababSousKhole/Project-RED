# Gestion de projet — Legends of RED

## 1. Équipe et répartition

Le projet est découpé en **3 pôles** (et pas en « tâches 1 à 10 / 11 à 20 / 21 à 30 »),
parce que les fonctionnalités sont liées entre elles.

| Personne | Pôle | Fichiers | Tâches |
|---|---|---|---|
| 👤 Personne 1 | Personnage + inventaire | `character.go`, `inventory.go` | 1 → 12 |
| 💰 Personne 2 | Économie + craft + équipement | `shop.go`, `equipment.go` | 13 → 18 |
| ⚔️ Personne 3 | Combat + monstres | `combat.go`, `monster.go` | 19 → 22 |
| 🤝 Tous | Menus, affichage, intégration | `main.go`, `ui.go` | — |

### Fonctions principales par pôle

**Personne 1**
`initCharacter`, `characterCreation`, `displayInfo`, `isDead`,
`accessInventory`, `addInventory`, `removeInventory`, `checkInventoryCapacity`,
`takePot`, `poisonPot`, `spellBook`

**Personne 2**
`marchand`, `buyItem`, `upgradeInventorySlot`,
`forgeron`, `hasResources`, `craftEquipment`, `equipItem`, `desequiper`

**Personne 3**
`initGoblin`, `initBoss`, `goblinPattern`, `characterTurn`,
`trainingFight`, `bossFight`, `lancerCombat`, `dessinerMonstre`

---

## 2. Structures communes (décidées ensemble en phase 1)

C'est le point le plus important : **ces structures sont figées dès le début**,
comme ça chacun peut travailler sans casser le code des autres.

```
Character                 Monster                  Equipment
 ├── Name                  ├── Name                  ├── Head
 ├── Class                 ├── MaxHP                 ├── Body
 ├── Level                 ├── CurrentHP             └── Feet
 ├── MaxHP                 ├── Attack
 ├── CurrentHP             ├── Initiative
 ├── Inventory []string    ├── XP
 ├── InventoryMax          └── EstBoss
 ├── Upgrades
 ├── Skills []string
 ├── Money
 ├── Equip (Equipment)
 ├── Mana / MaxMana
 ├── XP / XPMax
 └── Initiative
```

### Conventions de nommage
- Les **structures et fonctions imposées par le sujet** gardent leur nom anglais
  (`Character`, `initCharacter`, `takePot`, `goblinPattern`…).
- Les fonctions « maison » sont en français (`marchand`, `forgeron`, `lancerCombat`…).
- Les commentaires sont en français, au-dessus de chaque fonction, avec le numéro de tâche.
- Un fichier = un thème. Pas de fonction d'affichage ailleurs que dans `ui.go`.

---

## 3. Organisation Git

```
main
│
├── feature/personnage    (Personne 1)
├── feature/economie      (Personne 2)
└── feature/combat        (Personne 3)
```

Règles de l'équipe :
- Personne ne code directement sur `main`.
- On merge dans `main` **au minimum une fois par jour** (fin de journée).
- Avant de merger : `go build ./src` doit passer, et le jeu doit se lancer.
- Message de commit : `[pole] ce qui a été fait`
  → ex. `[combat] ajout de goblinPattern et de l'attaque a 200%`

Commandes de base :
```bash
git checkout -b feature/combat      # créer sa branche
git add . && git commit -m "[combat] init du gobelin"
git push origin feature/combat

git checkout main && git pull       # récupérer le travail des autres
git merge feature/combat
```

---

## 4. Planning sur 5 jours

| Jour | Personne 1 | Personne 2 | Personne 3 |
|---|---|---|---|
| **J1 matin** | Réunion commune : structures, fichiers, menus, conventions | ← | ← |
| **J1 aprèm** | Tâches 1-3 (Character, init, displayInfo) | Prépare le catalogue d'articles + prix | Tâche 19 (structure Monster, initGoblin) |
| **J2 matin** | Tâches 4-6 (inventaire, potion, menu) | Tâche 13-14 (argent, marchand) | Tâche 20 (goblinPattern, 200 % tous les 3 tours) |
| **J2 aprèm** | Tâches 7-9 (marchand de base, mort, poison) | Tâche 15 (forgeron, recettes) | Tâche 21 (characterTurn, menu de combat) |
| **J3 matin** | Tâches 10-12 (sorts, classes, limite) | Tâches 16-17 (Equipment, équiper, bonus PV) | Tâche 22 (trainingFight, boucle de combat) |
| **J3 aprèm** | **Intégration n°1** : tout le monde merge dans `main` et on teste ensemble | ← | ← |
| **J4 matin** | Mission 2 (XP / niveaux) | Tâche 18 (agrandissement du sac) | Mission 1 (initiative) |
| **J4 aprèm** | Mission 4 (mana) | Mission 5 (améliorations, messages d'erreur) | Mission 3 (sorts en combat) + boss final ASCII |
| **J5 matin** | Mission 6 + README + docs | Tests de tous les cas d'erreur | Équilibrage du combat |
| **J5 aprèm** | **Répétition de l'oral** (démo CLI + explication du code) | ← | ← |

---

## 5. Tests d'intégration à faire ensemble

Ce sont les scénarios qui traversent les 3 pôles — c'est là que ça casse.

- [ ] Potion achetée → apparaît dans l'inventaire → utilisable en combat
- [ ] Livre de sort acheté → utilisé → le sort apparaît dans la fiche personnage → utilisable en combat
- [ ] Ressources achetées → forgeron → équipement fabriqué → équipé → PV max augmentés
- [ ] Sac plein → le marchand refuse la vente (message clair)
- [ ] Pas assez d'or / pas assez de ressources → messages d'erreur différents
- [ ] Équiper un 2ᵉ chapeau → le premier revient dans le sac, les PV max sont recalculés
- [ ] Le joueur tombe à 0 PV en combat → « WASTED » → résurrection à 50 % → retour au menu
- [ ] Le gobelin tape bien 10 dégâts aux tours 3, 6, 9…
- [ ] Sort lancé sans assez de mana → refusé, le tour n'est pas perdu

---

## 6. Points d'attention pour l'oral

- Montrer **une démo qui tourne** (créer un perso, acheter, crafter, équiper, combattre).
- Expliquer **pourquoi** on a découpé le code en fichiers, et ce que fait chacun.
- Savoir expliquer une fonction au hasard : `goblinPattern`, `equipItem`, `removeInventory`.
- Assumer les choix d'univers et les améliorations personnelles (boss final, phases ASCII).
- Répartir la parole : chacun présente son pôle.
