package entity

type Neutral interface {
	Hostile

	//IsHostile() checks if the neutral mob is hostile in that instance
	IsHostile() bool

	//SetHostile() will try and make the entity hostile towards players and villagers
	SetHostile(hostile bool)
}
