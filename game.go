package main

type Weapon struct {
	Damage          [2]int
	StrengthReq     int
	AgilityReq      int
	IntelligenceReq int
}

type Consummables struct {
	Duration        int
	HitsPointEffect int
	StrengthEffect  int
	AgilityEffect   int
	IntellectEffect int
}

type Gopher struct {
	Hitpoints int
	Weapon    Weapon
	Inventory []Consummables
	Strength  int
	Agility   int
	Intellect int
	Coins     int
}
