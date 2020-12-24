package generator

import (
	"github.com/df-mc/dragonfly/dragonfly/block"
	"github.com/df-mc/dragonfly/dragonfly/block/wood"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/dragonfly/dragonfly/world/chunk"
	"math/rand"
)

type Skygrid struct {
}

var blocks = [...]world.Block{block.Chest{}, block.Grass{}, block.Log{Wood: wood.Oak()}, block.Log{Wood: wood.Acacia()}, block.Log{Wood: wood.Jungle()},
	block.Andesite{}, block.Stone{Smooth: true}, block.Stone{}, block.Cobblestone{}, block.Cobblestone{Mossy: true}, block.CoalOre{}, block.CoalBlock{},
	block.IronOre{}, block.IronBlock{}, block.Leaves{Wood: wood.Jungle()}, block.Leaves{Wood: wood.Jungle()},
}

func (s Skygrid) GenerateChunk(pos world.ChunkPos, chunk *chunk.Chunk) {
	for x := uint8(0); x < 16; x += 4 {
		for y := uint8(0); y < 16; y += 4 {
			s.GenerateLine(x, y, chunk)
		}
	}
}

func (s Skygrid) GenerateLine(x, z uint8, chunk *chunk.Chunk) {
	for y := uint8(0); y < 225; y += 4 {
		s.GenerateBlock(x, y, z, chunk)
	}
}

func (s Skygrid) GenerateBlock(x, y, z uint8, chunk *chunk.Chunk) {
	chunk.SetRuntimeID(x, y, z, 0, s.RandomBlock())
}

func (s Skygrid) RandomBlock() (j uint32) {
	block := rand.Intn(len(blocks) - 1)
	j, _ = world.BlockRuntimeID(blocks[block])
	return
}
