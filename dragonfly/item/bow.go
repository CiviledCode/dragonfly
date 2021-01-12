package item

// Bow is an item that allows you to shoot arrows at other players or blocks.
type Bow struct {
	Releasable
}

func FuelTime() int {
	return 200
}

func getMaxDurability() int {
	return 385
}

// EncodeItem ...
func (b Bow) EncodeItem() (id int32, meta int16) {
	return 261, 0
}
