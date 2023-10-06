package main

import "math"

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
	Hitpoints int
	Weapon    Weapon
	Inventory []Consummable
	Strength  int
	Agility   int
	Intellect int
	Coins     int
}

func main() {

	//Weapons
	Weapons := make(map[string]Weapon)
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
		CoinsReq:       5,
		HitPointEffect: 5,
		Duration:       math.MaxInt64,
		StrengthEffect: 0,
		AgilityEffect: 0,
		IntellectEffect: 0,
	}
	Consummables["strength_potion"] = Consummable{
		CoinsReq:       10,
		HitPointEffect: 0,
		Duration:       3,
		StrengthEffect: 3,
		AgilityEffect: 0,
		IntellectEffect: 0,
	}
	Consummables["agility_potion"] = Consummable{
		CoinsReq:       10,
		HitPointEffect: 0,
		Duration:       3,
		StrengthEffect: 0,
		AgilityEffect: 3,
		IntellectEffect: 0,
	}
	Consummables["intellect_potion"] = Consummable{
		CoinsReq:       10,
		HitPointEffect: 0,
		Duration:       3,
		StrengthEffect: 0,
		AgilityEffect: 0,
		IntellectEffect: 3,
	}

}
