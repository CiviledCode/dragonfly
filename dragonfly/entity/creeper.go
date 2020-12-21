package entity

import (
	"github.com/df-mc/dragonfly/dragonfly/entity/damage"
	"github.com/df-mc/dragonfly/dragonfly/entity/healing"
	"github.com/go-gl/mathgl/mgl64"
	"sync/atomic"
)

type Creeper struct {
	Primed, Powered bool

	ExplosionRadius, Fuse uint8

	pos atomic.Value

	health, maxHealth, speed float64

	damageImmune bool
}

func (c *Creeper) Health() float64 {
	return c.health
}

func (c *Creeper) MaxHealth() float64 {
	return c.maxHealth
}

func (c *Creeper) SetMaxHealth(health float64) {
	c.maxHealth = health
}

func (c *Creeper) AttackImmune() bool {
	return c.damageImmune
}

func (c *Creeper) SetImmune(immune bool) {
	c.damageImmune = immune
}

func (c *Creeper) Hurt(damage float64, source damage.Source) {
	c.health -= damage
}

func (c *Creeper) Heal(health float64, source healing.Source) {
	c.health += health
}

func (c *Creeper) KnockBack(src mgl64.Vec3, force, height float64) {
	//TODO: Implement Knockback
}

func (c *Creeper) Speed() float64 {
	return c.speed
}

func (c *Creeper) SetSpeed(speed float64) {
	c.speed = speed
}

func (c *Creeper) Pathfind() {
	//TODO: Implement pathfinding
}

// Ignite is called when the creeper was interacted with while holding a flint and steel.
func (c *Creeper) Ignite() {
	c.Primed = true
}

// Charge allows us to charge the entity for higher explosion damage and radius
func (c *Creeper) Charge(charged bool) {

}

// Blowup makes the Creeper blowup immediately without any priming
func (c *Creeper) Blowup() {

}

// Prime makes the Creeper expand and act as if it's going to blowup.
// This takes into account the fuse timer set inside of struct
func (c *Creeper) Prime() {

}
