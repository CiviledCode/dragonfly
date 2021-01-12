package item

// Arrow is a projectile that is shot from a bow.
type Arrow struct{}

// EncodeItem ...
func (a Arrow) EncodeItem() (id int32, meta int16) {
	return 262, 0
}
