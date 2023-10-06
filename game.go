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
	Buy(item string, weapons map[string]Weapon, consumables map[string]Consumable)
	Use(Item string, consumables map[string]Consumable)
	Train(skill string)
}

type Weapon struct {
	Damage          [2]int
	StrengthReq     int
	AgilityReq      int
	IntelligenceReq int
	CoinsReq        int
}

type Consumable struct {
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
	Inventory []Consumable
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
func (player Gopher) Buy(item string, weapons map[string]Weapon, consumables map[string]Consumable) Gopher{
	item = strings.TrimSpace(item)
	if value, exists := weapons[item]; exists {
        fmt.Printf("%s exists in the weapons, and its value is %d\n", item, value)
		if player.Coins >= value.CoinsReq && player.Strength >= value.StrengthReq && player.Agility >= value.AgilityReq &&
			player.Intellect >= value.IntelligenceReq {
				player.Weapon = value
				player.Coins -= value.CoinsReq

		} else {
			fmt.Println("Cant buy weapon")
		}
		return player
    }

	if value, exists := consumables[item]; exists {
		fmt.Printf("%s exists in the consumables, and its value is %d\n", item, value)
		if player.Coins >= value.CoinsReq {
			player.Coins -= value.CoinsReq
			player.Inventory = append(player.Inventory, value)
		} else {
			fmt.Println("Cant buy consumable")
		}
		return player
	}


	fmt.Println("Invalid Operation / Command")
	return player
}

// Use
func (player Gopher) Use(item string, consummables map[string]Consumable) {
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
	Consummables := make(map[string]Consumable)
	Consummables["health_potion"] = Consumable{
		CoinsReq:        5,
		HitPointEffect:  5,
		Duration:        math.MaxInt32,
		StrengthEffect:  0,
		AgilityEffect:   0,
		IntellectEffect: 0,
	}
	Consummables["strength_potion"] = Consumable{
		CoinsReq:        10,
		HitPointEffect:  0,
		Duration:        3,
		StrengthEffect:  3,
		AgilityEffect:   0,
		IntellectEffect: 0,
	}
	Consummables["agility_potion"] = Consumable{
		CoinsReq:        10,
		HitPointEffect:  0,
		Duration:        3,
		StrengthEffect:  0,
		AgilityEffect:   3,
		IntellectEffect: 0,
	}
	Consummables["intellect_potion"] = Consumable{
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
		Inventory: []Consumable{},
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
					action[0] = strings.TrimSpace(action[0])
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
						player1 = player1.Buy(action[1], Weapons, Consummables)
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
					action[0] = strings.TrimSpace(action[0])
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
						player2 = player2.Buy(action[1], Weapons, Consummables)
					}
					if action[0] == "use" {
						player2.Use(action[1], Consummables)
					}
					if action[0] == "train" {
						player2 = player2.Train(action[1])
					}
				}
			}
			turn++
		}
	}
}
