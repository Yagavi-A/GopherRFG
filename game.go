package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
)

type Action interface {
	Attack(opponent Gopher)
	Work()
	Buy(item string, weapons map[string]Weapon, consumables map[string]Consummable)
	Use(Item string, consumables map[string]Consummable)
	Train(skill string)
}

type Weapon struct {
	Damage          [2]int
	StrengthReq     int
	AgilityReq      int
	IntelligenceReq int
	CoinsReq        int
}

type Consummable struct {
	Duration        int
	HitPointEffect  int
	StrengthEffect  int
	AgilityEffect   int
	IntellectEffect int
	CoinsReq        int
}

type Gopher struct {
	Name      string
	Hitpoints int
	Weapon    Weapon
	Inventory []Consummable
	Strength  int
	Agility   int
	Intellect int
	Coins     int
}

func ExitGame() string {
	return "Game Over"
}

// Attack
func (player Gopher) Attack(opponent Gopher) Gopher {
	fmt.Println(player.Name, " Attacked ", opponent.Name)
	leastDamage := player.Weapon.Damage[0]
	highestDamage := player.Weapon.Damage[1]
	damageCaused := rand.Intn(highestDamage-leastDamage+1) + leastDamage
	opponent.Hitpoints -= damageCaused
	return opponent

}

// Work
func (player Gopher) Work() Gopher {
	coinsEarned := rand.Intn(11) + 5
	player.Coins += coinsEarned
	fmt.Printf("%s worked and earned %d coins \n", player.Name, coinsEarned)
	return player
}

// Buy
func (player Gopher) Buy(item string, weapons map[string]Weapon, consummables map[string]Consummable) {
	fmt.Println("Buy here")
}

// Use
func (player Gopher) Use(item string, consummables map[string]Consummable) {
	fmt.Println("Use here")
}

// Train
func (player Gopher) Train(skill string) Gopher {
	skill = strings.TrimSpace(skill)
	if player.Coins < 5 {
		fmt.Println("Not enough coins")
		return player
	}

	if skill == "strength" {
		player.Strength += 2
	} else if skill == "agility" {
		player.Agility += 2
	} else if skill == "intellect" {
		player.Intellect += 2
	} else {
		fmt.Println("Invalid skill")
	}

	player.Coins -= 5
	fmt.Printf("%s trained in %s and increased from %s to 2.\n", player.Name, skill, skill)
	return player
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	//Weapons
	Weapons := make(map[string]Weapon)
	Weapons["hand"] = Weapon{
		Damage:          [2]int{1, 1},
		StrengthReq:     0,
		AgilityReq:      0,
		IntelligenceReq: 0,
		CoinsReq:        0,
	}
	Weapons["knife"] = Weapon{
		Damage:          [2]int{2, 3},
		StrengthReq:     0,
		AgilityReq:      0,
		IntelligenceReq: 0,
		CoinsReq:        10,
	}
	Weapons["sword"] = Weapon{
		Damage:          [2]int{3, 5},
		StrengthReq:     2,
		AgilityReq:      0,
		IntelligenceReq: 0,
		CoinsReq:        35,
	}
	Weapons["ninjaku"] = Weapon{
		Damage:          [2]int{1, 7},
		StrengthReq:     0,
		AgilityReq:      2,
		IntelligenceReq: 0,
		CoinsReq:        25,
	}
	Weapons["wand"] = Weapon{
		Damage:          [2]int{3, 3},
		StrengthReq:     0,
		AgilityReq:      0,
		IntelligenceReq: 2,
		CoinsReq:        30,
	}
	Weapons["gophermourne"] = Weapon{
		Damage:          [2]int{6, 7},
		StrengthReq:     3,
		AgilityReq:      0,
		IntelligenceReq: 2,
		CoinsReq:        65,
	}

	//Consumables
	Consummables := make(map[string]Consummable)
	Consummables["health_potion"] = Consummable{
		CoinsReq:        5,
		HitPointEffect:  5,
		Duration:        math.MaxInt64,
		StrengthEffect:  0,
		AgilityEffect:   0,
		IntellectEffect: 0,
	}
	Consummables["strength_potion"] = Consummable{
		CoinsReq:        10,
		HitPointEffect:  0,
		Duration:        3,
		StrengthEffect:  3,
		AgilityEffect:   0,
		IntellectEffect: 0,
	}
	Consummables["agility_potion"] = Consummable{
		CoinsReq:        10,
		HitPointEffect:  0,
		Duration:        3,
		StrengthEffect:  0,
		AgilityEffect:   3,
		IntellectEffect: 0,
	}
	Consummables["intellect_potion"] = Consummable{
		CoinsReq:        10,
		HitPointEffect:  0,
		Duration:        3,
		StrengthEffect:  0,
		AgilityEffect:   0,
		IntellectEffect: 3,
	}

	//2 players
	player1 := Gopher{
		Name:      "PLAYER>1",
		Hitpoints: 3,
		Coins:     20,
		Inventory: []Consummable{},
		Strength:  0,
		Agility:   0,
		Intellect: 0,
		Weapon:    Weapons["hand"],
	}

	player2 := player1
	player2.Name = "PLAYER>2"

	var command string
	var turn int = 0
	for {
		command, _ = reader.ReadString('\n')
		if command == "exit" {
			fmt.Println(ExitGame())
			break
		} else {
			if turn%2 == 0 {
				action := strings.Split(command, " ")
				if len(action) == 1 {
					if action[0] == "attack" {
						player2 = player1.Attack(player2)
						if player2.Hitpoints <= 0 {
							fmt.Println(ExitGame())
							break
						}
					}
					if action[0] == "work" {
						player1 = player1.Work()
					}
				} else {
					if action[0] == "buy" {
						player1.Buy(action[1], Weapons, Consummables)
					}
					if action[0] == "use" {
						player1.Use(action[1], Consummables)
					}
					if action[0] == "train" {
						player1 = player1.Train(action[1])
					}
				}
			} else {
				action := strings.Split(command, " ")
				if len(action) == 1 {
					if action[0] == "attack" {
						player1 = player2.Attack(player1)
						if player1.Hitpoints <= 0 {
							fmt.Println(ExitGame())
							break
						}
					}
					if action[0] == "work" {
						player2 = player2.Work()
					}
				} else {
					if action[0] == "buy" {
						player2.Buy(action[1], Weapons, Consummables)
					}
					if action[0] == "use" {
						player2.Use(action[1], Consummables)
					}
					if action[0] == "train" {
						player2.Train(action[1])
					}
				}
			}
			turn++
		}
	}

}
