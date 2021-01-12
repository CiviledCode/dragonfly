package math

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func spaceship(first, second float64) float64 {
	if first == second {
		return 0
	} else if second > first {
		return -1
	} else if first > second {
		return 1
	}
	panic("should never happen")
}

func BetweenPoints(start, end mgl64.Vec3) []mgl64.Vec3 {
	currentBlock := mgl64.Vec3{math.Floor(start.X()), math.Floor(start.Y()), math.Floor(start.Z())}
	directionVector := end.Sub(start).Normalize()

	if directionVector.LenSqr() <= 0 {
		panic("start and end points are the same, giving a zero direction vector")
	}

	radius := world.Distance(start, end)

	stepX := spaceship(directionVector.X(), 0)
	stepY := spaceship(directionVector.Y(), 0)
	stepZ := spaceship(directionVector.Z(), 0)

	tMaxX := rayTraceDistanceToBoundary(start.X(), directionVector.X())
	tMaxY := rayTraceDistanceToBoundary(start.Y(), directionVector.Y())
	tMaxZ := rayTraceDistanceToBoundary(start.Z(), directionVector.Z())

	var tDeltaX float64
	var tDeltaY float64
	var tDeltaZ float64

	if directionVector.X() != 0 {
		tDeltaX = stepX / directionVector.X()
	}

	if directionVector.Y() != 0 {
		tDeltaY = stepY / directionVector.Y()
	}

	if directionVector.Z() != 0 {
		tDeltaZ = stepZ / directionVector.Z()
	}

	var locations []mgl64.Vec3
	for {
		fmt.Println(currentBlock)
		locations = append(locations, currentBlock)

		if tMaxX < tMaxY && tMaxX < tMaxZ {
			if tMaxX > radius {
				break
			}
			currentBlock = currentBlock.Add(mgl64.Vec3{stepX, 0, 0})
			tMaxX += tDeltaX
		} else if tMaxY < tMaxZ {
			if tMaxY > radius {
				break
			}
			currentBlock = currentBlock.Add(mgl64.Vec3{0, stepY, 0})
			tMaxY += tDeltaY
		} else {
			if tMaxZ > radius {
				break
			}
			currentBlock = currentBlock.Add(mgl64.Vec3{0, 0, stepZ})
			tMaxZ += tDeltaZ
		}
	}
	return locations
}

func rayTraceDistanceToBoundary(s, ds float64) float64 {
	if ds == 0 {
		return 0x7FF0000000000000
	}

	if ds < 0 {
		s = -s
		ds = -ds

		if math.Floor(s) == s {
			return 0
		}
	}

	return (1 - (s - math.Floor(s))) / ds
}
