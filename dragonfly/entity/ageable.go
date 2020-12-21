package entity

// Ageable represents a living entity that has both an adult and baby form
type Ageable interface {
	Living

	// IsBaby() checks if the mob is in it's baby form or not
	IsBaby() bool
}
