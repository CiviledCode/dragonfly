package pathfinding

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
	"golang.org/x/image/math/f32"
)

type EOmap [][]uint8

// NewEOMap takes the blocks within a defined radius and puts the y difference from the origin into an object
func NewEOMap(wrld world.World, origin f32.Vec3, radius int) (eo EOmap) {
	var xPoint, zPoint uint8 = 0, 0
	for x := int(origin[0]) - radius; x < int(origin[0])+radius; x++ {
		for z := int(origin[2]) - radius; z < int(origin[2])+radius; z++ {
			//TODO: Get the highest relative block at a X, Z coordinate

			// TODO: Find a better implementation of this
			// Increment from 2 blocks above our last position to 15 blocks below it
			/*
				for rel := int(eo[xPoint][zPoint-1])+2; rel > rel - 15; rel-- {
					if wrld.Block(x, rel, z) != block.Air{}.(block.) {

					}
				}
			*/
			wrld.Block(world.BlockPos{x, 2, z})

			zPoint++
		}
		zPoint = 0
		xPoint++
	}
	return
}
