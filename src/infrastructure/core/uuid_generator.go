package core

import (
	"github.com/google/uuid"
	"github.com/safalife/core-api/src/domain/core"
)

// GoogleUUID wraps google/uuid.UUID to implement core.UUID interface
type GoogleUUID struct {
	uuid.UUID
}

// String returns the string representation of the UUID
func (u GoogleUUID) String() string {
	return u.UUID.String()
}

// NewGoogleUUID creates a new GoogleUUID from uuid.UUID
func NewGoogleUUID(id uuid.UUID) GoogleUUID {
	return GoogleUUID{UUID: id}
}

// ToGoogleUUID converts core.UUID to uuid.UUID for database operations
func ToGoogleUUID(coreUUID core.UUID) uuid.UUID {
	if coreUUID == nil {
		return uuid.Nil
	}

	if googleUUID, ok := coreUUID.(GoogleUUID); ok {
		return googleUUID.UUID
	}

	// Fallback: parse from string
	parsed, _ := uuid.Parse(coreUUID.String())
	return parsed
}

// FromGoogleUUID converts uuid.UUID to core.UUID
func FromGoogleUUID(googleUUID uuid.UUID) core.UUID {
	return GoogleUUID{UUID: googleUUID}
}

// UUIDGenerator implements core.UUIDGenerator using google/uuid
type UUIDGenerator struct{}

// NewUUIDGenerator creates a new UUID generator
func NewUUIDGenerator() core.UUIDGenerator {
	return &UUIDGenerator{}
}

// New generates a new UUID
func (g *UUIDGenerator) New() core.UUID {
	return GoogleUUID{UUID: uuid.New()}
}

// Parse parses a string into a UUID
func (g *UUIDGenerator) Parse(s string) (core.UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return nil, err
	}
	return GoogleUUID{UUID: parsed}, nil
}

// Nil returns a nil UUID
func (g *UUIDGenerator) Nil() core.UUID {
	return GoogleUUID{UUID: uuid.Nil}
}

func ToNullGoogleUUID(coreUUID *core.UUID) *uuid.UUID {
	if coreUUID == nil || *coreUUID == nil {
		return nil
	}

	googleUUID := ToGoogleUUID(*coreUUID)
	return &googleUUID
}

func FromNullGoogleUUID(googleUUID *uuid.UUID) *core.UUID {
	if googleUUID == nil {
		return nil
	}

	coreUUID := FromGoogleUUID(*googleUUID)
	return &coreUUID
}
