package entity

import (
	"github.com/df-mc/dragonfly/dragonfly/entity/physics"
	"github.com/df-mc/dragonfly/dragonfly/entity/state"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"go.uber.org/atomic"
)

type Arrow struct {
	velocity, pos              atomic.Value
	yaw, pitch                 float64
	points                     []mgl64.Vec3
	pointIndex                 int
	inBlock, inEntity          bool
	entityIntersectionPosition mgl64.Vec3
	*MovementComputer
}

// NewArrow ...
func NewArrow(pos, velocity mgl64.Vec3, yaw, pitch float64, force bool) *Arrow {
	a := &Arrow{}
	a.pos.Store(pos)
	a.velocity.Store(velocity)
	a.yaw = yaw
	a.pitch = pitch
	return a
}

// tickMovement performs the movement and velocity decreases of the arrow.
// if the arrow hits a block or is stuck inside of a block, tickMovement will still be called but nothing will happen
func (a *Arrow) tickMovement(e world.Entity) mgl64.Vec3 {

	// Check if the arrow is stuck in a block
	if !a.inBlock && !a.inEntity {
		// Decrease the velocity by multiplying it by .99
		velocity := a.Velocity().Mul(.99)
		a.SetVelocity(velocity)

		lastPoint := a.points[a.pointIndex-1]

		//TODO: Calculate the current points location

		// Calculate the current point and update it
		var currentPoint mgl64.Vec3

		currentPoint = mgl64.Vec3{lastPoint.X() + .2, lastPoint.Y() - .1, lastPoint.Z()}

		/*
			// Check if the current point is a block so we can set inBlock to true, terminating movement calculations
			blockAtPoint := a.World().Block(world.BlockPos{int(currentPoint.X()), int(currentPoint.Y()), int(currentPoint.Z())})
			if _, ok := blockAtPoint.(block.Air); !ok {
				a.inBlock = true
			} else if _, ok := blockAtPoint.(block.Lava); ok {

			}
		*/
		//TODO: Check if the arrow is going through water or lava for slowdowns

		// Check if the arrow is intersecting with any entities
		if len(a.World().EntitiesWithin(a.AABB())) > 0 {
			a.inEntity = true
			//TODO: Deal damage to the entity based on velocity and do effects
		}

		// Store the current position of the area and move the movable entity
		a.pos.Store(currentPoint)
		a.move(a, currentPoint, a.World().Viewers(currentPoint))
		return currentPoint
	}

	if a.inEntity {
		// Fill this out to move with the entity in the correct position
	}
	a.pointIndex++

	return mgl64.Vec3{}
}

func (a *Arrow) Name() string {
	return "Arrow"
}

// Close ...
func (a *Arrow) Close() error {
	if a.World() != nil {
		a.World().RemoveEntity(a)
	}
	return nil
}

// AABB ...
func (a *Arrow) AABB() physics.AABB {
	return physics.NewAABB(mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0.5, 0.5, 0.5})
}

// Position ...
func (a *Arrow) Position() mgl64.Vec3 {
	return a.pos.Load().(mgl64.Vec3)
}

// World ...
func (a *Arrow) World() *world.World {
	w, _ := world.OfEntity(a)
	return w
}

// Yaw ...
func (a *Arrow) Yaw() float64 {
	return a.yaw
}

// OnGround ...
func (a *Arrow) OnGround() bool {
	return false
}

// Pitch ...
func (a *Arrow) Pitch() float64 {
	return a.pitch
}

// State ...
func (a *Arrow) State() []state.State {
	return nil
}

// Velocity ...
func (a *Arrow) Velocity() mgl64.Vec3 {
	return a.velocity.Load().(mgl64.Vec3)
}

// SetVelocity ...
func (a *Arrow) SetVelocity(v mgl64.Vec3) {
	a.velocity.Store(v)
}

// EncodeEntity ...
func (a *Arrow) EncodeEntity() string {
	return "minecraft:arrow"
}
