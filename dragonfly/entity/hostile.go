package entity

import "github.com/df-mc/dragonfly/dragonfly/world"

type Hostile interface {
	Living

	// AttackDamage() returns the amount of base damage the entity inflicts when attacking
	AttackDamage() float64

	// SetAttackDamage() sets the base attack damage that the entity inflicts when attacking
	SetAttackDamage(float64)

	// DiscoveryRange() returns the range in which a mob is able to see a target
	DiscoveryRange() uint8

	// SetTarget() sets the target that the entity should be focused on
	SetTarget(entity world.Entity)

	// Target() returns the target the entity is focused on
	Target() world.Entity

	// DayLightSensitive() returns if the entity burns in daylight
	DayLightSensitive() bool

	// SpawnLightRequired() returns the maximum allowed light for the entity to spawn properly
	SpawnLightRequired() uint8

	// AttackDistance() returns the distance in which an entity can attack from
	AttackDistance() uint8

	// Attack() is called when the distance between the target and the entity is smaller than AttackDistance()
	Attack()
}
