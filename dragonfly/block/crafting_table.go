package block

import (
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/item/tool"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

type CraftingTable struct {
	noNBT
	bass
	solid
	Activatable
}

func (c CraftingTable) EncodeItem() (id int32, meta int16) {
	return 58, 0
}

// EncodeBlock ...
func (c CraftingTable) EncodeBlock() (name string, properties map[string]interface{}) {
	//noinspection SpellCheckingInspection
	return "minecraft:crafting_table", map[string]interface{}{}
}

// BreakInfo ...
func (c CraftingTable) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 2.5,
		Harvestable: func(_ tool.Tool) bool {
			return true
		},
		Effective: axeEffective,
		Drops:     simpleDrops(item.NewStack(c, 1)),
	}
}

func (c CraftingTable) Activate(pos world.BlockPos, _ world.Face, _ *world.World, u item.User) {
	if opener, ok := u.(ContainerOpener); ok {
		opener.OpenBlockContainer(pos)
	}
}
