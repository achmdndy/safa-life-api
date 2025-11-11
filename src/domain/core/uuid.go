package core

// UUID represents a universally unique identifier
type UUID interface {
	// String returns the string representation of the UUID
	String() string
}

// UUIDGenerator provides an abstraction for UUID generation
type UUIDGenerator interface {
	// New generates a new UUID
	New() UUID
	// Parse parses a string into a UUID
	Parse(s string) (UUID, error)
	// Nil returns a nil UUID
	Nil() UUID
}
