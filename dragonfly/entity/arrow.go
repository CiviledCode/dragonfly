package entity

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/entity/physics"
	"github.com/df-mc/dragonfly/dragonfly/entity/state"
	"github.com/df-mc/dragonfly/dragonfly/math"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"go.uber.org/atomic"
)

type Arrow struct {
	velocity, pos atomic.Value
	yaw, pitch    float64
	hasTicked     bool
	points        []mgl64.Vec3
	pointIndex    int
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

// tickMovement performs a movement tick on an entity. Velocity is applied and changed according to the values
// of its drag and gravity.
// The new position of the entity after movement is returned.
func (a *Arrow) tickMovement(e world.Entity) mgl64.Vec3 {
	if !a.hasTicked {
		a.hasTicked = true

		velocity := a.velocity.Load().(mgl64.Vec3)

		start := a.Position()
		end := start.Add(velocity)
		a.points = math.BetweenPoints(start, end)
	} else {
		if a.pointIndex != len(a.points) {
			point := a.points[a.pointIndex]

			a.World().Block(world.BlockPos{int(point.X()), int(point.Y()), int(point.Z())})
			fmt.Println("Moved to", point)
			a.pos.Store(a.points[a.pointIndex])
			a.move(a, a.points[a.pointIndex], a.World().Viewers(a.points[a.pointIndex]))
			a.pointIndex++
		}
	}
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
